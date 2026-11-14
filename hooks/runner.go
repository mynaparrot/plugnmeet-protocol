package hooks

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// FileDownloader fetches a file from a URL into a destination directory.
type FileDownloader interface {
	Download(ctx context.Context, fileURL, dstDir string, timeout time.Duration) (fileFullPath string, err error)
}

// runPipeline executes a hook pipeline and decodes the final payload.
// It returns nil without error when the category has no scripts configured.
func runPipeline[T interface{ GetError() string }](manager *HookProcessManager, config *HookScriptConfig, req *T, log *logrus.Entry, name HookCategory) (*T, error) {
	if config == nil || len(config.Scripts) == 0 {
		return nil, nil
	}

	resBytes, err := ExecuteHookPipeline(manager, config.Scripts, req, config.HookTimeout, log)
	if err != nil {
		return nil, fmt.Errorf("%s pipeline failed with error: %w", name, err)
	}

	var res T
	if err := json.Unmarshal(resBytes, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s response: %w", name, err)
	}

	if msg := res.GetError(); msg != "" {
		return nil, fmt.Errorf("%s script returned an error: %s", name, msg)
	}

	return &res, nil
}

// RunUploadHook executes the pipeline for uploading files.
// Returns nil if no upload hook is configured.
func (c *Config) RunUploadHook(req *UploadHookData, log *logrus.Entry) (*UploadHookData, error) {
	res, err := runPipeline(c.hookManager, c.UploadHook, req, log, HookNameUpload)
	if err != nil || res == nil {
		return nil, err
	}

	log.Infof("Upload hook successful proceeded with output %s ", res.OutputPath)
	return res, nil
}

// RunDownloadHook executes the pipeline for downloading files.
// If localDownloadDirPath is provided and the script returns a redirect URL,
// the file is downloaded to that directory.
// Returns nil if no download hook is configured.
func (c *Config) RunDownloadHook(ctx context.Context, req *DownloadHookData, localDownloadDirPath *string, downloadTimeout time.Duration, log *logrus.Entry) (*DownloadHookData, error) {
	res, err := runPipeline(c.hookManager, c.DownloadHook, req, log, HookNameDownload)
	if err != nil || res == nil {
		return nil, err
	}

	if localDownloadDirPath != nil && res.RedirectUrl != "" {
		if c.downloader == nil {
			return nil, fmt.Errorf("%s: no file downloader configured", HookNameDownload)
		}
		if err := os.MkdirAll(*localDownloadDirPath, 0755); err != nil {
			return nil, err
		}
		filePath, err := c.downloader.Download(ctx, res.RedirectUrl, *localDownloadDirPath, downloadTimeout)
		if err != nil {
			return nil, err
		}
		res.OutputPath = filePath
	}

	return res, nil
}

// RunDeleteHook executes the pipeline for deleting files.
// Returns nil if no delete hook is configured.
func (c *Config) RunDeleteHook(req *DeleteHookData, log *logrus.Entry) (*DeleteHookData, error) {
	return runPipeline(c.hookManager, c.DeleteHook, req, log, HookNameDelete)
}

// RunResumableUploadHook executes the pipeline for resumable uploads.
// Returns nil if no resumable upload hook is configured.
func (c *Config) RunResumableUploadHook(req *ResumableUploadHookData, log *logrus.Entry) (*ResumableUploadHookData, error) {
	return runPipeline(c.hookManager, c.ResumableUploadHook, req, log, HookNameResumableUpload)
}

// RunRoomEndHook executes the pipeline for room end hooks.
// Returns nil if no room end hook is configured.
func (c *Config) RunRoomEndHook(req *RoomEndHookData, log *logrus.Entry) (*RoomEndHookData, error) {
	return runPipeline(c.hookManager, c.RoomEndHook, req, log, HookNameRoomEnd)
}

// RunPostRecordingHook executes the pipeline for post recording hooks.
// Returns nil if no post recording hook is configured.
func (c *Config) RunPostRecordingHook(req *RecordingHookData, log *logrus.Entry) (*RecordingHookData, error) {
	return runPipeline(c.hookManager, c.PostRecording, req, log, HookNamePostRecording)
}

// RunPreTranscodingHook executes the pipeline for pre transcoding hooks.
// Returns nil if no pre transcoding hook is configured.
func (c *Config) RunPreTranscodingHook(req *RecordingHookData, log *logrus.Entry) (*RecordingHookData, error) {
	return runPipeline(c.hookManager, c.PreTranscoding, req, log, HookNamePreTranscoding)
}

// RunPostTranscodingHook executes the pipeline for post transcoding hooks.
// Returns nil if no post transcoding hook is configured.
func (c *Config) RunPostTranscodingHook(req *RecordingHookData, log *logrus.Entry) (*RecordingHookData, error) {
	return runPipeline(c.hookManager, c.PostTranscoding, req, log, HookNamePostTranscoding)
}
