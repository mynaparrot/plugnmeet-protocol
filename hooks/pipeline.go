package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/shell"
)

const HookCommandHttpRequest = "http-request"

// HookScript defines a single script with its properties.
type HookScript struct {
	Script    string `yaml:"script"`
	IsOneShot bool   `yaml:"is_one_shot"`
}

// HookScriptConfig defines the configuration for a specific hook category.
type HookScriptConfig struct {
	PoolSize    int           `yaml:"pool_size"`
	Scripts     []HookScript  `yaml:"scripts"`
	HookTimeout time.Duration `yaml:"hook_timeout"`
}

func ValidateHookScript(scriptFullCommand string, hookType string) error {
	if scriptFullCommand == "" {
		return nil // Empty script string is valid (no script configured for this step)
	}

	// a special case for built-in command
	if strings.HasPrefix(scriptFullCommand, HookCommandHttpRequest) {
		parts, err := shell.Fields(scriptFullCommand, nil)
		if err != nil {
			return fmt.Errorf("failed to parse %s command '%s': %w", HookCommandHttpRequest, scriptFullCommand, err)
		}
		if len(parts) < 2 {
			return fmt.Errorf("missing URL for %s command in %s", HookCommandHttpRequest, hookType)
		}
		_, err = url.ParseRequestURI(parts[1])
		if err != nil {
			return fmt.Errorf("invalid URL for %s command in %s: %s", HookCommandHttpRequest, hookType, parts[1])
		}
		return nil
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
func ExecuteHookPipeline(manager *HookProcessManager, scripts []HookScript, initialData interface{}, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	if manager == nil {
		return nil, fmt.Errorf("hook manager is nil")
	}

	jsonData, err := json.Marshal(initialData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal initial data for hook script: %w", err)
	}

	for _, script := range scripts {
		log.Infof("Executing hook script via manager: %s", script.Script)
		res, err := manager.executeHook(script, jsonData, timeout, log)
		if err != nil {
			return nil, fmt.Errorf("failed to execute hook '%s': %w", script.Script, err)
		}

		trimmedRes := bytes.TrimSpace(res)
		if len(trimmedRes) > 0 {
			if json.Valid(trimmedRes) {
				jsonData = trimmedRes
			} else {
				log.Warnf("hook '%s' returned invalid JSON, discarding output: %s", script.Script, string(trimmedRes))
			}
		}
	}

	return jsonData, nil
}
