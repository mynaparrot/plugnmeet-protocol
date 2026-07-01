package utils

import (
	"time"

	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/proto"
)

// PrepareDefaultRoomFeatures will initialize all the features and lock settings with default values
// this ensure that none of the feature has nil values
func PrepareDefaultRoomFeatures(r *plugnmeet.CreateRoomReq) {
	if r.Metadata == nil {
		r.Metadata = new(plugnmeet.RoomMetadata)
	}
	if r.Metadata.DefaultLockSettings == nil {
		r.Metadata.DefaultLockSettings = proto.Clone(defaultLockSettings).(*plugnmeet.LockSettings)
	}
	if r.Metadata.ExtraData == nil {
		r.Metadata.ExtraData = make(map[string]string)
	}
	if r.Metadata.CopyrightConf == nil {
		r.Metadata.CopyrightConf = proto.Clone(defaultCopyrightConf).(*plugnmeet.CopyrightConf)
	}

	// now with all room features
	if r.Metadata.RoomFeatures == nil {
		r.Metadata.RoomFeatures = &plugnmeet.RoomCreateFeatures{
			AllowWebcams:            true,
			MuteOnStart:             false,
			AllowScreenShare:        true,
			AllowViewOtherWebcams:   true,
			AllowViewOtherUsersList: true,
			AdminOnlyWebcams:        false,
			EnableAnalytics:         true,
			AllowVirtualBg:          new(true),
			AllowRaiseHand:          new(true),
			AllowReactions:          new(false),
		}
	}
	rf := r.Metadata.RoomFeatures

	if rf.RecordingFeatures == nil {
		rf.RecordingFeatures = proto.Clone(defaultRecordingFeatures).(*plugnmeet.RecordingFeatures)
	}
	if !rf.RecordingFeatures.IsAllow {
		rf.RecordingFeatures.IsAllowCloud = false
		rf.RecordingFeatures.IsAllowLocal = false
		rf.RecordingFeatures.EnableAutoCloudRecording = false
	}

	if rf.ChatFeatures == nil {
		rf.ChatFeatures = proto.Clone(defaultChatFeatures).(*plugnmeet.ChatFeatures)
	}
	if !rf.ChatFeatures.IsAllow {
		rf.ChatFeatures.IsAllowFileUpload = false
	}

	if rf.SharedNotePadFeatures == nil {
		rf.SharedNotePadFeatures = proto.Clone(defaultSharedNotePadFeatures).(*plugnmeet.SharedNotePadFeatures)
	}

	if rf.WhiteboardFeatures == nil {
		rf.WhiteboardFeatures = proto.Clone(defaultWhiteboardFeatures).(*plugnmeet.WhiteboardFeatures)
	}

	if rf.ExternalMediaPlayerFeatures == nil {
		rf.ExternalMediaPlayerFeatures = proto.Clone(defaultExternalMediaPlayerFeatures).(*plugnmeet.ExternalMediaPlayerFeatures)
	}

	if rf.WaitingRoomFeatures == nil {
		rf.WaitingRoomFeatures = proto.Clone(defaultWaitingRoomFeatures).(*plugnmeet.WaitingRoomFeatures)
	}

	if rf.BreakoutRoomFeatures == nil {
		rf.BreakoutRoomFeatures = proto.Clone(defaultBreakoutRoomFeatures).(*plugnmeet.BreakoutRoomFeatures)
	}

	if rf.DisplayExternalLinkFeatures == nil {
		rf.DisplayExternalLinkFeatures = proto.Clone(defaultDisplayExternalLinkFeatures).(*plugnmeet.DisplayExternalLinkFeatures)
	}

	if rf.IngressFeatures == nil {
		rf.IngressFeatures = proto.Clone(defaultIngressFeatures).(*plugnmeet.IngressFeatures)
	}

	if rf.EndToEndEncryptionFeatures == nil {
		rf.EndToEndEncryptionFeatures = proto.Clone(defaultEndToEndEncryptionFeatures).(*plugnmeet.EndToEndEncryptionFeatures)
	}
	if !rf.EndToEndEncryptionFeatures.IsEnabled {
		rf.EndToEndEncryptionFeatures.EnabledSelfInsertEncryptionKey = false
		rf.EndToEndEncryptionFeatures.IncludedChatMessages = false
		rf.EndToEndEncryptionFeatures.IncludedWhiteboard = false
	}

	if rf.PollsFeatures == nil {
		rf.PollsFeatures = proto.Clone(defaultPollsFeatures).(*plugnmeet.PollsFeatures)
	}

	if rf.InsightsFeatures == nil {
		rf.InsightsFeatures = proto.Clone(defaultInsightsFeatures).(*plugnmeet.InsightsFeatures)
	}

	if rf.SipDialInFeatures == nil {
		rf.SipDialInFeatures = proto.Clone(defaultSipDialInFeatures).(*plugnmeet.SipDialInFeatures)
	}

	if rf.ExternalBroadcastingFeatures == nil {
		rf.ExternalBroadcastingFeatures = proto.Clone(defaultExternalBroadcastingFeatures).(*plugnmeet.ExternalBroadcastingFeatures)
	}
	if rf.AllowRtmp != nil {
		rf.ExternalBroadcastingFeatures.IsAllow = *rf.AllowRtmp
		rf.ExternalBroadcastingFeatures.IsAllowRtmp = *rf.AllowRtmp
	}
	if !rf.ExternalBroadcastingFeatures.IsAllow {
		rf.ExternalBroadcastingFeatures.IsAllowRtmp = false
	}

	r.Metadata.StartedAt = uint64(time.Now().UTC().Unix())
	r.Metadata.RoomFeatures = rf
}

func SetCreateRoomDefaultValues(r *plugnmeet.CreateRoomReq, maxSize, maxSizeWhiteboardFile uint64, allowedTypes []string, allowedNotepad bool) {
	rf := r.Metadata.RoomFeatures

	if rf.AutoGenUserId == nil {
		// by default, auto user id generation will be disabled
		rf.AutoGenUserId = new(false)
	}

	if rf.SharedNotePadFeatures.IsAllow && !allowedNotepad {
		rf.SharedNotePadFeatures.IsAllow = false
	}

	if rf.ChatFeatures.IsAllowFileUpload {
		if len(rf.ChatFeatures.AllowedFileTypes) == 0 {
			rf.ChatFeatures.AllowedFileTypes = allowedTypes
		}
		rf.ChatFeatures.MaxFileSize = &maxSize
	}

	if rf.WhiteboardFeatures.IsAllow {
		if maxSizeWhiteboardFile < 1 {
			maxSizeWhiteboardFile = 30
		}
		rf.WhiteboardFeatures.MaxAllowedFileSize = &maxSizeWhiteboardFile
	}

	if rf.BreakoutRoomFeatures.IsAllow && rf.BreakoutRoomFeatures.AllowedNumberRooms == 0 {
		rf.BreakoutRoomFeatures.AllowedNumberRooms = 6
	}

	if rf.EndToEndEncryptionFeatures.IsEnabled {
		r.Metadata.RoomFeatures.SipDialInFeatures.IsAllow = false
		r.Metadata.RoomFeatures.IngressFeatures.IsAllow = false

		if !rf.EndToEndEncryptionFeatures.EnabledSelfInsertEncryptionKey {
			randomKey, err := GenerateSecureRandomStrings(32)
			if err != nil {
				randomKey = GenerateRandomStrings(32)
			}
			rf.EndToEndEncryptionFeatures.EncryptionKey = &randomKey
		} else {
			// if self insert key enabled
			r.Metadata.RoomFeatures.ExternalBroadcastingFeatures.IsAllow = false
			r.Metadata.RoomFeatures.RecordingFeatures.IsAllowCloud = false
			r.Metadata.RoomFeatures.RecordingFeatures.EnableAutoCloudRecording = false

			insightsFeatures := r.Metadata.RoomFeatures.InsightsFeatures
			if insightsFeatures.TranscriptionFeatures != nil {
				insightsFeatures.TranscriptionFeatures.IsAllow = false
			}
			if insightsFeatures.AiFeatures != nil && insightsFeatures.AiFeatures.MeetingSummarizationFeatures != nil {
				insightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow = false
			}
		}
	}

	if !rf.InsightsFeatures.IsAllow {
		if rf.InsightsFeatures.TranscriptionFeatures != nil {
			rf.InsightsFeatures.TranscriptionFeatures.IsAllow = false
		}
		if rf.InsightsFeatures.ChatTranslationFeatures != nil {
			rf.InsightsFeatures.ChatTranslationFeatures.IsAllow = false
		}
		if rf.InsightsFeatures.AiFeatures != nil {
			rf.InsightsFeatures.AiFeatures.IsAllow = false
		}
	}

	if rf.InsightsFeatures.TranscriptionFeatures != nil && !rf.InsightsFeatures.TranscriptionFeatures.IsAllow {
		rf.InsightsFeatures.TranscriptionFeatures.IsAllowTranslation = false
	}

	if rf.InsightsFeatures.AiFeatures != nil && !rf.InsightsFeatures.AiFeatures.IsAllow {
		if rf.InsightsFeatures.AiFeatures.AiTextChatFeatures != nil {
			rf.InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow = false
		}
		if rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures != nil {
			rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow = false
		}
	}
}

func SetRoomDefaultLockSettings(r *plugnmeet.CreateRoomReq) {
	if r.Metadata.DefaultLockSettings.LockScreenSharing == nil {
		r.Metadata.DefaultLockSettings.LockScreenSharing = new(true)
	}
	if r.Metadata.DefaultLockSettings.LockWhiteboard == nil {
		r.Metadata.DefaultLockSettings.LockWhiteboard = new(true)
	}
	if r.Metadata.DefaultLockSettings.LockSharedNotepad == nil {
		r.Metadata.DefaultLockSettings.LockSharedNotepad = new(true)
	}
}

type RoomDefaultSettings struct {
	MaxParticipants                    *uint32 `yaml:"max_participants"`
	MaxDuration                        *uint64 `yaml:"max_duration"`
	MaxNumBreakoutRooms                *uint32 `yaml:"max_num_breakout_rooms"`
	MaxPreloadedWhiteboardFileSize     *string `yaml:"max_preloaded_wb_file_size"`
	MaxPreloadedWhiteboardFileSizeByte *uint64 `yaml:"-"`
}

func SetDefaultRoomSettings(s *RoomDefaultSettings, r *plugnmeet.CreateRoomReq) {
	if s == nil {
		return
	}

	if s.MaxParticipants != nil && *s.MaxParticipants > 0 {
		if r.MaxParticipants != nil {
			if *r.MaxParticipants == 0 || *r.MaxParticipants > *s.MaxParticipants {
				r.MaxParticipants = s.MaxParticipants
			}
		} else {
			r.MaxParticipants = s.MaxParticipants
		}
	}

	if s.MaxDuration != nil && *s.MaxDuration > 0 {
		if r.Metadata.RoomFeatures.RoomDuration != nil {
			if *r.Metadata.RoomFeatures.RoomDuration == 0 || *r.Metadata.RoomFeatures.RoomDuration > *s.MaxDuration {
				r.Metadata.RoomFeatures.RoomDuration = s.MaxDuration
			}
		} else {
			r.Metadata.RoomFeatures.RoomDuration = s.MaxDuration
		}
	}

	if r.EmptyTimeout == nil || *r.EmptyTimeout < 120 {
		var et uint32 = 1800 // 1800 seconds = 30 minutes
		r.EmptyTimeout = &et
	}

	// at present, we will allow max 16 rooms
	var maxNum uint32 = 16
	if s.MaxNumBreakoutRooms == nil {
		s.MaxNumBreakoutRooms = &maxNum
	} else if s.MaxNumBreakoutRooms != nil && *s.MaxParticipants > maxNum {
		s.MaxNumBreakoutRooms = &maxNum
	}

	if r.Metadata.RoomFeatures.BreakoutRoomFeatures.AllowedNumberRooms > *s.MaxNumBreakoutRooms {
		r.Metadata.RoomFeatures.BreakoutRoomFeatures.AllowedNumberRooms = *s.MaxNumBreakoutRooms
	}
}
