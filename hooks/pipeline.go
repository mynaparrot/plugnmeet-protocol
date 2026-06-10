package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/shell"
)

func ValidateHookScript(scriptFullCommand string, hookType string) error {
	if scriptFullCommand == "" {
		return nil // Empty script string is valid (no script configured for this step)
	}

	// Use shell.Fields to parse the command string
	parts, err := shell.Fields(scriptFullCommand, nil)
	if err != nil {
		return fmt.Errorf("failed to parse %s script command '%s': %w", hookType, scriptFullCommand, err)
	}
	if len(parts) == 0 {
		return fmt.Errorf("empty %s script command provided", hookType)
	}
	scriptPath := parts[0] // The first part is the executable path

	info, err := os.Stat(scriptPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s script not found at %s", hookType, scriptPath)
	}
	if info.IsDir() {
		return fmt.Errorf("%s script at %s is a directory", hookType, scriptPath)
	}
	if info.Mode()&0111 == 0 {
		return fmt.Errorf("%s script at %s is not executable", hookType, scriptPath)
	}
	return nil
}

// ExecuteHookPipeline runs a series of scripts, passing data from one to the next.
func ExecuteHookPipeline(ctx context.Context, scripts []string, initialData interface{}, log *logrus.Entry) (json.RawMessage, error) {
	jsonData, err := json.Marshal(initialData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal initial data for hook script: %w", err)
	}

	for _, script := range scripts {
		log.Infof("Running storage hook script: %s", script)

		// Use mvdan.cc/sh/v3/shell to parse the command string robustly
		parts, err := shell.Fields(script, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to parse script command '%s': %w", script, err)
		}
		if len(parts) == 0 {
			return nil, fmt.Errorf("empty script command provided")
		}

		cmdPath := parts[0]
		cmdArgs := parts[1:]

		cmd := exec.CommandContext(ctx, cmdPath, cmdArgs...)
		cmd.Stdin = bytes.NewReader(jsonData)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("storage hook script '%s' failed: %w, stderr: %s", script, err, stderr.String())
		}

		if len(bytes.TrimSpace(out.Bytes())) > 0 {
			if !json.Valid(out.Bytes()) {
				return nil, fmt.Errorf("storage hook script '%s' returned invalid JSON: %s", script, out.String())
			}
			jsonData = out.Bytes()
		}
	}

	return jsonData, nil
}
