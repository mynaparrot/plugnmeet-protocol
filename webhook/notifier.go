package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/livekit/protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	authHeader = "Authorization"
	// in various Apache modules will strip the Authorization header,
	// so we'll use additional one
	hashToken = "Hash-Token"
	maxRetry  = 2 // retryablehttp will use maxRetry + 1
)

// sharedClient is a reusable HTTP client for all webhook requests.
var sharedClient *retryablehttp.Client

func init() {
	sharedClient = retryablehttp.NewClient()
	sharedClient.Logger = nil
	sharedClient.RetryMax = maxRetry
}

type Notifier struct {
	worker *SimpleQueueWorker
	logger *logrus.Entry
}

func NewNotifier(ctx context.Context, queueSize int, logger *logrus.Logger) *Notifier {
	loggerEntry := logger.WithField("component", "webhook-notifier")
	w := &Notifier{
		logger: loggerEntry,
		worker: NewSimpleQueueWorker(ctx, queueSize, loggerEntry.WithField("sub-component", "queue-worker")),
	}

	return w
}

func (n *Notifier) AddInNotifyQueue(event *plugnmeet.CommonNotifyEvent, apiKey, apiSecret string, urls []string) {
	if len(urls) < 1 {
		return
	}

	for _, u := range urls {
		n.worker.Submit(func() {
			res, err := n.sendWebhookRequest(event, apiKey, apiSecret, u)
			logFields := logrus.Fields{
				"url":   u,
				"event": event.GetEvent(),
				"room":  event.GetRoom().GetRoomId(),
				"sid":   event.GetRoom().GetSid(),
			}

			if err != nil {
				n.logger.WithFields(logFields).WithError(err).Error("failed to send webhook")
			} else if res != nil {
				defer res.Body.Close()
				logFields["http_status_code"] = res.StatusCode
				logFields["http_status"] = res.Status
				n.logger.WithFields(logFields).Info("webhook sent successfully")
			}
		})
	}
}

// StopGracefully waits for all queued items to be processed before stopping.
func (n *Notifier) StopGracefully() {
	if n.worker != nil {
		n.worker.StopGracefully()
	}
}

// Kill stops the worker immediately, dropping any unprocessed items in the queue.
func (n *Notifier) Kill() {
	if n.worker != nil {
		n.worker.Kill()
	}
}

// sendWebhookRequest sends a single webhook event synchronously.
func (n *Notifier) sendWebhookRequest(event *plugnmeet.CommonNotifyEvent, apiKey, apiSecret, url string) (*http.Response, error) {
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
		return nil, err
	}
	r.Header.Set(authHeader, token)
	r.Header.Set(hashToken, token)
	r.Header.Set("content-type", "application/webhook+json")
	res, err := sharedClient.Do(r)

	if err != nil {
		return nil, err
	}
	// The caller is responsible for closing the response body.
	return res, nil
}
