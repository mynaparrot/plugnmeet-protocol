package bbbapiwrapper

import (
	"encoding/json"
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/mynaparrot/plugnmeet-protocol/utils"
)

type CreateMeetingReq struct {
	Name                    string `query:"name"`
	MeetingID               string `query:"meetingID"`
	AttendeePW              string `query:"attendeePW"`  // Deprecated
	ModeratorPW             string `query:"moderatorPW"` // Deprecated
	Welcome                 string `query:"welcome"`
	MaxParticipants         uint32 `query:"maxParticipants"`
	LogoutURL               string `query:"logoutURL"`
	Duration                uint64 `query:"duration"`
	Record                  bool   `query:"record"`
	AutoStartRecording      bool   `query:"autoStartRecording"`
	WebcamsOnlyForModerator bool   `query:"webcamsOnlyForModerator"`
	MuteOnStart             bool   `query:"muteOnStart"`
	GuestPolicy             string `query:"guestPolicy"` // ALWAYS_ACCEPT, ASK_MODERATOR
	MeetingKeepEvents       bool   `query:"meetingKeepEvents"`
	Logo                    string `query:"logo"`
	DisabledFeatures        string `query:"disabledFeatures"` //breakoutRooms,chat,externalVideos,polls,screenshare,sharedNotes,virtualBackgrounds,liveTranscription,presentation,virtualBackgrounds,raiseHand
	PreUploadedPresentation string `query:"preUploadedPresentation"`

	// few locks
	LockSettingsDisableCam         bool `query:"lockSettingsDisableCam"`
	LockSettingsDisableMic         bool `query:"lockSettingsDisableMic"`
	LockSettingsDisablePrivateChat bool `query:"lockSettingsDisablePrivateChat"`
	LockSettingsDisablePublicChat  bool `query:"lockSettingsDisablePublicChat"`
	LockSettingsDisableNotes       bool `query:"lockSettingsDisableNotes"`
	LockSettingsHideUserList       bool `query:"lockSettingsHideUserList"`

	// to avoid incompatibility
	VoiceBridge string `query:"voiceBridge"`
	DialNumber  string `query:"dialNumber"`
}

const BbbExtraDataTag = "bbb_extra_data"

type CreateMeetingDefaultExtraData struct {
	AttendeePW        string            `json:"attendeePW"`
	ModeratorPW       string            `json:"moderatorPW"`
	Logo              string            `json:"logo"`
	OriginalMeetingId string            `json:"originalMeetingId"`
	Meta              map[string]string `json:"meta"`
}

type CreateMeetingResp struct {
	XMLName              xml.Name `xml:"response"`
	ReturnCode           string   `xml:"returncode"`
	MessageKey           string   `xml:"messageKey"`
	Message              string   `xml:"message"`
	MeetingID            string   `xml:"meetingID"`
	InternalMeetingID    string   `xml:"internalMeetingID"`
	ParentMeetingID      string   `xml:"parentMeetingID"`
	AttendeePW           string   `xml:"attendeePW"`  // Deprecated
	ModeratorPW          string   `xml:"moderatorPW"` // Deprecated
	CreateTime           int64    `xml:"createTime"`
	CreateDate           string   `xml:"createDate"`
	HasUserJoined        bool     `xml:"hasUserJoined"`
	Duration             int64    `xml:"duration"`
	VoiceBridge          string   `xml:"voiceBridge"`
	DialNumber           string   `xml:"dialNumber"`
	HasBeenForciblyEnded bool     `xml:"hasBeenForciblyEnded"`
}

type PreUploadWhiteboardPostFile struct {
	XMLName xml.Name `xml:"modules"`
	Module  struct {
		Name      string `xml:"name,attr"`
		Documents []struct {
			URL      string `xml:"url,attr"`
			Filename string `xml:"filename,attr"`
			Name     string `xml:"name,attr"`
		} `xml:"document"`
	} `xml:"module"`
}

func ConvertCreateRequest(r *CreateMeetingReq, rawQueries map[string]string) (*plugnmeet.CreateRoomReq, error) {
	req := plugnmeet.CreateRoomReq{
		RoomId: CheckMeetingIdToMatchFormat(r.MeetingID),
		Metadata: &plugnmeet.RoomMetadata{
			RoomTitle:    r.Name,
			RoomFeatures: new(plugnmeet.RoomCreateFeatures),
		},
	}
	// set defaults
	utils.PrepareDefaultRoomFeatures(&req)
	// now we can customize features
	setFeatures(r, req.Metadata.RoomFeatures)

	if r.MaxParticipants > 0 {
		req.MaxParticipants = &r.MaxParticipants
	}
	if r.Duration > 0 {
		req.Metadata.RoomFeatures.RoomDuration = &r.Duration
	}
	if r.LogoutURL != "" {
		req.Metadata.LogoutUrl = &r.LogoutURL
	}

	if r.Welcome != "" {
		r.Welcome = strings.ReplaceAll(r.Welcome, "%%CONFNAME%%", r.Name)
		req.Metadata.WelcomeMessage = &r.Welcome
	}

	if r.GuestPolicy != "" {
		if r.GuestPolicy == "ASK_MODERATOR" {
			req.Metadata.RoomFeatures.WaitingRoomFeatures = &plugnmeet.WaitingRoomFeatures{
				IsActive: true,
			}
		}
	}

	if r.DisabledFeatures != "" {
		setDifferentFeatures(req.Metadata.RoomFeatures, r.DisabledFeatures)
	}

	if r.PreUploadedPresentation != "" && req.Metadata.RoomFeatures.WhiteboardFeatures.IsAllow {
		req.Metadata.RoomFeatures.WhiteboardFeatures.PreloadFile = &r.PreUploadedPresentation
	}

	// we'll only consider if some value was sent
	if rawQueries["meetingKeepEvents"] != "" {
		req.Metadata.RoomFeatures.EnableAnalytics = r.MeetingKeepEvents
	}

	// lock settings
	if r.LockSettingsHideUserList {
		req.Metadata.RoomFeatures.AllowViewOtherUsersList = false
	}
	// now adjust lock settings
	setLockSettings(req.Metadata.DefaultLockSettings, r)

	meta := map[string]string{}
	for k, qq := range rawQueries {
		if strings.Contains(k, "meta_") {
			val := qq
			unescape, err := url.QueryUnescape(qq)
			if err == nil {
				val = unescape
			}
			meta[strings.Replace(k, "meta_", "", 1)] = val
		}
	}

	marshal, err := json.Marshal(CreateMeetingDefaultExtraData{
		ModeratorPW:       r.ModeratorPW,
		AttendeePW:        r.AttendeePW,
		OriginalMeetingId: r.MeetingID,
		Logo:              r.Logo,
		Meta:              meta,
	})
	if err != nil {
		return nil, err
	}
	req.Metadata.ExtraData = map[string]string{
		BbbExtraDataTag: string(marshal),
	}

	return &req, nil
}

func setFeatures(r *CreateMeetingReq, f *plugnmeet.RoomCreateFeatures) {
	b := true
	f.AllowWebcams = true
	f.AdminOnlyWebcams = r.WebcamsOnlyForModerator
	f.EnableAnalytics = true
	f.MuteOnStart = r.MuteOnStart
	f.AllowScreenShare = true
	f.AllowRaiseHand = &b
	f.AllowVirtualBg = &b
	f.AutoGenUserId = &b

	f.RecordingFeatures.IsAllow = r.Record
	f.RecordingFeatures.IsAllowCloud = r.Record
	f.RecordingFeatures.EnableAutoCloudRecording = r.AutoStartRecording
}

func setDifferentFeatures(f *plugnmeet.RoomCreateFeatures, disabledFeatures string) {
	features := strings.Split(disabledFeatures, ",")
	fVal := false

	for _, ff := range features {
		switch ff {
		case "breakoutRooms":
			f.BreakoutRoomFeatures.IsAllow = fVal
		case "chat":
			f.ChatFeatures.IsAllow = fVal
		case "externalVideos":
			f.ExternalMediaPlayerFeatures.IsAllow = fVal
		case "polls":
			f.PollsFeatures.IsAllow = fVal
		case "screenshare":
			f.AllowScreenShare = fVal
		case "sharedNotes":
			f.SharedNotePadFeatures.IsAllow = fVal
		case "liveTranscription":
			f.InsightsFeatures.TranscriptionFeatures.IsAllow = fVal
			f.InsightsFeatures.ChatTranslationFeatures.IsAllow = fVal
		case "presentation":
			f.WhiteboardFeatures.IsAllow = fVal
		case "virtualBackgrounds":
			f.AllowVirtualBg = &fVal
		case "raiseHand":
			f.AllowRaiseHand = &fVal
		}
	}
}

func setLockSettings(f *plugnmeet.LockSettings, r *CreateMeetingReq) {
	l := true

	if r.LockSettingsDisableCam {
		f.LockWebcam = &l
	}
	if r.LockSettingsDisableMic {
		f.LockMicrophone = &l
	}
	if r.LockSettingsDisablePublicChat {
		f.LockChatSendMessage = &l
	}
	if r.LockSettingsDisablePrivateChat {
		f.LockPrivateChat = &l
	}
	if r.LockSettingsDisableNotes {
		f.LockSharedNotepad = &l
	}
}
