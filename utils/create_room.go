package utils

import (
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"time"
)

func PrepareDefaultRoomFeatures(r *plugnmeet.CreateRoomReq) {
	rf := r.Metadata.RoomFeatures

	if rf.RecordingFeatures == nil {
		rf.RecordingFeatures = &plugnmeet.RecordingFeatures{
			IsAllow:                  true,
			IsAllowCloud:             true,
			IsAllowLocal:             true,
			EnableAutoCloudRecording: false,
		}
	}

	if rf.ChatFeatures == nil {
		rf.ChatFeatures = &plugnmeet.ChatFeatures{
			AllowChat:       false,
			AllowFileUpload: false,
		}
	}

	if rf.SharedNotePadFeatures == nil {
		rf.SharedNotePadFeatures = &plugnmeet.SharedNotePadFeatures{
			AllowedSharedNotePad: false,
			IsActive:             false,
			Visible:              false,
		}
	}

	if rf.WhiteboardFeatures == nil {
		rf.WhiteboardFeatures = &plugnmeet.WhiteboardFeatures{
			AllowedWhiteboard: false,
			Visible:           false,
			WhiteboardFileId:  "default",
			FileName:          "default",
			TotalPages:        10,
		}
	}

	if rf.ExternalMediaPlayerFeatures == nil {
		rf.ExternalMediaPlayerFeatures = &plugnmeet.ExternalMediaPlayerFeatures{
			AllowedExternalMediaPlayer: false,
			IsActive:                   false,
		}
	}

	if rf.WaitingRoomFeatures == nil {
		rf.WaitingRoomFeatures = &plugnmeet.WaitingRoomFeatures{
			IsActive:       false,
			WaitingRoomMsg: "",
		}
	}

	if rf.BreakoutRoomFeatures == nil {
		rf.BreakoutRoomFeatures = &plugnmeet.BreakoutRoomFeatures{
			IsAllow:            false,
			IsActive:           false,
			AllowedNumberRooms: 6,
		}
	}

	if rf.DisplayExternalLinkFeatures == nil {
		rf.DisplayExternalLinkFeatures = &plugnmeet.DisplayExternalLinkFeatures{
			IsAllow:  false,
			IsActive: false,
		}
	}

	if r.Metadata.DefaultLockSettings == nil {
		r.Metadata.DefaultLockSettings = new(plugnmeet.LockSettings)
	}

	r.Metadata.StartedAt = uint64(time.Now().Unix())
	r.Metadata.RoomFeatures = rf
}

func SetCreateRoomDefaultValues(r *plugnmeet.CreateRoomReq, maxSize uint64, allowedTypes []string, allowedNotepad bool) {
	rf := r.Metadata.RoomFeatures

	if rf.SharedNotePadFeatures.AllowedSharedNotePad && !allowedNotepad {
		rf.SharedNotePadFeatures.AllowedSharedNotePad = false
	}

	if rf.ChatFeatures.AllowFileUpload {
		if len(rf.ChatFeatures.AllowedFileTypes) == 0 {
			rf.ChatFeatures.AllowedFileTypes = allowedTypes
		}
		rf.ChatFeatures.MaxFileSize = &maxSize
	}

	if rf.BreakoutRoomFeatures.IsAllow && rf.BreakoutRoomFeatures.AllowedNumberRooms == 0 {
		rf.BreakoutRoomFeatures.AllowedNumberRooms = 6
	}
}

func SetRoomDefaultLockSettings(r *plugnmeet.CreateRoomReq) {
	lock := new(bool)
	if r.Metadata.DefaultLockSettings.LockScreenSharing == nil {
		*lock = true
		r.Metadata.DefaultLockSettings.LockScreenSharing = lock
	}
	if r.Metadata.DefaultLockSettings.LockWhiteboard == nil {
		*lock = true
		r.Metadata.DefaultLockSettings.LockWhiteboard = lock
	}
	if r.Metadata.DefaultLockSettings.LockSharedNotepad == nil {
		*lock = true
		r.Metadata.DefaultLockSettings.LockSharedNotepad = lock
	}
}
