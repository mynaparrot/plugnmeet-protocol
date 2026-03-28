package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	projectRoot string
	rootOnce    sync.Once
)

// findAndCacheProjectRoot finds the project root by looking for go.mod and caches it.
// It starts searching from the given path and moves upwards.
func findAndCacheProjectRoot(startPath string) {
	rootOnce.Do(func() {
		dir := filepath.Dir(startPath)
		// Loop until we find go.mod or hit the filesystem root.
		for {
			if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
				projectRoot = dir
				return
			}
			parentDir := filepath.Dir(dir)
			if parentDir == dir {
				// We've reached the root of the filesystem without finding go.mod.
				projectRoot = "" // Mark as not found
				return
			}
			dir = parentDir
		}
	})
}

// SourceFormatter is a custom formatter to control the caller output.
// It wraps a standard formatter (like TextFormatter or JSONFormatter).
type SourceFormatter struct {
	// Underlying is the formatter (e.g., &logrus.TextFormatter{}) that we will delegate to.
	Underlying logrus.Formatter
	// AddSpace indicates whether to add an extra newline for readability, typically for text format.
	AddSpace bool
}

// Format renders a single log entry.
func (f *SourceFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// If the logger has caller reporting enabled, the entry will have a `Caller` field.
	if entry.HasCaller() {
		// Find and cache the project root on the first run.
		findAndCacheProjectRoot(entry.Caller.File)

		fileDir := filepath.Dir(entry.Caller.File)
		baseName := filepath.Base(entry.Caller.File)
		var fileName string

		// If the file's directory is the project root, just use the filename.
		if projectRoot != "" && fileDir == projectRoot {
			fileName = baseName
		} else {
			// Otherwise, use the parent directory and the filename.
			fileName = filepath.Join(filepath.Base(fileDir), baseName)
		}

		sourceLocation := fmt.Sprintf("%s:%d", fileName, entry.Caller.Line)

		// Add our custom-formatted source as a new field.
		entry.Data["x_file_source"] = sourceLocation
	}

	// Now, let the underlying formatter do the rest of the work with our modified entry.
	formatted, err := f.Underlying.Format(entry)
	if err != nil {
		return nil, err
	}

	if f.AddSpace {
		return append(formatted, '\n'), nil
	}

	return formatted, nil
}
