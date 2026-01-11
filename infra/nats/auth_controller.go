package nats

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mynaparrot/plugnmeet-protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/nats-io/nkeys"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

const (
	RecorderUserAuthName      = "PLUGNMEET_RECORDER_AUTH"
	TranscoderConsumerDurable = "transcoderWorker"
)

type NatsAuthController struct {
	apiKey        string
	apiSecret     string
	natsInfo      *NatsInfo
	nc            *nats.Conn
	issuerKeyPair nkeys.KeyPair
	curveKeyPair  nkeys.KeyPair
	logger        *logrus.Entry
}

func NewNatsAuthController(apiKey, apiSecret string, info *NatsInfo, nc *nats.Conn, issuerKeyPair nkeys.KeyPair, curveKeyPair nkeys.KeyPair, logger *logrus.Entry) *NatsAuthController {
	return &NatsAuthController{
		apiKey:        apiKey,
		apiSecret:     apiSecret,
		natsInfo:      info,
		nc:            nc,
		issuerKeyPair: issuerKeyPair,
		curveKeyPair:  curveKeyPair,
		logger:        logger.WithField("sub-controller", "nats-auth"),
	}
}

func (s *NatsAuthController) Handle(r micro.Request) {
	var data []byte
	var err error

	xKey := r.Headers().Get("Nats-Server-Xkey")
	if len(xKey) > 0 {
		if s.curveKeyPair == nil {
			s.logger.Errorln("received encrypted data from nats server but curveKeyPair is nil")
			_ = r.Error("500", "xKey not supported", nil)
			return
		}

		data, err = s.curveKeyPair.Open(r.Data(), xKey)
		if err != nil {
			s.logger.WithError(err).Errorln("error decrypting message from nats server")
			_ = r.Error("500", err.Error(), nil)
			return
		}
	} else {
		data = r.Data()
	}

	rc, err := jwt.DecodeAuthorizationRequestClaims(string(data))
	if err != nil {
		s.logger.WithError(err).Errorln("error decoding authorization request")
		_ = r.Error("500", err.Error(), nil)
		return
	}

	userNkey := rc.UserNkey
	serverId := rc.Server.ID

	claims, err := s.handleClaims(rc)
	if err != nil {
		s.logger.WithError(err).Errorln("error handling claims")
		s.respond(r, userNkey, serverId, "", err)
		return
	}

	token, err := s.validateAndSign(claims, s.issuerKeyPair)
	s.respond(r, userNkey, serverId, token, err)
}

func (s *NatsAuthController) handleClaims(req *jwt.AuthorizationRequestClaims) (*jwt.UserClaims, error) {
	claims := jwt.NewUserClaims(req.UserNkey)
	claims.Audience = s.natsInfo.Account

	// nats v2.10.28 & v2.11.2 (#6808) Auth tokens are now redacted
	// but for our case it is necessary to have it, so we'll add here
	// otherwise user CONNECT & DISCONNECT logics will be breaking
	claims.Name = req.ConnectOptions.Token

	// check the info first
	data, err := auth.VerifyPlugNmeetAccessToken(s.apiKey, s.apiSecret, req.ConnectOptions.Token, 0)
	if err != nil {
		return nil, err
	}

	if data.GetName() == RecorderUserAuthName {
		s.setPermissionForRecorder(data, claims)
		return claims, nil
	}

	err = s.setPermissionForClient(data, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *NatsAuthController) setPermissionForRecorder(data *plugnmeet.PlugNmeetTokenClaims, claims *jwt.UserClaims) {
	pubAllow := jwt.StringList{
		"$JS.API.INFO",
		"_INBOX.>", // otherwise won't be able to send respond msg
		fmt.Sprintf("$JS.API.STREAM.INFO.KV_%s-%s", s.natsInfo.Recorder.RecorderInfoKv, data.GetUserId()),
		fmt.Sprintf("$JS.API.STREAM.UPDATE.KV_%s-%s", s.natsInfo.Recorder.RecorderInfoKv, data.GetUserId()),
		fmt.Sprintf("$JS.API.STREAM.CREATE.KV_%s-%s", s.natsInfo.Recorder.RecorderInfoKv, data.GetUserId()),
		fmt.Sprintf("$KV.%s-%s.>", s.natsInfo.Recorder.RecorderInfoKv, data.GetUserId()),
		fmt.Sprintf("$JS.API.DIRECT.GET.KV_%s-%s.>", s.natsInfo.Recorder.RecorderInfoKv, data.GetUserId()),
		// Allow publishing the job to the stream
		s.natsInfo.Recorder.TranscodingJobs,
		// Allow fetching the next message from the consumer & send ack
		fmt.Sprintf("$JS.API.CONSUMER.MSG.NEXT.%s.%s", s.natsInfo.Recorder.TranscodingJobs, TranscoderConsumerDurable),
		fmt.Sprintf("$JS.API.CONSUMER.INFO.%s.%s", s.natsInfo.Recorder.TranscodingJobs, TranscoderConsumerDurable),
		fmt.Sprintf("$JS.ACK.%s.%s.>", s.natsInfo.Recorder.TranscodingJobs, TranscoderConsumerDurable),
	}

	claims.Permissions = jwt.Permissions{
		Pub: jwt.Permission{
			Allow: pubAllow,
		},
		Sub: jwt.Permission{
			Allow: jwt.StringList{
				s.natsInfo.Recorder.RecorderChannel,
				"_INBOX.>",
			},
		},
	}
}

func (s *NatsAuthController) setPermissionForClient(data *plugnmeet.PlugNmeetTokenClaims, claims *jwt.UserClaims) error {
	roomId := data.GetRoomId()
	userId := data.GetUserId()

	payload, err := proto.Marshal(&plugnmeet.NatsSystemApiWorker{
		ApiTask: &plugnmeet.NatsSystemApiWorker_CreateConsumerWithPermission{
			CreateConsumerWithPermission: &plugnmeet.ApiWorkerTaskCreateConsumerWithPermissionReq{
				RoomId: roomId,
				UserId: userId,
			},
		},
	})
	if err != nil {
		s.logger.WithError(err).Errorln("error marshaling payload")
		return err
	}

	request, err := s.nc.Request(s.natsInfo.Subjects.SystemApiWorker, payload, time.Second*3)
	if err != nil {
		s.logger.WithError(err).Errorln("error sending request to system api worker")
		return err
	}

	res := new(plugnmeet.ApiWorkerTaskCreateConsumerWithPermissionRes)
	err = proto.Unmarshal(request.Data, res)
	if err != nil {
		s.logger.WithError(err).Errorf("error unmarshalling response: %s", string(request.Data))
		return err
	}
	if !res.Status {
		s.logger.Errorf("error creating consumer with permission %s", res.Msg)
		return errors.New(res.Msg)
	}

	var allowPub jwt.StringList
	err = json.Unmarshal([]byte(res.Msg), &allowPub)
	if err != nil {
		s.logger.WithError(err).Errorln("error unmarshalling response")
		return err
	}

	// Assign Permissions
	claims.Permissions = jwt.Permissions{
		Pub: jwt.Permission{
			Allow: allowPub,
		},
		Sub: jwt.Permission{
			Allow: jwt.StringList{
				"_INBOX.>", // otherwise break request-reply patterns
			},
		},
	}

	return nil
}

func (s *NatsAuthController) respond(req micro.Request, userNKey, serverId, userJWT string, err error) {
	rc := jwt.NewAuthorizationResponseClaims(userNKey)
	rc.Audience = serverId
	rc.Jwt = userJWT
	if err != nil {
		rc.Error = err.Error()
	}

	token, err := rc.Encode(s.issuerKeyPair)
	if err != nil {
		s.logger.WithError(err).Errorln("error encoding response jwt")
		_ = req.Respond(nil)
		return
	}
	data := []byte(token)

	// Check if encryption is required.
	xKey := req.Headers().Get("Nats-Server-Xkey")
	if len(xKey) > 0 {
		data, err = s.curveKeyPair.Seal(data, xKey)
		if err != nil {
			s.logger.WithError(err).Errorln("error encrypting response JWT")
			_ = req.Respond(nil)
			return
		}
	}

	_ = req.Respond(data)
}

func (s *NatsAuthController) validateAndSign(claims *jwt.UserClaims, kp nkeys.KeyPair) (string, error) {
	// Validate the claims.
	vr := jwt.CreateValidationResults()
	claims.Validate(vr)
	if len(vr.Errors()) > 0 {
		return "", errors.Join(vr.Errors()...)
	}

	// Sign it with the issuer key since this is non-operator mode.
	return claims.Encode(kp)
}
