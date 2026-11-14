package hooks

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

// httpDownloader is a minimal FileDownloader used as a reference in tests.
// Real applications can implement the same interface with their preferred
// HTTP client (e.g. b-server uses the grab library).
type httpDownloader struct {
	client *http.Client
}

func (d *httpDownloader) Download(ctx context.Context, fileURL, dstDir string, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	dst := filepath.Join(dstDir, path.Base(req.URL.Path))
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return "", err
	}

	return dst, nil
}

// newDownloadHookConfig writes a one-shot script that always responds with
// the given download hook JSON payload.
func newDownloadHookConfig(t *testing.T, response string) *Config {
	t.Helper()

	root := t.TempDir()
	script := writeHookScript(t, root, response)

	config := &Config{
		DownloadHook: &HookScriptConfig{
			Scripts: []HookScript{{Script: script, IsOneShot: true}},
		},
	}

	if err := config.Initialize(context.Background(), InitOptions{
		RootWorkingDir: root,
		Logger:         logrus.New(),
		Downloader:     &httpDownloader{client: http.DefaultClient},
	}, Category{Name: HookNameDownload, Config: config.DownloadHook}); err != nil {
		t.Fatal(err)
	}

	return config
}

func TestRunDownloadHookDownloadsRedirectURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "hook file content")
	}))
	defer server.Close()

	response := fmt.Sprintf(`{"action":"redirect","redirect_url":%q}`, server.URL+"/file.txt")
	config := newDownloadHookConfig(t, response)

	dstDir := t.TempDir()
	res, err := config.RunDownloadHook(
		context.Background(),
		&DownloadHookData{InputPath: "remote/file.txt"},
		&dstDir,
		time.Minute,
		logrus.New().WithField("test", "download"),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Action != DownloadHookDataActionRedirect {
		t.Errorf("expected action %q, got %q", DownloadHookDataActionRedirect, res.Action)
	}

	expectedPath := filepath.Join(dstDir, "file.txt")
	if res.OutputPath != expectedPath {
		t.Errorf("expected output path %q, got %q", expectedPath, res.OutputPath)
	}

	content, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("failed to read downloaded file: %v", err)
	}
	if string(content) != "hook file content" {
		t.Errorf("unexpected file content: %q", content)
	}
}

func TestRunDownloadHookWithoutDownloader(t *testing.T) {
	root := t.TempDir()
	script := filepath.Join(root, "download_hook.sh")
	if err := os.WriteFile(script, []byte("#!/bin/sh\necho '{\"redirect_url\":\"http://example.com/file.txt\"}'\n"), 0755); err != nil {
		t.Fatal(err)
	}

	config := &Config{
		DownloadHook: &HookScriptConfig{
			Scripts: []HookScript{{Script: script, IsOneShot: true}},
		},
	}
	if err := config.Initialize(context.Background(), InitOptions{
		RootWorkingDir: root,
		Logger:         logrus.New(),
	}, Category{Name: HookNameDownload, Config: config.DownloadHook}); err != nil {
		t.Fatal(err)
	}

	dstDir := t.TempDir()
	_, err := config.RunDownloadHook(
		context.Background(),
		&DownloadHookData{InputPath: "remote/file.txt"},
		&dstDir,
		time.Minute,
		logrus.New().WithField("test", "download"),
	)
	if err == nil || !strings.Contains(err.Error(), "no file downloader configured") {
		t.Fatalf("expected missing downloader error, got %v", err)
	}
}
