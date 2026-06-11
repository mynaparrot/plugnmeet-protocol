package hooks

// RecordingHookData defines the JSON payload passed to and from external scripts from recorder.
type RecordingHookData struct {
	Task        string   `json:"task,omitempty"` // e.g., "single", "merge"
	RecordingID string   `json:"recording_id"`
	RoomTableID int64    `json:"room_table_id"`
	RoomID      string   `json:"room_id"`
	RoomSID     string   `json:"room_sid"`
	FileName    string   `json:"file_name,omitempty"`  // For single file tasks
	FilePath    string   `json:"file_path,omitempty"`  // For single file tasks or as output
	FilePaths   []string `json:"file_paths,omitempty"` // For merge tasks
	RecorderID  string   `json:"recorder_id"`
	FileSize    float32  `json:"file_size,omitempty"` // Used in post-transcoding scripts

	// fields only for response back from scripts to recorder
	ShouldCleanup    bool   `json:"should_cleanup,omitempty"`     // Used in post-transcoding scripts to clean up dir
	SourceForCleanup string `json:"source_for_cleanup,omitempty"` // Used in post-transcoding scripts to clean up original file
	Error            string `json:"error,omitempty"`              // mostly response error message from script back to recorder
}
