package bbbapiwrapper

import (
	"testing"

	"github.com/livekit/protocol/livekit"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

func TestConvertActiveRoomInfoToBBBMeetingInfoWithBreakoutRooms(t *testing.T) {
	room := &plugnmeet.ActiveRoomWithParticipant{
		RoomInfo: &plugnmeet.ActiveRoomInfo{
			RoomTitle: "Main room",
			RoomId:    "room-1",
			Sid:       "sid-1",
		},
		ParticipantsInfo: []*livekit.ParticipantInfo{
			{
				Identity: "user-1",
				Name:     "Alice",
				Tracks:   []*livekit.TrackInfo{{Source: livekit.TrackSource_MICROPHONE}},
			},
		},
		BreakoutRooms: []*plugnmeet.ActiveRoomWithParticipant{
			{
				RoomInfo: &plugnmeet.ActiveRoomInfo{RoomId: "room-1-b1"},
				ParticipantsInfo: []*livekit.ParticipantInfo{
					// duplicate of the main room participant
					{Identity: "user-1", Name: "Alice"},
					{
						Identity: "user-2",
						Name:     "Bob",
						Tracks:   []*livekit.TrackInfo{{Source: livekit.TrackSource_CAMERA}},
					},
				},
			},
		},
	}

	res := ConvertActiveRoomInfoToBBBMeetingInfo(room)

	if res.ParticipantCount != 2 {
		t.Errorf("expected participant count 2, got %d", res.ParticipantCount)
	}
	if !res.HasUserJoined {
		t.Error("expected HasUserJoined to be true")
	}
	if res.VoiceParticipantCount != 1 {
		t.Errorf("expected voice participant count 1, got %d", res.VoiceParticipantCount)
	}
	if res.VideoCount != 1 {
		t.Errorf("expected video count 1, got %d", res.VideoCount)
	}
	if len(res.AttendeesInfo.Attendees) != 2 {
		t.Fatalf("expected 2 attendees, got %d", len(res.AttendeesInfo.Attendees))
	}

	ids := map[string]int{}
	for _, a := range res.AttendeesInfo.Attendees {
		ids[a.UserID]++
	}
	if ids["user-1"] != 1 || ids["user-2"] != 1 {
		t.Errorf("expected deduplicated attendees by identity, got %v", ids)
	}
}
