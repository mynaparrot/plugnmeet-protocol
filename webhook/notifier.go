package webhook

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"github.com/frostbyte73/core"
	"github.com/google/uuid"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/livekit/protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strings"
	"time"
)

const (
	authHeader = "Authorization"
	// in various Apache modules will strip the Authorization header,
	// so we'll use additional one
	hashToken = "Hash-Token"
	maxRetry  = 2 // retryablehttp will use maxRetry + 1
)

type Notifier struct {
	client *retryablehttp.Client
	debug  bool
	worker core.QueueWorker
	logger *logrus.Logger
}

func newWebhookNotifier(queueSize int, debug bool, logger *logrus.Logger) *Notifier {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.RetryMax = maxRetry

	w := &Notifier{
		client: client,
		debug:  debug,
		logger: logger,
	}

	w.worker = core.NewQueueWorker(core.QueueWorkerParams{
		QueueSize:    queueSize,
		DropWhenFull: true,
	})

	return w
}

func (n *Notifier) AddInNotifyQueue(event *plugnmeet.CommonNotifyEvent, apiKey, apiSecret string, urls []string) {
	if len(urls) < 1 {
		return
	}

	for _, u := range urls {
		n.worker.Submit(func() {
			res, err := n.sendPostRequest(event, apiKey, apiSecret, u)
			if err != nil {
				n.logger.Errorln("failed to sendPostRequest webhook,", "url:", u, "event:", event.GetEvent(), "roomId:", event.GetRoom().GetRoomId(), "sid:", event.Room.GetSid(), "error:", err)
			} else {
				if n.debug {
					n.logger.Println("webhook sent for event:", event.GetEvent(), "roomID:", event.Room.GetRoomId(), "sid:", event.Room.GetSid(), "to URL:", u, "with http response code:", res.StatusCode, "msg:", res.Status)
				}
			}
		})
	}
}

func (n *Notifier) sendPostRequest(event *plugnmeet.CommonNotifyEvent, apiKey, apiSecret, url string) (*http.Response, error) {
	op := protojson.MarshalOptions{
		EmitUnpopulated: false,
		UseProtoNames:   true,
	}
	// make sure the event name is lowercase
	ev := strings.ToLower(event.GetEvent())
	event.Event = &ev

	if event.CreatedAt == nil {
		now := time.Now().UTC().Unix()
		event.CreatedAt = &now
	}
	if event.Id == nil {
		mId := uuid.NewString()
		event.Id = &mId
	}

	encoded, err := op.Marshal(event)
	if err != nil {
		return nil, err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return nil, err
	}
	r, err := retryablehttp.NewRequest("POST", url, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return nil, err
	}
	r.Header.Set(authHeader, token)
	// in various Apache modules will strip the Authorization header,
	// so we'll use additional one for easy use
	r.Header.Set(hashToken, token)
	// use a custom mime type to ensure the signature is checked before parsing
	r.Header.Set("content-type", "application/webhook+json")
	res, err := n.client.Do(r)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()

	return res, nil
}

var webhookNotifier *Notifier

func GetWebhookNotifier(queueSize int, debug bool, logger *logrus.Logger) *Notifier {
	if webhookNotifier != nil {
		return webhookNotifier
	}
	webhookNotifier = newWebhookNotifier(queueSize, debug, logger)

	return webhookNotifier
}
