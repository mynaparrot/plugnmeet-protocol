package hooks

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestInitializeAppliesDefaultsAndResolvesPaths(t *testing.T) {
	root := t.TempDir()
	scriptPath := filepath.Join(root, "scripts", "hook.sh")
	if err := os.MkdirAll(filepath.Dir(scriptPath), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(scriptPath, []byte("#!/bin/sh\n"), 0755); err != nil {
		t.Fatal(err)
	}

	config := &Config{
		UploadHook: &HookScriptConfig{
			Scripts: []HookScript{{Script: "scripts/hook.sh"}},
		},
	}

	err := config.Initialize(context.Background(), InitOptions{
		RootWorkingDir: root,
		DefaultTimeout: time.Minute,
		Logger:         logrus.New(),
	}, Category{Name: HookNameUpload, Config: config.UploadHook})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if config.UploadHook.PoolSize != 1 {
		t.Errorf("expected default pool size 1, got %d", config.UploadHook.PoolSize)
	}
	if config.UploadHook.HookTimeout != time.Minute {
		t.Errorf("expected default timeout %s, got %s", time.Minute, config.UploadHook.HookTimeout)
	}
	if got := config.UploadHook.Scripts[0].Script; got != scriptPath {
		t.Errorf("expected resolved path %s, got %s", scriptPath, got)
	}
}

func TestConfigHasHookGetters(t *testing.T) {
	var nilConfig *Config
	if nilConfig.HasUploadHook() {
		t.Error("nil config must not report hooks")
	}

	config := &Config{
		UploadHook: &HookScriptConfig{
			Scripts: []HookScript{{Script: "/bin/true", IsOneShot: true}},
		},
		// Configured but without scripts.
		PostRecording: &HookScriptConfig{},
	}

	checks := []struct {
		name string
		has  func() bool
		want bool
	}{
		{"UploadHook", config.HasUploadHook, true},
		{"DownloadHook", config.HasDownloadHook, false},
		{"DeleteHook", config.HasDeleteHook, false},
		{"ResumableUploadHook", config.HasResumableUploadHook, false},
		{"RoomEndHook", config.HasRoomEndHook, false},
		{"PostRecording", config.HasPostRecording, false},
		{"PreTranscoding", config.HasPreTranscoding, false},
		{"PostTranscoding", config.HasPostTranscoding, false},
	}

	for _, check := range checks {
		if got := check.has(); got != check.want {
			t.Errorf("%s: expected %v, got %v", check.name, check.want, got)
		}
	}
}

func TestInitializeSkipsUnlistedCategories(t *testing.T) {
	config := &Config{
		// Not passed to Initialize, so it must not be validated or modified.
		PostRecording: &HookScriptConfig{
			Scripts: []HookScript{{Script: "/does/not/exist"}},
		},
	}

	err := config.Initialize(context.Background(), InitOptions{
		RootWorkingDir: t.TempDir(),
		Logger:         logrus.New(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if config.PostRecording.HookTimeout != 0 || config.PostRecording.PoolSize != 0 {
		t.Errorf("unlisted category was modified: %+v", config.PostRecording)
	}
}
