package hooks

// UploadHookRequest is the JSON payload sent to the *first* script in the upload pipeline.
type UploadHookRequest struct {
	SourceFilePath string `json:"source_file_path"`
	ServiceType    string `json:"service_type"`
	RoomId         string `json:"room_id"`
	RoomSid        string `json:"room_sid"`
	RoomTableId    uint64 `json:"room_table_id"`
}

// UploadHookResponse is the JSON payload expected from the *last* script in the upload pipeline.
type UploadHookResponse struct {
	Error       string `json:"error,omitempty"`
	LogicalPath string `json:"logical_path,omitempty"`
}

// DownloadHookRequest is the JSON payload sent to the *first* script in the download pipeline.
type DownloadHookRequest struct {
	LogicalPath string `json:"logical_path"`
	ServiceType string `json:"service_type"`
}

// DownloadHookResponse is the JSON payload expected from the *last* script in the download pipeline.
type DownloadHookResponse struct {
	Error       string `json:"error,omitempty"`
	Action      string `json:"action,omitempty"` // serve_local or redirect
	RedirectUrl string `json:"redirect_url,omitempty"`
	LocalPath   string `json:"local_path,omitempty"`
	MimeType    string `json:"mime_type,omitempty"`
}

// DeleteHookRequest is the JSON payload sent to the delete hook pipeline.
type DeleteHookRequest struct {
	LogicalPath string `json:"logical_path"`
	ServiceType string `json:"service_type"`
}

// DeleteHookResponse is the JSON payload expected from the delete hook pipeline.
type DeleteHookResponse struct {
	Error string `json:"error,omitempty"`
	Msg   string `json:"msg,omitempty"`
}
