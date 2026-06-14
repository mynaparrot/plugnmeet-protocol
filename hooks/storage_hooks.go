package hooks

type HookFileType string

const (
	HookFileTypeArtifact                HookFileType = "artifact"
	HookFileTypeRecording               HookFileType = "recording"
	HookFileTypeRoomFile                HookFileType = "room-file"
	HookFileTypeWhiteboardConvertedImgs HookFileType = "whiteboard-converted-imgs"
)

// UploadHookData is used for the upload pipeline.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type UploadHookData struct {
	InputPath          string       `json:"input_path"` // Path to the file on the local system to be uploaded.
	InputDirectoryPath string       `json:"input_directory_path,omitempty"`
	HookFileType       HookFileType `json:"hook_file_type"`
	RoomId             string       `json:"room_id"`
	RoomSid            string       `json:"room_sid"`
	RoomTableId        uint64       `json:"room_table_id,omitempty"`
	FileId             string       `json:"file_id,omitempty"`

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error      string `json:"error,omitempty"`       // Error message from the script.
	OutputPath string `json:"output_path,omitempty"` // Path/URL where the file was uploaded (e.g., S3 URL).
}

type DownloadHookDataAction string

const (
	DownloadHookDataActionServeLocal DownloadHookDataAction = "serve_local"
	DownloadHookDataActionRedirect   DownloadHookDataAction = "redirect"
)

// DownloadHookData is used for the download pipeline.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type DownloadHookData struct {
	InputPath    string       `json:"input_path"` // Path/URL of the file in remote storage to be downloaded.
	HookFileType HookFileType `json:"hook_file_type"`

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error       string                 `json:"error,omitempty"`        // Error message from the script.
	Action      DownloadHookDataAction `json:"action,omitempty"`       // e.g., "serve_local" or "redirect".
	RedirectUrl string                 `json:"redirect_url,omitempty"` // URL to redirect for download.
	OutputPath  string                 `json:"output_path,omitempty"`  // Local path where the file was downloaded.
	MimeType    string                 `json:"mime_type,omitempty"`
}

// DeleteHookData is used for the delete hook pipeline.
// It serves as both input to the script (from the server) and output from the script (back to the server).
// In a script chain, subsequent scripts receive the modified data from the previous script.
type DeleteHookData struct {
	InputPath    string       `json:"input_path"` // Path/URL of the file in remote storage to be deleted.
	HookFileType HookFileType `json:"hook_file_type"`

	// Fields below are typically set by the script as output, or can be passed along in a chain.
	Error string `json:"error,omitempty"` // Error message from the script.
	Msg   string `json:"msg,omitempty"`   // General message from the script.
}

// ResumableUploadHookType defines the stage of the resumable upload.
type ResumableUploadHookType string

const (
	ResumableUploadHookTypeCheck  ResumableUploadHookType = "part-check"
	ResumableUploadHookTypeUpload ResumableUploadHookType = "part-upload"
	ResumableUploadHookTypeMerge  ResumableUploadHookType = "merge"
)

// ResumableUploadOutputType defines the response type from the script.
type ResumableUploadOutputType string

const (
	ResumableUploadOutputTypePartExists    ResumableUploadOutputType = "part_exists"
	ResumableUploadOutputTypePartNotExists ResumableUploadOutputType = "part_not_exists"
	ResumableUploadOutputTypePartUploaded  ResumableUploadOutputType = "part_uploaded"
	ResumableUploadOutputTypeMergeSuccess  ResumableUploadOutputType = "merge_success"
)

// ResumableUploadHookData is used for the resumable upload pipeline.
type ResumableUploadHookData struct {
	Type ResumableUploadHookType `json:"type"`

	// Identifiers, always present
	RoomSid             string `json:"room_sid"`
	RoomId              string `json:"room_id"`
	UserId              string `json:"user_id"`
	FileType            string `json:"file_type,omitempty"`
	ResumableIdentifier string `json:"resumable_identifier"`
	ResumableFilename   string `json:"resumable_filename,omitempty"`

	// Present for "part-check" and "part-upload"
	ResumableChunkNumber int `json:"resumable_chunk_number,omitempty"`

	// Present for "part-upload"
	InputPath string `json:"input_path,omitempty"` // Path to the temporary local chunk file

	// Present for "merge"
	ResumableTotalChunks int32 `json:"resumable_total_chunks,omitempty"`

	// --- Fields set by the script as output ---
	Error              string                    `json:"error,omitempty"`
	OutputResponseType ResumableUploadOutputType `json:"output_response_type,omitempty"`
	// OutputPath should return Final URL/path of the merged file. The format should be roomSid/other parts
	// for example: roomSid/fileName or in case converted images dir roomSid/FileId/page_1.png ...
	OutputPath    string `json:"output_path,omitempty"`
	FileMimeType  string `json:"file_mime_type,omitempty"`
	FileExtension string `json:"file_extension,omitempty"`
}

type RoomEndHookData struct {
	RoomId  string `json:"room_id"`
	RoomSid string `json:"room_sid"`
	// --- Fields set by the script as output ---
	Error string `json:"error,omitempty"`
	Msg   string `json:"msg,omitempty"`
}
