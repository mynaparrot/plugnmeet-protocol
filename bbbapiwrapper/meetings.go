package bbbapiwrapper

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/encoding/protojson"
)

type MeetingReq struct {
	MeetingID string `query:"meetingID"`
}

type IsMeetingRunningRes struct {
	XMLName    xml.Name `xml:"response"`
	ReturnCode string   `xml:"returncode"`
	Running    bool     `xml:"running"`
}

type MeetingInfo struct {
	XMLName               xml.Name `xml:"meeting"`
	MeetingName           string   `xml:"meetingName"`
	MeetingID             string   `xml:"meetingID"`
	InternalMeetingID     string   `xml:"internalMeetingID"`
	CreateTime            int64    `xml:"createTime"`
	CreateDate            string   `xml:"createDate"`
	AttendeePW            string   `xml:"attendeePW"`  // Deprecated
	ModeratorPW           string   `xml:"moderatorPW"` // Deprecated
	Running               bool     `xml:"running"`
	Duration              uint32   `xml:"duration"`
	HasUserJoined         bool     `xml:"hasUserJoined"`
	Recording             bool     `xml:"recording"`
	HasBeenForciblyEnded  bool     `xml:"hasBeenForciblyEnded"`
	StartTime             int64    `xml:"startTime"`
	EndTime               int64    `xml:"endTime"`
	ParticipantCount      int64    `xml:"participantCount"`
	ListenerCount         int64    `xml:"listenerCount"`
	VoiceParticipantCount int64    `xml:"voiceParticipantCount"`
	VideoCount            int64    `xml:"videoCount"`
	MaxUsers              int64    `xml:"maxUsers"`
	ModeratorCount        int64    `xml:"moderatorCount"`
	AttendeesInfo         struct {
		Attendees []Attendee `xml:"attendee"`
	} `xml:"attendees"`
	IsBreakout bool        `xml:"isBreakout"`
	Metadata   MetadataMap `xml:"metadata"`
}

type Attendee struct {
	XMLName         xml.Name `xml:"attendee"`
	UserID          string   `xml:"userID"`
	FullName        string   `xml:"fullName"`
	Role            string   `xml:"role"`
	IsPresenter     bool     `xml:"isPresenter"`
	IsListeningOnly bool     `xml:"isListeningOnly"`
	HasJoinedVoice  bool     `xml:"hasJoinedVoice"`
	HasVideo        bool     `xml:"hasVideo"`
	ClientType      string   `xml:"clientType"`
}

type MeetingInfoRes struct {
	XMLName     xml.Name     `xml:"response"`
	ReturnCode  string       `xml:"returncode"`
	MeetingInfo *MeetingInfo `xml:",chardata"`
}

type GetMeetingsRes struct {
	XMLName      xml.Name `xml:"response"`
	ReturnCode   string   `xml:"returncode"`
	MeetingsInfo struct {
		Meetings []*MeetingInfo `xml:"meeting"`
	} `xml:"meetings"`
}

func ConvertActiveRoomInfoToBBBMeetingInfo(r *plugnmeet.ActiveRoomWithParticipant) *MeetingInfo {
	rInfo := r.RoomInfo

	res := MeetingInfo{
		MeetingName:           rInfo.RoomTitle,
		MeetingID:             rInfo.RoomId,
		InternalMeetingID:     rInfo.Sid,
		CreateTime:            rInfo.CreationTime * 1000,
		CreateDate:            time.Unix(rInfo.CreationTime, 0).Format(time.RFC1123),
		Running:               true,
		StartTime:             rInfo.CreationTime * 1000,
		Duration:              uint32(time.Now().UTC().Unix() - rInfo.CreationTime),
		ListenerCount:         0,
		VideoCount:            0,
		ModeratorCount:        0,
		VoiceParticipantCount: 0,
	}

	if rInfo.IsRecording > 0 {
		res.Recording = true
	}

	// we'll need to
	rm := new(plugnmeet.RoomMetadata)
	err := protojson.Unmarshal([]byte(rInfo.Metadata), rm)
	if err == nil {
		if rm.ExtraData != nil {
			if data, ok := rm.ExtraData[BbbExtraDataTag]; ok {
				ex := new(CreateMeetingDefaultExtraData)
				err = json.Unmarshal([]byte(data), ex)
				if err == nil {
					res.MeetingID = ex.OriginalMeetingId
					res.ModeratorPW = ex.ModeratorPW
					res.AttendeePW = ex.AttendeePW
					res.Metadata = ex.Meta
				}
			}
		}
	}

	participants := participantsWithBreakoutRooms(r)
	var attendees []Attendee
	for _, p := range participants {
		a := Attendee{
			UserID:   p.Identity,
			FullName: p.Name,
			Role:     "VIEWER",
		}
		if m := new(plugnmeet.UserMetadata); protojson.Unmarshal([]byte(p.Metadata), m) == nil {
			a.IsPresenter = m.IsPresenter
			if m.IsAdmin {
				a.Role = "MODERATOR"
				res.ModeratorCount++
			}
		}

		if len(p.GetTracks()) > 0 {
			for _, t := range p.GetTracks() {
				if t.Source == livekit.TrackSource_CAMERA {
					a.HasVideo = true
					res.VideoCount++
				} else if t.Source == livekit.TrackSource_MICROPHONE {
					a.HasJoinedVoice = true
					res.VoiceParticipantCount++
				}
			}
		} else {
			a.IsListeningOnly = true
			res.ListenerCount++
		}

		attendees = append(attendees, a)
	}

	if len(attendees) > 0 {
		res.AttendeesInfo.Attendees = attendees
	}

	res.ParticipantCount = int64(len(participants))
	res.HasUserJoined = res.ParticipantCount > 0

	return &res
}

// participantsWithBreakoutRooms returns the main room participants plus those of
// its breakout rooms, de-duplicated by identity.
func participantsWithBreakoutRooms(r *plugnmeet.ActiveRoomWithParticipant) []*livekit.ParticipantInfo {
	participants := make([]*livekit.ParticipantInfo, 0, len(r.ParticipantsInfo))
	seen := make(map[string]struct{}, len(r.ParticipantsInfo))

	add := func(p *livekit.ParticipantInfo) {
		if p.Identity != "" {
			if _, ok := seen[p.Identity]; ok {
				return
			}
			seen[p.Identity] = struct{}{}
		}
		participants = append(participants, p)
	}

	for _, p := range r.ParticipantsInfo {
		add(p)
	}
	for _, br := range r.BreakoutRooms {
		for _, p := range br.GetParticipantsInfo() {
			add(p)
		}
	}

	return participants
}
