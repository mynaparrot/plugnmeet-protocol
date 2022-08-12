package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livekit/protocol/livekit"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/proto"
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

func SendCommonResponse(c *fiber.Ctx, s bool, m string) error {
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

func SendProtoResponse(c *fiber.Ctx, res proto.Message) error {
	marshal, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/protobuf")
	return c.Send(marshal)
}
