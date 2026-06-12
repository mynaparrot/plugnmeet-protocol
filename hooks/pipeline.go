package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

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

// ExecuteHookPipeline runs a series of scripts using the HookProcessManager.
// The manager will decide whether to use a long-lived process or a one-shot command.
func ExecuteHookPipeline(manager *HookProcessManager, scripts []string, initialData interface{}, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	if manager == nil {
		return nil, fmt.Errorf("hook manager is nil")
	}

	jsonData, err := json.Marshal(initialData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal initial data for hook script: %w", err)
	}

	for _, script := range scripts {
		log.Infof("Executing hook script via manager: %s", script)
		res, err := manager.ExecuteHook(script, jsonData, timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to execute hook '%s': %w", script, err)
		}
		if len(bytes.TrimSpace(res)) > 0 {
			jsonData = res
		}
	}

	return jsonData, nil
}
