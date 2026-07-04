package utils

import (
	"net/url"
	"strconv"

	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

func AssignLTIV1CustomParams(params *url.Values, claims *plugnmeet.LtiClaims) {
	customPara := new(plugnmeet.LtiCustomParameters)

	if params.Get("custom_room_duration") != "" {
		duration, _ := strconv.Atoi(params.Get("custom_room_duration"))
		customPara.RoomDuration = new(uint64(duration))
	}

	// helper function to safely parse boolean params
	parseBool := func(param string) *bool {
		if val := params.Get(param); val != "" {
			b, err := strconv.ParseBool(val)
			if err == nil {
				return &b
			}
		}
		return nil
	}

	if val := parseBool("custom_allow_polls"); val != nil && !*val {
		customPara.AllowPolls = val
	}
	if val := parseBool("custom_allow_shared_note_pad"); val != nil && !*val {
		customPara.AllowSharedNotePad = val
	}
	if val := parseBool("custom_allow_breakout_room"); val != nil && !*val {
		customPara.AllowBreakoutRoom = val
	}
	if val := parseBool("custom_allow_recording"); val != nil && !*val {
		customPara.AllowRecording = val
	}
	if val := parseBool("custom_allow_rtmp"); val != nil && !*val {
		customPara.AllowRtmp = val
	}
	if val := parseBool("custom_allow_view_other_webcams"); val != nil && !*val {
		customPara.AllowViewOtherWebcams = val
	}
	if val := parseBool("custom_allow_view_other_users_list"); val != nil && !*val {
		customPara.AllowViewOtherUsersList = val
	}
	if val := parseBool("custom_mute_on_start"); val != nil && *val {
		customPara.MuteOnStart = val
	}

	// custom design
	customDesign := new(plugnmeet.LtiCustomDesign)
	if params.Get("custom_primary_color") != "" {
		customDesign.PrimaryColor = new(params.Get("custom_primary_color"))
	}
	if params.Get("custom_secondary_color") != "" {
		customDesign.SecondaryColor = new(params.Get("custom_secondary_color"))
	}
	if params.Get("custom_background_color") != "" {
		customDesign.BackgroundColor = new(params.Get("custom_background_color"))
	}
	if params.Get("custom_custom_logo") != "" {
		customDesign.CustomLogo = new(params.Get("custom_custom_logo"))
	}

	claims.LtiCustomParameters = customPara
	claims.LtiCustomParameters.LtiCustomDesign = customDesign
}

func PrepareLTIV1RoomCreateReq(c *plugnmeet.LtiClaims) *plugnmeet.CreateRoomReq {
	// minimum initialization
	req := &plugnmeet.CreateRoomReq{
		RoomId: c.RoomId,
		Metadata: &plugnmeet.RoomMetadata{
			RoomTitle: c.RoomTitle,
		},
	}

	// set all default room features
	PrepareDefaultRoomFeatures(req)

	if c.LtiCustomParameters != nil {
		p := c.LtiCustomParameters
		f := req.Metadata.RoomFeatures

		if p.RoomDuration != nil && *p.RoomDuration > 0 {
			f.RoomDuration = p.RoomDuration
		}
		if p.MuteOnStart != nil {
			f.MuteOnStart = *p.MuteOnStart
		}
		if p.AllowSharedNotePad != nil {
			f.SharedNotePadFeatures.IsAllow = *p.AllowSharedNotePad
		}
		if p.AllowBreakoutRoom != nil {
			f.BreakoutRoomFeatures.IsAllow = *p.AllowBreakoutRoom
		}
		if p.AllowPolls != nil {
			f.PollsFeatures.IsAllow = *p.AllowPolls
		}
		if p.AllowRecording != nil {
			f.RecordingFeatures.IsAllow = *p.AllowRecording
		}
		if p.AllowRtmp != nil {
			f.ExternalBroadcastingFeatures.IsAllow = *p.AllowRtmp
			f.ExternalBroadcastingFeatures.IsAllowRtmp = *p.AllowRtmp
		}
		if p.AllowViewOtherWebcams != nil {
			f.AllowViewOtherWebcams = *p.AllowViewOtherWebcams
		}
		if p.AllowViewOtherUsersList != nil {
			f.AllowViewOtherUsersList = *p.AllowViewOtherUsersList
		}

		req.Metadata.RoomFeatures = f
	}

	return req
}
