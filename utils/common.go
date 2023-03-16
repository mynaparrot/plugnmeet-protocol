package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livekit/protocol/livekit"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/fs"
	"path/filepath"
	"sort"
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

func SendCommonProtobufResponse(c *fiber.Ctx, s bool, m string) error {
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

func SendProtobufResponse(c *fiber.Ctx, res proto.Message) error {
	marshal, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/protobuf")
	return c.Send(marshal)
}

func SendCommonProtoJsonResponse(c *fiber.Ctx, s bool, m string) error {
	res := &plugnmeet.CommonResponse{
		Status: s,
		Msg:    m,
	}
	op := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
	marshal, err := op.Marshal(res)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/json")
	return c.Send(marshal)
}

func SendProtoJsonResponse(c *fiber.Ctx, res proto.Message) error {
	op := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
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
		if filepath.Ext(d.Name()) == ext {
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
