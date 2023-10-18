package webhook

import (
	"context"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const maxDurationWaitBeforeCleanRoom = 2 // minutes

type WebhookNotifier struct {
	sync.RWMutex
	ctx                   context.Context
	webhookQueuedNotifier map[string]QueuedNotifier

	apiKey    string
	apiSecret string
	urls      []string
	logger    *logrus.Logger
}

func NewWebhookNotifier(apiKey, apiSecret string, logger *logrus.Logger) *WebhookNotifier {
	return &WebhookNotifier{
		ctx:                   context.Background(),
		webhookQueuedNotifier: make(map[string]QueuedNotifier),
		apiKey:                apiKey,
		apiSecret:             apiSecret,
		logger:                logger,
	}
}

func (a *WebhookNotifier) AddToWebhookQueuedNotifier(roomId string, urls []string) {
	if a == nil || len(urls) == 0 || roomId == "" {
		return
	}

	a.Lock()
	defer a.Unlock()
	if _, ok := a.webhookQueuedNotifier[roomId]; !ok {
		a.webhookQueuedNotifier[roomId] = NewDefaultNotifier(a.apiKey, a.apiSecret, urls, a.logger)
	}
}

func (a *WebhookNotifier) RoomExist(roomId string) bool {
	if a == nil || roomId == "" {
		return false
	}

	a.Lock()
	defer a.Unlock()
	_, ok := a.webhookQueuedNotifier[roomId]
	return ok
}

func (a *WebhookNotifier) SendWebhook(event *plugnmeet.CommonNotifyEvent) error {
	if a == nil || event == nil || event.GetRoom().GetRoomId() == "" {
		return nil
	}
	a.Lock()
	defer a.Unlock()
	if queued, ok := a.webhookQueuedNotifier[event.GetRoom().GetRoomId()]; ok {
		err := queued.QueueNotify(a.ctx, event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *WebhookNotifier) DeleteWebhookQueuedNotifier(roomId string) {
	if a == nil || roomId == "" {
		return
	}
	// we'll wait long time before delete WebhookQueued
	// to make sure that we've completed everything else
	time.Sleep(time.Minute * maxDurationWaitBeforeCleanRoom)
	a.Lock()
	defer a.Unlock()
	if _, ok := a.webhookQueuedNotifier[roomId]; ok {
		delete(a.webhookQueuedNotifier, roomId)
	}
}
