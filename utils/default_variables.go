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
		IsAllow:           true,
		IsAllowFileUpload: true,
	}

	defaultSharedNotePadFeatures = &plugnmeet.SharedNotePadFeatures{
		IsAllow: true,
	}

	defaultWhiteboardFeatures = &plugnmeet.WhiteboardFeatures{
		IsAllow: true,
	}

	defaultExternalMediaPlayerFeatures = &plugnmeet.ExternalMediaPlayerFeatures{
		IsAllow: true,
	}

	defaultWaitingRoomFeatures = &plugnmeet.WaitingRoomFeatures{
		IsActive: false,
	}

	defaultBreakoutRoomFeatures = &plugnmeet.BreakoutRoomFeatures{
		IsAllow:            true,
		AllowedNumberRooms: 6,
	}

	defaultDisplayExternalLinkFeatures = &plugnmeet.DisplayExternalLinkFeatures{
		IsAllow: true,
	}

	defaultIngressFeatures = &plugnmeet.IngressFeatures{
		IsAllow: true,
	}

	defaultEndToEndEncryptionFeatures = &plugnmeet.EndToEndEncryptionFeatures{
		IsEnabled: false,
	}

	defaultPollsFeatures = &plugnmeet.PollsFeatures{
		IsAllow: true,
	}

	defaultInsightsFeatures = &plugnmeet.InsightsFeatures{
		IsAllow: true,
		TranscriptionFeatures: &plugnmeet.InsightsTranscriptionFeatures{
			IsAllow:               true,
			MaxSelectedTransLangs: 2,
		},
		ChatTranslationFeatures: &plugnmeet.InsightsChatTranslationFeatures{
			IsAllow:               true,
			MaxSelectedTransLangs: 5,
		},
		AiFeatures: &plugnmeet.InsightsAIFeatures{
			IsAllow: true,
			AiTextChatFeatures: &plugnmeet.InsightsAITextChatFeatures{
				IsAllow: true,
			},
			MeetingSummarizationFeatures: &plugnmeet.InsightsAIMeetingSummarizationFeatures{
				IsAllow: true,
			},
		},
	}

	defaultSipDialInFeatures = &plugnmeet.SipDialInFeatures{
		IsAllow: true,
	}

	defaultExternalBroadcastingFeatures = &plugnmeet.ExternalBroadcastingFeatures{
		IsAllow:     true,
		IsAllowRtmp: true,
	}

	defaultLockSettings = &plugnmeet.LockSettings{
		LockWhiteboard:    new(true),
		LockScreenSharing: new(true),
		LockSharedNotepad: new(true),
	}

	defaultCopyrightConf = &plugnmeet.CopyrightConf{
		Display: true,
		Text:    "Powered by <a href=\"https://www.plugnmeet.org\" target=\"_blank\">plugNmeet</a>",
	}
)
