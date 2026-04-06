package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/google/uuid"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/livekit/protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	authHeader = "Authorization"
	hashToken  = "Hash-Token"
)

type Notifier struct {
	wp     *workerpool.WorkerPool
	logger *logrus.Entry
	ctx    context.Context
	cancel context.CancelFunc
	client *retryablehttp.Client
}

// NewNotifier creates a new Notifier with a specified worker count and max retry attempts.
func NewNotifier(ctx context.Context, workerCount int, maxRetry int, logger *logrus.Entry) *Notifier {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.RetryMax = maxRetry
	// Create custom transport with InsecureSkipVerify
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.HTTPClient.Transport = tr

	ctx, cancel := context.WithCancel(ctx)
	loggerEntry := logger.WithField("component", "webhook-notifier")

	w := &Notifier{
		wp:     workerpool.New(workerCount),
		logger: loggerEntry,
		ctx:    ctx,
		cancel: cancel,
		client: client,
	}

	return w
}

func (n *Notifier) AddInNotifyQueue(event *plugnmeet.CommonNotifyEvent, apiKey, apiSecret string, urls []string) {
	if len(urls) < 1 {
		return
	}

	for _, u := range urls {
		url := u // create a new variable for the closure
		n.wp.Submit(func() {
			l := n.logger.WithFields(logrus.Fields{
				"url":   url,
				"event": event.GetEvent(),
				"room":  event.GetRoom().GetRoomId(),
				"sid":   event.GetRoom().GetSid(),
			})

			statusCode, err := n.sendWebhookRequest(event, apiKey, apiSecret, url, l)
			if err != nil {
				l.WithError(err).Error("Failed to send webhook")
			} else {
				l.WithField("http_status_code", statusCode).Infof("Successfully sent webhook with status code %d", statusCode)
			}
		})
	}
}

// StopGracefully waits for all queued items to be processed before stopping.
func (n *Notifier) StopGracefully() {
	if n.wp != nil {
		n.wp.StopWait()
	}
	n.cancel()
}

// Kill stops the worker immediately, dropping any unprocessed items in the queue.
func (n *Notifier) Kill() {
	if n.wp != nil {
		n.wp.Stop()
	}
	n.cancel()
}

// sendWebhookRequest sends a single webhook event synchronously.
func (n *Notifier) sendWebhookRequest(event *plugnmeet.CommonNotifyEvent, apiKey, apiSecret, url string, l *logrus.Entry) (int, error) {
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
		l.WithError(err).Error("Failed to marshal event")
		return 0, err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		l.WithError(err).Error("Failed to generate token")
		return 0, err
	}

	r, err := retryablehttp.NewRequestWithContext(n.ctx, "POST", url, bytes.NewReader(encoded))
	if err != nil {
		l.WithError(err).Error("Failed to create request")
		return 0, err
	}
	r.Header.Set(authHeader, token)
	r.Header.Set(hashToken, token)
	r.Header.Set("content-type", "application/webhook+json")

	res, err := n.client.Do(r)
	statusCode := 0
	if res != nil {
		statusCode = res.StatusCode
		defer res.Body.Close()
		// Read and discard the body to allow connection reuse, even on error.
		_, _ = io.Copy(io.Discard, res.Body)
	}

	if err != nil {
		return statusCode, err
	}

	return statusCode, nil
}
