package hooks

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// HookCategory identifies a hook category and is also used in validation error messages.
type HookCategory string

const (
	HookNameUpload          HookCategory = "upload_hook"
	HookNameDownload        HookCategory = "download_hook"
	HookNameDelete          HookCategory = "delete_hook"
	HookNameResumableUpload HookCategory = "resumable_upload_hook"
	HookNameRoomEnd         HookCategory = "room_end_hook"

	HookNamePostRecording   HookCategory = "post_recording"
	HookNamePreTranscoding  HookCategory = "pre_transcoding"
	HookNamePostTranscoding HookCategory = "post_transcoding"
)

// Config defines the YAML hook configuration shared by all plugnmeet apps.
// An app only initializes the categories it uses.
type Config struct {
	UploadHook          *HookScriptConfig `yaml:"upload_hook"`
	DownloadHook        *HookScriptConfig `yaml:"download_hook"`
	DeleteHook          *HookScriptConfig `yaml:"delete_hook"`
	ResumableUploadHook *HookScriptConfig `yaml:"resumable_upload_hook"`
	RoomEndHook         *HookScriptConfig `yaml:"room_end_hook"`

	PostRecording   *HookScriptConfig `yaml:"post_recording"`
	PreTranscoding  *HookScriptConfig `yaml:"pre_transcoding"`
	PostTranscoding *HookScriptConfig `yaml:"post_transcoding"`

	hookManager *HookProcessManager
	downloader  FileDownloader
}

// Category pairs a hook category name with its configuration.
type Category struct {
	Name   HookCategory
	Config *HookScriptConfig
}

// InitOptions controls how the hook configuration is prepared.
type InitOptions struct {
	RootWorkingDir string
	DefaultTimeout time.Duration
	Logger         *logrus.Logger
	// Downloader is optional; it is only required when RunDownloadHook must
	// fetch a redirect URL into a local directory.
	Downloader FileDownloader
}

// Initialize resolves script paths, applies defaults, validates every script
// and starts the long-lived process pools. Categories that are not passed in
// are left untouched.
func (c *Config) Initialize(ctx context.Context, opts InitOptions, categories ...Category) error {
	if opts.DefaultTimeout == 0 {
		opts.DefaultTimeout = 5 * time.Minute
	}

	resolvePath := func(scriptPath string) string {
		if !filepath.IsAbs(scriptPath) {
			scriptPath = filepath.Join(opts.RootWorkingDir, scriptPath)
		}
		return filepath.Clean(scriptPath)
	}

	// scriptsWithPoolSize maps each unique script path to its required pool size.
	// We take the max pool size if a script is used in multiple categories.
	scriptsWithPoolSize := make(map[string]int)

	for _, category := range categories {
		if category.Config == nil {
			continue
		}
		config := category.Config
		if config.PoolSize <= 0 {
			config.PoolSize = 1
		}
		if config.HookTimeout == 0 {
			config.HookTimeout = opts.DefaultTimeout
		}

		for i, script := range config.Scripts {
			var resolved string
			if strings.HasPrefix(script.Script, HookCommandHttpRequest) {
				resolved = script.Script
			} else {
				resolved = resolvePath(script.Script)
			}

			if err := ValidateHookScript(resolved, category.Name); err != nil {
				return err
			}
			config.Scripts[i].Script = resolved

			if !script.IsOneShot {
				if currentSize, ok := scriptsWithPoolSize[resolved]; !ok || config.PoolSize > currentSize {
					scriptsWithPoolSize[resolved] = config.PoolSize
				}
			}
		}
	}

	c.downloader = opts.Downloader
	c.hookManager = NewHookProcessManager(ctx, opts.Logger.WithField("service", "hook_manager"))
	if len(scriptsWithPoolSize) > 0 {
		if err := c.hookManager.StartHookProcesses(scriptsWithPoolSize); err != nil {
			return fmt.Errorf("failed to start hook processes: %w", err)
		}
	}

	return nil
}

// hasScripts reports whether a category is configured with at least one script.
func hasScripts(config *HookScriptConfig) bool {
	return config != nil && len(config.Scripts) > 0
}

func (c *Config) HasUploadHook() bool          { return c != nil && hasScripts(c.UploadHook) }
func (c *Config) HasDownloadHook() bool        { return c != nil && hasScripts(c.DownloadHook) }
func (c *Config) HasDeleteHook() bool          { return c != nil && hasScripts(c.DeleteHook) }
func (c *Config) HasResumableUploadHook() bool { return c != nil && hasScripts(c.ResumableUploadHook) }
func (c *Config) HasRoomEndHook() bool         { return c != nil && hasScripts(c.RoomEndHook) }
func (c *Config) HasPostRecording() bool       { return c != nil && hasScripts(c.PostRecording) }
func (c *Config) HasPreTranscoding() bool      { return c != nil && hasScripts(c.PreTranscoding) }
func (c *Config) HasPostTranscoding() bool     { return c != nil && hasScripts(c.PostTranscoding) }
