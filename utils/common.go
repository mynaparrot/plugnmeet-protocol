package utils

import (
	"github.com/livekit/protocol/livekit"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

func PrepareCommonWebhookNotifyEvent(event *livekit.WebhookEvent) *plugnmeet.CommonNotifyEvent {
	ct := uint64(event.Room.CreationTime)
	return &plugnmeet.CommonNotifyEvent{
		Event: &event.Event,
		Room: &plugnmeet.NotifyEventRoom{
			Sid:             &event.Room.Sid,
			RoomId:          &event.Room.Name,
			EmptyTimeout:    &event.Room.EmptyTimeout,
			MaxParticipants: &event.Room.MaxParticipants,
			CreationTime:    &ct,
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
