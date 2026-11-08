package utils

import (
	"crypto/rand"
	"fmt"
	"io/fs"
	"math/big"
	rd "math/rand"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v3"
	"github.com/livekit/protocol/livekit"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var op = protojson.MarshalOptions{
	EmitUnpopulated: true,
	UseProtoNames:   true,
}

// StatusCodeToHTTPStatus maps a protocol StatusCode to the HTTP status of an API
// response. The protocol enum stays transport agnostic: HTTP specific cases must
// never be added as new StatusCode values, they are mapped here instead.
func StatusCodeToHTTPStatus(code plugnmeet.StatusCode) int {
	switch code {
	case plugnmeet.StatusCode_SUCCESS:
		return fiber.StatusOK
	case plugnmeet.StatusCode_INVALID_PARAMETERS, plugnmeet.StatusCode_MISSING_REQUIRED_PARAMETER:
		return fiber.StatusBadRequest
	case plugnmeet.StatusCode_INVALID_TOKEN_OR_SIGNATURE, plugnmeet.StatusCode_INVALID_API_KEY:
		// 401 for invalid credentials (token or API key), so clients know they
		// can retry with renewed credentials.
		return fiber.StatusUnauthorized
	case plugnmeet.StatusCode_PERMISSION_DENIED,
		plugnmeet.StatusCode_USER_BLOCKED, plugnmeet.StatusCode_USER_NOT_ACTIVE:
		// 403 so clients do not try to renew a token that is not the problem.
		return fiber.StatusForbidden
	case plugnmeet.StatusCode_NOT_FOUND, plugnmeet.StatusCode_ROOM_NOT_FOUND,
		plugnmeet.StatusCode_USER_NOT_FOUND:
		return fiber.StatusNotFound
	case plugnmeet.StatusCode_CONFLICT, plugnmeet.StatusCode_ROOM_NOT_ACTIVE:
		return fiber.StatusConflict
	default: // UNKNOWN_STATUS, INTERNAL_SERVER_ERROR and any future value
		return fiber.StatusInternalServerError
	}
}

// CommonErrorHTTPStatus returns the HTTP status for a failed response. The result
// is always non-2xx, even when a success code was passed by mistake.
func CommonErrorHTTPStatus(code plugnmeet.StatusCode) int {
	if status := StatusCodeToHTTPStatus(code); status >= 400 {
		return status
	}
	return fiber.StatusInternalServerError
}

func PrepareCommonWebhookNotifyEvent(event *livekit.WebhookEvent) *plugnmeet.CommonNotifyEvent {
	return &plugnmeet.CommonNotifyEvent{
		Event: &event.Event,
		Room: &plugnmeet.NotifyEventRoom{
			Sid:             &event.Room.Sid,
			RoomId:          &event.Room.Name,
			EmptyTimeout:    &event.Room.EmptyTimeout,
			MaxParticipants: &event.Room.MaxParticipants,
			CreationTime:    new(uint64(event.Room.CreationTime)),
			EnabledCodecs:   event.Room.EnabledCodecs,
			Metadata:        &event.Room.Metadata,
			NumParticipants: &event.Room.NumParticipants,
		},
		Participant: event.Participant,
		Track:       event.Track,
		Id:          &event.Id,
		CreatedAt:   &event.CreatedAt,
	}
}

// SendCommonProtobufResponse sends a CommonResponse as protobuf. It does not set
// an HTTP status, so it is only meant for successful responses; failures must use
// SendCommonErrorProtobufResponse.
func SendCommonProtobufResponse(c fiber.Ctx, s bool, m string) error {
	res := &plugnmeet.CommonResponse{
		Status: s,
		Msg:    m,
	}
	marshal, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/protobuf")
	return c.Send(marshal)
}

// SendCommonErrorProtobufResponse sends a failed response as a protobuf
// CommonResponse together with a non-2xx HTTP status derived from code.
func SendCommonErrorProtobufResponse(c fiber.Ctx, msg string, code plugnmeet.StatusCode) error {
	res := &plugnmeet.CommonResponse{
		Status:     false,
		Msg:        msg,
		StatusCode: code,
	}
	marshal, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/protobuf")
	return c.Status(CommonErrorHTTPStatus(code)).Send(marshal)
}

func SendProtobufResponse(c fiber.Ctx, res proto.Message) error {
	marshal, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/protobuf")
	return c.Send(marshal)
}

// SendCommonProtoJsonResponse sends a CommonResponse as protojson. It does not set
// an HTTP status, so it is only meant for successful responses; failures must use
// SendCommonErrorJsonResponse.
func SendCommonProtoJsonResponse(c fiber.Ctx, s bool, m string, statusCode plugnmeet.StatusCode) error {
	res := &plugnmeet.CommonResponse{
		Status:     s,
		Msg:        m,
		StatusCode: statusCode,
	}
	marshal, err := op.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/json")
	return c.Send(marshal)
}

// SendCommonErrorJsonResponse sends a failed response as a protojson CommonResponse
// together with a non-2xx HTTP status derived from code.
func SendCommonErrorJsonResponse(c fiber.Ctx, msg string, code plugnmeet.StatusCode) error {
	res := &plugnmeet.CommonResponse{
		Status:     false,
		Msg:        msg,
		StatusCode: code,
	}
	marshal, err := op.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/json")
	return c.Status(CommonErrorHTTPStatus(code)).Send(marshal)
}

func SendProtoJsonResponse(c fiber.Ctx, res proto.Message) error {
	marshal, err := op.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/json")
	return c.Send(marshal)
}

func GetFilesFromDir(path, ext, s string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext && checkIfAllowedFilePrefix(d.Name()) {
			files = append(files, d.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if s == "asc" {
		sort.Strings(files)
	} else if s == "des" {
		sort.Sort(sort.Reverse(sort.StringSlice(files)))
	}
	return files, nil
}

func checkIfAllowedFilePrefix(f string) bool {
	allowedPrefix := []string{"main", "runtime", "vendor", "tflite"}
	for _, p := range allowedPrefix {
		if strings.HasPrefix(f, p) {
			return true
		}
	}
	return false
}

func GenerateSecureRandomStrings(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateRandomStrings(n int) string {
	r := rd.New(rd.NewSource(time.Now().UnixNano()))
	b := make([]byte, n+2)
	r.Read(b)
	return fmt.Sprintf("%x", b)[2 : n+2]
}

func ParseFileSizeToBytes(sizeStr string) (uint64, error) {
	sizeStr = strings.TrimSpace(strings.ToLower(sizeStr))
	var multiplier uint64 = 1
	var numPart string

	for i, r := range sizeStr {
		if unicode.IsDigit(r) || r == '.' {
			numPart += string(r)
		} else {
			unit := strings.TrimSpace(sizeStr[i:])
			switch unit {
			case "kb", "k":
				multiplier = 1024
			case "mb", "m":
				multiplier = 1024 * 1024
			case "gb", "g":
				multiplier = 1024 * 1024 * 1024
			case "tb", "t":
				multiplier = 1024 * 1024 * 1024 * 1024
			case "b", "":
				multiplier = 1
			default:
				return 0, fmt.Errorf("invalid size unit: %s", unit)
			}
			break
		}
	}

	if numPart == "" {
		return 0, fmt.Errorf("invalid size format: no number found")
	}

	size, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid size format: %w", err)
	}

	return uint64(size * float64(multiplier)), nil
}
