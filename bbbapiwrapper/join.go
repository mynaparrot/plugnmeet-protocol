package bbbapiwrapper

import (
	"encoding/xml"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

type JoinMeetingReq struct {
	FullName  string `query:"fullName"`
	MeetingID string `query:"meetingID"`
	Password  string `query:"password"` // Deprecated
	Role      string `query:"role"`     // MODERATOR or VIEWER
	UserID    string `query:"userID"`
	AvatarURL string `query:"avatarURL"`
}

type JoinMeetingRes struct {
	XMLName    xml.Name `xml:"response"`
	ReturnCode string   `xml:"returncode"`
	MessageKey string   `xml:"messageKey"`
	Message    string   `xml:"message"`
}

func ConvertJoinRequest(r *JoinMeetingReq, isAdmin bool) *plugnmeet.GenerateTokenReq {
	req := plugnmeet.GenerateTokenReq{
		RoomId: CheckMeetingIdToMatchFormat(r.MeetingID),
		UserInfo: &plugnmeet.UserInfo{
			Name:    r.FullName,
			UserId:  CheckForUserId(r.UserID),
			IsAdmin: isAdmin,
		},
	}

	if r.AvatarURL != "" {
		req.UserInfo.UserMetadata.ProfilePic = &r.AvatarURL
	}

	return &req
}
