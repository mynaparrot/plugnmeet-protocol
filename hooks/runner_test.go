package hooks

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
)

// writeHookScript creates a one-shot hook script that ignores its stdin and
// always prints the given JSON response. The response payloads used in these
// tests follow the hooks contract implemented by x-storage (pnm-hooks-storage).
func writeHookScript(t *testing.T, root, response string) string {
	t.Helper()

	script := filepath.Join(root, "hook.sh")
	content := "#!/bin/sh\ncat >/dev/null\necho '" + response + "'\n"
	if err := os.WriteFile(script, []byte(content), 0755); err != nil {
		t.Fatal(err)
	}
	return script
}

func testLog() *logrus.Entry {
	return logrus.New().WithField("test", "hooks")
}

// testHook wires a single hook category to a one-shot script and runs it.
func testHook[T interface{ GetError() string }](
	t *testing.T,
	categoryName HookCategory,
	response string,
	assign func(*Config) *HookScriptConfig,
	run func(*Config) (*T, error),
	assert func(*testing.T, *T),
) {
	t.Helper()

	root := t.TempDir()
	script := writeHookScript(t, root, response)

	config := &Config{}
	hookConfig := assign(config)
	hookConfig.Scripts = []HookScript{{Script: script, IsOneShot: true}}

	if err := config.Initialize(context.Background(), InitOptions{
		RootWorkingDir: root,
		Logger:         logrus.New(),
		Downloader:     &httpDownloader{client: http.DefaultClient},
	}, Category{Name: categoryName, Config: hookConfig}); err != nil {
		t.Fatal(err)
	}

	res, err := run(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected hook result, got nil")
	}
	assert(t, res)
}

func TestRunHooksWithNilConfig(t *testing.T) {
	var config *Config
	log := testLog()
	ctx := context.Background()

	checks := []struct {
		name string
		run  func() (bool, error)
	}{
		{"UploadHook", func() (bool, error) {
			res, err := config.RunUploadHook(&UploadHookData{}, log)
			return res != nil, err
		}},
		{"DownloadHook", func() (bool, error) {
			res, err := config.RunDownloadHook(ctx, &DownloadHookData{}, nil, 0, log)
			return res != nil, err
		}},
		{"DeleteHook", func() (bool, error) {
			res, err := config.RunDeleteHook(&DeleteHookData{}, log)
			return res != nil, err
		}},
		{"ResumableUploadHook", func() (bool, error) {
			res, err := config.RunResumableUploadHook(&ResumableUploadHookData{}, log)
			return res != nil, err
		}},
		{"RoomEndHook", func() (bool, error) {
			res, err := config.RunRoomEndHook(&RoomEndHookData{}, log)
			return res != nil, err
		}},
		{"PostRecording", func() (bool, error) {
			res, err := config.RunPostRecordingHook(&RecordingHookData{}, log)
			return res != nil, err
		}},
		{"PreTranscoding", func() (bool, error) {
			res, err := config.RunPreTranscodingHook(&RecordingHookData{}, log)
			return res != nil, err
		}},
		{"PostTranscoding", func() (bool, error) {
			res, err := config.RunPostTranscodingHook(&RecordingHookData{}, log)
			return res != nil, err
		}},
	}

	for _, check := range checks {
		t.Run(check.name, func(t *testing.T) {
			hasResult, err := check.run()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if hasResult {
				t.Error("expected nil result")
			}
		})
	}
}

func TestRunServerHooks(t *testing.T) {
	t.Run("upload", func(t *testing.T) {
		testHook(t, HookNameUpload,
			`{"output_path":"artifacts/meeting_analytics/1_room01/analytics.json"}`,
			func(c *Config) *HookScriptConfig { c.UploadHook = &HookScriptConfig{}; return c.UploadHook },
			func(c *Config) (*UploadHookData, error) {
				return c.RunUploadHook(&UploadHookData{
					HookFileType: HookFileTypeArtifact,
					RoomId:       "1_room01",
					InputPath:    "/tmp/analytics.json",
				}, testLog())
			},
			func(t *testing.T, res *UploadHookData) {
				if res.OutputPath != "artifacts/meeting_analytics/1_room01/analytics.json" {
					t.Errorf("unexpected output path: %s", res.OutputPath)
				}
			})
	})

	t.Run("download", func(t *testing.T) {
		// x-storage returns the request with action serve_local for file groups.
		testHook(t, HookNameDownload,
			`{"action":"serve_local","output_path":"/var/downloads/1_room01/group-id"}`,
			func(c *Config) *HookScriptConfig { c.DownloadHook = &HookScriptConfig{}; return c.DownloadHook },
			func(c *Config) (*DownloadHookData, error) {
				return c.RunDownloadHook(context.Background(), &DownloadHookData{
					HookFileType: HookFileTypeFileGroup,
					RoomSid:      "room-sid",
					GroupId:      "group-id",
				}, nil, 0, testLog())
			},
			func(t *testing.T, res *DownloadHookData) {
				if res.Action != DownloadHookDataActionServeLocal {
					t.Errorf("unexpected action: %s", res.Action)
				}
				if res.OutputPath != "/var/downloads/1_room01/group-id" {
					t.Errorf("unexpected output path: %s", res.OutputPath)
				}
			})
	})

	t.Run("delete", func(t *testing.T) {
		testHook(t, HookNameDelete,
			`{"msg":"File/directory moved to trash/delete successfully"}`,
			func(c *Config) *HookScriptConfig { c.DeleteHook = &HookScriptConfig{}; return c.DeleteHook },
			func(c *Config) (*DeleteHookData, error) {
				return c.RunDeleteHook(&DeleteHookData{
					HookFileType: HookFileTypeRecording,
					RoomId:       "room01",
					RoomSid:      "room-sid",
					InputPath:    "recordings/room01/rec.mp4",
				}, testLog())
			},
			func(t *testing.T, res *DeleteHookData) {
				if res.Msg != "File/directory moved to trash/delete successfully" {
					t.Errorf("unexpected msg: %s", res.Msg)
				}
			})
	})

	t.Run("resumable_upload", func(t *testing.T) {
		testHook(t, HookNameResumableUpload,
			`{"output_response_type":"merge_success","output_path":"room-sid/file.mp4","file_mime_type":"video/mp4","file_extension":"mp4"}`,
			func(c *Config) *HookScriptConfig {
				c.ResumableUploadHook = &HookScriptConfig{}
				return c.ResumableUploadHook
			},
			func(c *Config) (*ResumableUploadHookData, error) {
				return c.RunResumableUploadHook(&ResumableUploadHookData{
					Type:                 ResumableUploadHookTypeMerge,
					RoomId:               "room01",
					RoomSid:              "room-sid",
					ResumableIdentifier:  "identifier",
					ResumableFilename:    "file.mp4",
					ResumableTotalChunks: 2,
				}, testLog())
			},
			func(t *testing.T, res *ResumableUploadHookData) {
				if res.OutputResponseType != ResumableUploadOutputTypeMergeSuccess {
					t.Errorf("unexpected response type: %s", res.OutputResponseType)
				}
				if res.OutputPath != "room-sid/file.mp4" {
					t.Errorf("unexpected output path: %s", res.OutputPath)
				}
				if res.FileExtension != "mp4" {
					t.Errorf("unexpected file extension: %s", res.FileExtension)
				}
			})
	})

	t.Run("room_end", func(t *testing.T) {
		testHook(t, HookNameRoomEnd,
			`{"msg":"Cleanup for room room-sid completed"}`,
			func(c *Config) *HookScriptConfig { c.RoomEndHook = &HookScriptConfig{}; return c.RoomEndHook },
			func(c *Config) (*RoomEndHookData, error) {
				return c.RunRoomEndHook(&RoomEndHookData{
					RoomId:  "room01",
					RoomSid: "room-sid",
				}, testLog())
			},
			func(t *testing.T, res *RoomEndHookData) {
				if res.Msg != "Cleanup for room room-sid completed" {
					t.Errorf("unexpected msg: %s", res.Msg)
				}
			})
	})
}

func TestRunRecordingHooks(t *testing.T) {
	t.Run("post_recording", func(t *testing.T) {
		testHook(t, HookNamePostRecording,
			`{"task":"single","recording_id":"rec-1","output_path":"recordings/room01/rec.mp4"}`,
			func(c *Config) *HookScriptConfig { c.PostRecording = &HookScriptConfig{}; return c.PostRecording },
			func(c *Config) (*RecordingHookData, error) {
				return c.RunPostRecordingHook(&RecordingHookData{
					Task:        RecordingHookTaskSingle,
					RecordingID: "rec-1",
					RoomID:      "room01",
					FileName:    "rec.mp4",
					InputPath:   "/var/recordings/rec-1",
				}, testLog())
			},
			func(t *testing.T, res *RecordingHookData) {
				if res.OutputPath != "recordings/room01/rec.mp4" {
					t.Errorf("unexpected output path: %s", res.OutputPath)
				}
			})
	})

	t.Run("pre_transcoding", func(t *testing.T) {
		testHook(t, HookNamePreTranscoding,
			`{"task":"merge","output_path":"/tmp/plugnmeet-pre-transcode-123","should_cleanup":true}`,
			func(c *Config) *HookScriptConfig { c.PreTranscoding = &HookScriptConfig{}; return c.PreTranscoding },
			func(c *Config) (*RecordingHookData, error) {
				return c.RunPreTranscodingHook(&RecordingHookData{
					Task:        RecordingHookTaskMerge,
					RecordingID: "rec-1",
					RoomID:      "room01",
					InputPaths:  []string{"recordings/room01/part-1.mkv", "recordings/room01/part-2.mkv"},
				}, testLog())
			},
			func(t *testing.T, res *RecordingHookData) {
				if res.OutputPath != "/tmp/plugnmeet-pre-transcode-123" {
					t.Errorf("unexpected output path: %s", res.OutputPath)
				}
				if !res.ShouldCleanup {
					t.Error("expected should_cleanup to be true")
				}
			})
	})

	t.Run("post_transcoding", func(t *testing.T) {
		testHook(t, HookNamePostTranscoding,
			`{"output_path":"recordings/room01/rec.mp4","should_cleanup":true,"source_for_cleanup":"recordings/room01"}`,
			func(c *Config) *HookScriptConfig { c.PostTranscoding = &HookScriptConfig{}; return c.PostTranscoding },
			func(c *Config) (*RecordingHookData, error) {
				return c.RunPostTranscodingHook(&RecordingHookData{
					Task:        RecordingHookTaskSingle,
					RecordingID: "rec-1",
					RoomID:      "room01",
					FileName:    "rec.mp4",
					InputPath:   "/tmp/transcoded",
				}, testLog())
			},
			func(t *testing.T, res *RecordingHookData) {
				if res.OutputPath != "recordings/room01/rec.mp4" {
					t.Errorf("unexpected output path: %s", res.OutputPath)
				}
				if res.SourceForCleanup != "recordings/room01" {
					t.Errorf("unexpected source for cleanup: %s", res.SourceForCleanup)
				}
			})
	})
}
