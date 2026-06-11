package hooks

// UploadHookData is used for the upload pipeline.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type UploadHookData struct {
	InputPath   string `json:"input_path"` // Path to the file on the local system to be uploaded.
	ServiceType string `json:"service_type"`
	RoomId      string `json:"room_id"`
	RoomSid     string `json:"room_sid"`
	RoomTableId uint64 `json:"room_table_id"`

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error     string `json:"error,omitempty"`      // Error message from the script.
	OutputPath string `json:"output_path,omitempty"` // Path/URL where the file was uploaded (e.g., S3 URL).
}

// DownloadHookData is used for the download pipeline.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type DownloadHookData struct {
	InputPath   string `json:"input_path"` // Path/URL of the file in remote storage to be downloaded.
	ServiceType string `json:"service_type"`

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error       string `json:"error,omitempty"`      // Error message from the script.
	Action      string `json:"action,omitempty"`     // e.g., "serve_local" or "redirect".
	RedirectUrl string `json:"redirect_url,omitempty"` // URL to redirect for download.
	OutputPath  string `json:"output_path,omitempty"` // Local path where the file was downloaded.
	MimeType    string `json:"mime_type,omitempty"`
}

// DeleteHookData is used for the delete hook pipeline.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type DeleteHookData struct {
	InputPath   string `json:"input_path"` // Path/URL of the file in remote storage to be deleted.
	ServiceType string `json:"service_type"`

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error string `json:"error,omitempty"` // Error message from the script.
	Msg   string `json:"msg,omitempty"`   // General message from the script.
}
