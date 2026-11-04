package utils

import (
	"testing"

	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"google.golang.org/protobuf/proto"
)

// assertInsightsNotNil verifies that every nested insights sub-message is populated.
func assertInsightsNotNil(t *testing.T, rf *plugnmeet.RoomCreateFeatures) {
	t.Helper()

	if rf.InsightsFeatures == nil {
		t.Fatalf("InsightsFeatures is nil")
	}
	if rf.InsightsFeatures.TranscriptionFeatures == nil {
		t.Fatalf("InsightsFeatures.TranscriptionFeatures is nil")
	}
	if rf.InsightsFeatures.ChatTranslationFeatures == nil {
		t.Fatalf("InsightsFeatures.ChatTranslationFeatures is nil")
	}
	if rf.InsightsFeatures.AiFeatures == nil {
		t.Fatalf("InsightsFeatures.AiFeatures is nil")
	}
	if rf.InsightsFeatures.AiFeatures.AiTextChatFeatures == nil {
		t.Fatalf("InsightsFeatures.AiFeatures.AiTextChatFeatures is nil")
	}
	if rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures == nil {
		t.Fatalf("InsightsFeatures.AiFeatures.MeetingSummarizationFeatures is nil")
	}
}

func TestMergeRoomFeaturesNilUserUsesDefaults(t *testing.T) {
	rf := mergeRoomFeatures(nil)
	if rf == nil {
		t.Fatalf("mergeRoomFeatures(nil) returned nil")
	}
	if !proto.Equal(rf, defaultRoomFeatures) {
		t.Fatalf("mergeRoomFeatures(nil) should equal defaultRoomFeatures")
	}
	assertInsightsNotNil(t, rf)
}

func TestMergeRoomFeaturesKeepsExplicitScalarFalse(t *testing.T) {
	rf := mergeRoomFeatures(&plugnmeet.RoomCreateFeatures{
		RecordingFeatures: &plugnmeet.RecordingFeatures{
			IsAllow:      true,
			IsAllowCloud: false,
		},
	})

	if !rf.RecordingFeatures.IsAllow {
		t.Fatalf("RecordingFeatures.IsAllow should be true")
	}
	if rf.RecordingFeatures.IsAllowCloud {
		t.Fatalf("RecordingFeatures.IsAllowCloud should stay false, got true from defaults")
	}
	if rf.RecordingFeatures.IsAllowLocal {
		t.Fatalf("RecordingFeatures.IsAllowLocal should be false (set message replaces default)")
	}
	if rf.RecordingFeatures.RecorderBotOptions != nil {
		t.Fatalf("RecordingFeatures.RecorderBotOptions should stay nil when nobody sets it")
	}
	assertInsightsNotNil(t, rf)
}

func TestMergeRoomFeaturesFillsUnsetInsightsSubMessages(t *testing.T) {
	// Case B: user provides only InsightsFeatures.TranscriptionFeatures
	rf := mergeRoomFeatures(&plugnmeet.RoomCreateFeatures{
		InsightsFeatures: &plugnmeet.InsightsFeatures{
			IsAllow: true,
			TranscriptionFeatures: &plugnmeet.InsightsTranscriptionFeatures{
				IsAllow: true,
			},
		},
	})

	assertInsightsNotNil(t, rf)
	if !rf.InsightsFeatures.TranscriptionFeatures.IsAllow {
		t.Fatalf("TranscriptionFeatures.IsAllow should be true")
	}
	if rf.InsightsFeatures.TranscriptionFeatures.MaxSelectedTransLangs != 0 {
		t.Fatalf("TranscriptionFeatures.MaxSelectedTransLangs should be kept as user supplied 0, got %d",
			rf.InsightsFeatures.TranscriptionFeatures.MaxSelectedTransLangs)
	}
	if !rf.InsightsFeatures.ChatTranslationFeatures.IsAllow {
		t.Fatalf("ChatTranslationFeatures.IsAllow should come from defaults")
	}
	if rf.InsightsFeatures.ChatTranslationFeatures.MaxSelectedTransLangs != 5 {
		t.Fatalf("ChatTranslationFeatures.MaxSelectedTransLangs should come from defaults, got %d",
			rf.InsightsFeatures.ChatTranslationFeatures.MaxSelectedTransLangs)
	}
	if !rf.InsightsFeatures.AiFeatures.IsAllow {
		t.Fatalf("AiFeatures.IsAllow should come from defaults")
	}
	if !rf.InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow {
		t.Fatalf("AiFeatures.AiTextChatFeatures.IsAllow should come from defaults")
	}
	if !rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow {
		t.Fatalf("AiFeatures.MeetingSummarizationFeatures.IsAllow should come from defaults")
	}
}

func TestMergeRoomFeaturesFillsUnsetAISubMessages(t *testing.T) {
	// Case C: user provides AiFeatures.AiTextChatFeatures but omits MeetingSummarizationFeatures
	rf := mergeRoomFeatures(&plugnmeet.RoomCreateFeatures{
		InsightsFeatures: &plugnmeet.InsightsFeatures{
			IsAllow: true,
			AiFeatures: &plugnmeet.InsightsAIFeatures{
				IsAllow: true,
				AiTextChatFeatures: &plugnmeet.InsightsAITextChatFeatures{
					IsAllow: true,
				},
			},
		},
	})

	assertInsightsNotNil(t, rf)
	if !rf.InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow {
		t.Fatalf("AiFeatures.AiTextChatFeatures.IsAllow should be true")
	}
	if !rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow {
		t.Fatalf("AiFeatures.MeetingSummarizationFeatures.IsAllow should come from defaults")
	}
	// ChatTranslationFeatures was omitted on user's InsightsFeatures
	if !rf.InsightsFeatures.ChatTranslationFeatures.IsAllow {
		t.Fatalf("ChatTranslationFeatures.IsAllow should come from defaults")
	}
}

func TestMergeRoomFeaturesKeepsDeepExplicitFalse(t *testing.T) {
	rf := mergeRoomFeatures(&plugnmeet.RoomCreateFeatures{
		InsightsFeatures: &plugnmeet.InsightsFeatures{
			IsAllow: true,
			AiFeatures: &plugnmeet.InsightsAIFeatures{
				IsAllow: true,
				MeetingSummarizationFeatures: &plugnmeet.InsightsAIMeetingSummarizationFeatures{
					IsAllow: false,
				},
			},
		},
	})

	assertInsightsNotNil(t, rf)
	if rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow {
		t.Fatalf("AiFeatures.MeetingSummarizationFeatures.IsAllow should stay false, got true from defaults")
	}
	if !rf.InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow {
		t.Fatalf("AiFeatures.AiTextChatFeatures.IsAllow should come from defaults")
	}
}

func TestMergeRoomFeaturesDoesNotAliasUserMessages(t *testing.T) {
	user := &plugnmeet.RoomCreateFeatures{
		RecordingFeatures: &plugnmeet.RecordingFeatures{IsAllow: true},
		InsightsFeatures: &plugnmeet.InsightsFeatures{
			IsAllow: true,
			TranscriptionFeatures: &plugnmeet.InsightsTranscriptionFeatures{
				IsAllow: true,
			},
		},
	}

	rf := mergeRoomFeatures(user)
	rf.RecordingFeatures.IsAllowLocal = true
	rf.InsightsFeatures.TranscriptionFeatures.IsAllow = false
	rf.InsightsFeatures.ChatTranslationFeatures.IsAllow = false

	if user.RecordingFeatures.IsAllowLocal {
		t.Fatalf("user's RecordingFeatures was mutated by the merge result")
	}
	if !user.InsightsFeatures.TranscriptionFeatures.IsAllow {
		t.Fatalf("user's TranscriptionFeatures was mutated by the merge result")
	}
	if user.InsightsFeatures.ChatTranslationFeatures != nil {
		t.Fatalf("user's ChatTranslationFeatures was mutated by the merge result")
	}
}

func TestPrepareDefaultRoomFeaturesDisablesNestedInsights(t *testing.T) {
	r := &plugnmeet.CreateRoomReq{
		Metadata: &plugnmeet.RoomMetadata{
			RoomFeatures: &plugnmeet.RoomCreateFeatures{
				InsightsFeatures: &plugnmeet.InsightsFeatures{
					IsAllow: false,
					TranscriptionFeatures: &plugnmeet.InsightsTranscriptionFeatures{
						IsAllow:                true,
						IsAllowTranslation:     true,
						IsAllowSpeechSynthesis: true,
					},
				},
			},
		},
	}

	PrepareDefaultRoomFeatures(r)
	SetCreateRoomDefaultValues(r, 0, 0, nil)
	rf := r.Metadata.RoomFeatures
	assertInsightsNotNil(t, rf)

	if rf.InsightsFeatures.IsAllow {
		t.Fatalf("InsightsFeatures.IsAllow should be false")
	}
	if rf.InsightsFeatures.TranscriptionFeatures.IsAllow {
		t.Fatalf("InsightsFeatures.TranscriptionFeatures.IsAllow should be false")
	}
	if rf.InsightsFeatures.ChatTranslationFeatures.IsAllow {
		t.Fatalf("InsightsFeatures.ChatTranslationFeatures.IsAllow should be false")
	}
	if rf.InsightsFeatures.AiFeatures.IsAllow {
		t.Fatalf("InsightsFeatures.AiFeatures.IsAllow should be false")
	}
	if rf.InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow {
		t.Fatalf("InsightsFeatures.AiFeatures.AiTextChatFeatures.IsAllow should be false")
	}
	if rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow {
		t.Fatalf("InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow should be false")
	}
}

func TestPrepareDefaultRoomFeaturesKeepsInsightsEnabled(t *testing.T) {
	r := &plugnmeet.CreateRoomReq{
		Metadata: &plugnmeet.RoomMetadata{
			RoomFeatures: &plugnmeet.RoomCreateFeatures{
				InsightsFeatures: &plugnmeet.InsightsFeatures{
					IsAllow: true,
					AiFeatures: &plugnmeet.InsightsAIFeatures{
						IsAllow: true,
						AiTextChatFeatures: &plugnmeet.InsightsAITextChatFeatures{
							IsAllow: true,
						},
					},
				},
			},
		},
	}

	PrepareDefaultRoomFeatures(r)
	rf := r.Metadata.RoomFeatures
	assertInsightsNotNil(t, rf)

	if !rf.InsightsFeatures.AiFeatures.MeetingSummarizationFeatures.IsAllow {
		t.Fatalf("AiFeatures.MeetingSummarizationFeatures.IsAllow should come from defaults")
	}
	if !rf.InsightsFeatures.TranscriptionFeatures.IsAllow {
		t.Fatalf("TranscriptionFeatures.IsAllow should come from defaults")
	}
}
