package utils

import (
	"time"

	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// mergeRoomFeatures clones the default features and applies the user's
// explicitly set fields on top. Unlike proto.Merge, non-nil message fields
// replace the default entirely, preserving their internal false values.
func mergeRoomFeatures(user *plugnmeet.RoomCreateFeatures) *plugnmeet.RoomCreateFeatures {
	result := proto.Clone(defaultRoomFeatures).(*plugnmeet.RoomCreateFeatures)
	if user == nil {
		return result
	}

	userFields := user.ProtoReflect()
	resultFields := result.ProtoReflect()

	userFields.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Kind() == protoreflect.MessageKind && !fd.IsList() && !fd.IsMap() {
			// Clone message fields so that:
			// - we don't alias the caller's structs
			// - the user's explicit proto3 zero values (e.g. IsAllowCloud: false)
			//   survive, which Range() would never report
			cloned := proto.Clone(v.Message().Interface()).ProtoReflect()
			// nested message fields the user left unset must never stay nil
			fillUnsetMessageFields(cloned, resultFields.Get(fd).Message())
			resultFields.Set(fd, protoreflect.ValueOf(cloned))
		} else {
			resultFields.Set(fd, v)
		}
		return true
	})

	return result
}

// fillUnsetMessageFields recursively makes sure that every nested message field
// of msg is populated. Fields which the user has set are processed recursively
// (using the matching default sub-message), while fields the user left unset are
// taken from defaults. If defaults has no value for a field, a new empty message
// is created so that no nested message field stays nil.
func fillUnsetMessageFields(msg, defaults protoreflect.Message) {
	fields := msg.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if fd.Kind() != protoreflect.MessageKind || fd.IsList() || fd.IsMap() {
			continue
		}

		var defSub protoreflect.Message
		if defaults != nil && defaults.IsValid() && defaults.Has(fd) {
			defSub = defaults.Get(fd).Message()
		}

		if msg.Has(fd) {
			// set by the user: keep its explicit zero values, but fill its own
			// nested message fields
			fillUnsetMessageFields(msg.Get(fd).Message(), defSub)
			continue
		}

		if defSub != nil {
			cloned := proto.Clone(defSub.Interface()).ProtoReflect()
			msg.Set(fd, protoreflect.ValueOf(cloned))
			continue
		}

		if fd.ContainingOneof() != nil {
			// optional/oneof message field which neither the user nor the defaults
			// set: keep it unset so we don't change its presence semantics
			continue
		}

		msg.Set(fd, msg.NewField(fd))
	}
}

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
	// mergeRoomFeatures replaces if nil with default.
	r.Metadata.RoomFeatures = mergeRoomFeatures(r.Metadata.RoomFeatures)
	rf := r.Metadata.RoomFeatures

	// reset everything if disabled
	if !rf.RecordingFeatures.IsAllow {
		rf.RecordingFeatures = &plugnmeet.RecordingFeatures{}
	}
	if !rf.ChatFeatures.IsAllow {
		rf.ChatFeatures = &plugnmeet.ChatFeatures{}
	}
	if !rf.EndToEndEncryptionFeatures.IsEnabled {
		rf.EndToEndEncryptionFeatures = &plugnmeet.EndToEndEncryptionFeatures{}
	}
	if !rf.ExternalBroadcastingFeatures.IsAllow {
		rf.ExternalBroadcastingFeatures.IsAllowRtmp = false
	}

	r.Metadata.StartedAt = uint64(time.Now().UTC().Unix())
}

func SetCreateRoomDefaultValues(r *plugnmeet.CreateRoomReq, maxSize, maxSizeWhiteboardFile uint64, allowedTypes []string) {
	rf := r.Metadata.RoomFeatures

	if rf.AutoGenUserId == nil {
		// by default, auto user id generation will be disabled
		rf.AutoGenUserId = new(false)
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

		if rf.EndToEndEncryptionFeatures.EnabledSelfInsertEncryptionKey {
			// if self insert key enabled
			r.Metadata.RoomFeatures.ExternalBroadcastingFeatures.IsAllow = false
			r.Metadata.RoomFeatures.RecordingFeatures.IsAllowCloud = false
			r.Metadata.RoomFeatures.RecordingFeatures.EnableAutoCloudRecording = false

			// nested insights features are always non-nil after PrepareDefaultRoomFeatures
			insightsFeatures := r.Metadata.RoomFeatures.InsightsFeatures
			insightsFeatures.TranscriptionFeatures.IsAllow = false
			insightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow = false
		} else {
			// only generate key if nothing was delivered
			if rf.EndToEndEncryptionFeatures.EncryptionKey == nil ||
				*rf.EndToEndEncryptionFeatures.EncryptionKey == "" {
				randomKey, err := GenerateSecureRandomStrings(32)
				if err != nil {
					randomKey = GenerateRandomStrings(32)
				}
				rf.EndToEndEncryptionFeatures.EncryptionKey = &randomKey
			}
		}
	}

	// nested insights features are always non-nil after PrepareDefaultRoomFeatures
	if !rf.InsightsFeatures.IsAllow {
		rf.InsightsFeatures.TranscriptionFeatures.IsAllow = false
		rf.InsightsFeatures.ChatTranslationFeatures.IsAllow = false
		rf.InsightsFeatures.AiFeatures.IsAllow = false
	}

	if !rf.InsightsFeatures.TranscriptionFeatures.IsAllow {
		rf.InsightsFeatures.TranscriptionFeatures.IsAllowTranslation = false
	}

	if !rf.InsightsFeatures.AiFeatures.IsAllow {
		rf.InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow = false
		rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow = false
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
