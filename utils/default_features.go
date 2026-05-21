package utils

import "github.com/mynaparrot/plugnmeet-protocol/plugnmeet"

var (
	defaultRecordingFeatures = &plugnmeet.RecordingFeatures{
		IsAllow:                  true,
		IsAllowCloud:             true,
		IsAllowLocal:             true,
		EnableAutoCloudRecording: false,
	}

	defaultChatFeatures = &plugnmeet.ChatFeatures{
		IsAllow:           false,
		IsAllowFileUpload: false,
	}

	defaultSharedNotePadFeatures = &plugnmeet.SharedNotePadFeatures{
		IsAllow:  false,
		IsActive: false,
		Visible:  false,
	}

	defaultWhiteboardFeatures = &plugnmeet.WhiteboardFeatures{
		IsAllow:          false,
		Visible:          false,
		WhiteboardFileId: "default",
		FileName:         "default",
		TotalPages:       10,
	}

	defaultExternalMediaPlayerFeatures = &plugnmeet.ExternalMediaPlayerFeatures{
		IsAllow:  false,
		IsActive: false,
	}

	defaultWaitingRoomFeatures = &plugnmeet.WaitingRoomFeatures{
		IsActive:       false,
		WaitingRoomMsg: "",
	}

	defaultBreakoutRoomFeatures = &plugnmeet.BreakoutRoomFeatures{
		IsAllow:            false,
		IsActive:           false,
		AllowedNumberRooms: 6,
	}

	defaultDisplayExternalLinkFeatures = &plugnmeet.DisplayExternalLinkFeatures{
		IsAllow:  false,
		IsActive: false,
	}

	defaultIngressFeatures = &plugnmeet.IngressFeatures{
		IsAllow: false,
	}

	defaultSpeechToTextTranslationFeatures = &plugnmeet.SpeechToTextTranslationFeatures{
		IsAllow:            false,
		IsAllowTranslation: false,
	}

	defaultEndToEndEncryptionFeatures = &plugnmeet.EndToEndEncryptionFeatures{
		IsEnabled: false,
	}

	defaultPollsFeatures = &plugnmeet.PollsFeatures{
		IsAllow: false,
	}

	defaultInsightsFeatures = &plugnmeet.InsightsFeatures{
		IsAllow: false,
		TranscriptionFeatures: &plugnmeet.InsightsTranscriptionFeatures{
			IsAllow:               false,
			MaxSelectedTransLangs: 2,
		},
		ChatTranslationFeatures: &plugnmeet.InsightsChatTranslationFeatures{
			IsAllow:               false,
			MaxSelectedTransLangs: 5,
		},
		AiFeatures: &plugnmeet.InsightsAIFeatures{
			IsAllow: false,
			AiTextChatFeatures: &plugnmeet.InsightsAITextChatFeatures{
				IsAllow: false,
			},
			MeetingSummarizationFeatures: &plugnmeet.InsightsAIMeetingSummarizationFeatures{
				IsAllow: false,
			},
		},
	}

	defaultSipDialInFeatures = &plugnmeet.SipDialInFeatures{
		IsAllow: false,
	}
)
