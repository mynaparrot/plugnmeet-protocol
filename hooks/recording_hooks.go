package hooks

// RecordingHookData defines the JSON payload passed to and from external scripts for recorder operations.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type RecordingHookData struct {
	Task        string  `json:"task,omitempty"` // e.g., "single", "merge".
	RecordingID string  `json:"recording_id"`
	RoomTableID int64   `json:"room_table_id"`
	RoomID      string  `json:"room_id"`
	RoomSID     string  `json:"room_sid"`
	FileName    string  `json:"file_name,omitempty"` // Original file name, for single file tasks.
	RecorderID  string  `json:"recorder_id"`
	FileSize    float32 `json:"file_size,omitempty"`

	// Input paths for the script. Can be local file paths or remote storage URLs.
	InputPath  string   `json:"input_path,omitempty"`  // Primary input path for single file tasks.
	InputPaths []string `json:"input_paths,omitempty"` // Input paths for merge tasks.

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error            string `json:"error,omitempty"`              // Error message from the script.
	OutputPath       string `json:"output_path,omitempty"`        // Path/URL of the result after script execution.
	ShouldCleanup    bool   `json:"should_cleanup,omitempty"`     // Flag for post-transcoding scripts to clean up dir.
	SourceForCleanup string `json:"source_for_cleanup,omitempty"` // Original file path to clean up after transcoding.
}
