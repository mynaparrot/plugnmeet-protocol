package hooks

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/shell"
)

// ManagedHookProcess holds the state for a long-lived hook script.
type ManagedHookProcess struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	reader *bufio.Reader
	mutex  sync.Mutex
	log    *logrus.Entry
}

// HookProcessManager manages all the long-lived hook processes.
type HookProcessManager struct {
	processes map[string]*ManagedHookProcess
	log       *logrus.Entry
	ctx       context.Context
	startOnce sync.Once
	startErr  error
}

// NewHookProcessManager creates a new manager.
// The provided context should be the main application context to manage the lifecycle of the hook processes.
func NewHookProcessManager(ctx context.Context, log *logrus.Entry) *HookProcessManager {
	return &HookProcessManager{
		processes: make(map[string]*ManagedHookProcess),
		log:       log,
		ctx:       ctx,
	}
}

// StartHookProcesses starts all unique scripts as long-lived processes and starts a listener for context cancellation.
// It is protected by a sync.Once to ensure it only runs once.
func (h *HookProcessManager) StartHookProcesses(scripts []string) error {
	h.startOnce.Do(func() {
		uniqueScripts := make(map[string]bool)
		for _, s := range scripts {
			if s != "" {
				uniqueScripts[s] = true
			}
		}

		for script := range uniqueScripts {
			if err := h.startProcess(script); err != nil {
				h.stopAll() // Clean up any processes that might have started
				h.startErr = err
				return
			}
		}

		// Start a goroutine to listen for context cancellation and trigger a graceful shutdown.
		go func() {
			<-h.ctx.Done()
			h.log.Info("main context cancelled, stopping all long-lived hook processes.")
			h.stopAll()
		}()
	})

	return h.startErr
}

func (h *HookProcessManager) startProcess(script string) error {
	log := h.log.WithField("hook_script", script)
	log.Info("starting long-lived hook process")

	parts, err := shell.Fields(script, nil)
	if err != nil {
		return fmt.Errorf("failed to parse script command '%s': %w", script, err)
	}
	if len(parts) == 0 {
		return fmt.Errorf("empty script command provided")
	}

	cmd := exec.CommandContext(h.ctx, parts[0], parts[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe for '%s': %w", script, err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe for '%s': %w", script, err)
	}
	// Discard stderr to prevent conflicts and unwanted logging from the child process.
	cmd.Stderr = io.Discard

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start script '%s': %w", script, err)
	}

	h.processes[script] = &ManagedHookProcess{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
		reader: bufio.NewReader(stdout),
		log:    log,
	}

	log.Info("long-lived hook process started successfully")
	return nil
}

// ExecuteHook sends already marshaled JSON data to a long-lived hook script and reads its response with a timeout.
func (h *HookProcessManager) ExecuteHook(script string, jsonData json.RawMessage, timeout time.Duration) (json.RawMessage, error) {
	p, ok := h.processes[script]
	if !ok {
		return nil, fmt.Errorf("hook script '%s' not found or not started", script)
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Create a context with a timeout for this specific execution.
	ctx, cancel := context.WithTimeout(h.ctx, timeout)
	defer cancel()

	// Channel to receive the result from the blocking read operation.
	type result struct {
		data  []byte
		error error
	}
	resultChan := make(chan result, 1)

	// Perform the write and blocking read in a separate goroutine.
	go func() {
		_, err := p.stdin.Write(append(jsonData, '\n'))
		if err != nil {
			resultChan <- result{nil, fmt.Errorf("failed to write to stdin for script '%s': %w", script, err)}
			return
		}

		responseLine, err := p.reader.ReadBytes('\n')
		if err != nil {
			resultChan <- result{nil, fmt.Errorf("failed to read from stdout for script '%s': %w", script, err)}
			return
		}
		resultChan <- result{responseLine, nil}
	}()

	// Wait for either the result or the timeout.
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("hook script '%s' timed out after %s", script, timeout)
	case res := <-resultChan:
		if res.error != nil {
			return nil, res.error
		}
		responseLine := bytes.TrimSuffix(res.data, []byte{'\n'})
		if !json.Valid(responseLine) {
			return nil, fmt.Errorf("script '%s' returned invalid JSON: %s", script, string(responseLine))
		}
		return responseLine, nil
	}
}

// stopAll terminates all managed hook processes. This is now an internal method.
func (h *HookProcessManager) stopAll() {
	h.log.Info("internal stopAll called for hook processes")
	for _, p := range h.processes {
		p.mutex.Lock()
		p.stdin.Close()
		p.stdout.Close()
		p.mutex.Unlock()
		if p.cmd.Process != nil {
			// The context cancellation should have already signaled the process to stop.
			// Killing is a fallback.
			p.cmd.Process.Kill()
			p.cmd.Wait()
		}
	}
}
