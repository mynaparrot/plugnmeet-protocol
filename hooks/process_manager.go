package hooks

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/shell"
)

const maxRecoveryAttempts = 3

// ManagedHookProcess holds the state for a long-lived hook script.
type ManagedHookProcess struct {
	cmd              *exec.Cmd
	stdin            io.WriteCloser
	stdout           io.ReadCloser
	reader           *bufio.Reader
	mutex            sync.Mutex
	log              *logrus.Entry
	recoveryAttempts int
}

// HookProcessManager manages all the long-lived hook processes.
type HookProcessManager struct {
	processes    map[string]*ManagedHookProcess
	managerMutex sync.Mutex // To protect the 'processes' map itself
	log          *logrus.Entry
	ctx          context.Context
	startOnce    sync.Once
	startErr     error
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
			if err := h.startProcess(script, 0); err != nil {
				h.stopAll() // Clean up any processes that might have started
				h.startErr = err
				return
			}
		}

		// Start a goroutine to listen for context cancellation and trigger a graceful shutdown.
		go func() {
			<-h.ctx.Done()
			h.log.Info("Main context cancelled, stopping all long-lived hook processes.")
			h.stopAll()
		}()
	})

	return h.startErr
}

// startProcess will bootup the long-lived hook process.
// it's assume that caller have hold the mutex lock, if needed.
func (h *HookProcessManager) startProcess(script string, attempt int) error {
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
		cmd:              cmd,
		stdin:            stdin,
		stdout:           stdout,
		reader:           bufio.NewReader(stdout),
		log:              log,
		recoveryAttempts: attempt,
	}

	log.Info("long-lived hook process started successfully")
	return nil
}

// recoverProcess attempts to restart a failed hook process.
func (h *HookProcessManager) recoverProcess(script string) error {
	h.managerMutex.Lock()
	defer h.managerMutex.Unlock()

	p, ok := h.processes[script]
	if !ok {
		return fmt.Errorf("process %s not found for recovery", script)
	}

	p.recoveryAttempts++
	if p.recoveryAttempts > maxRecoveryAttempts {
		return fmt.Errorf("exceeded max recovery attempts (%d) for script %s", maxRecoveryAttempts, script)
	}

	p.log.Warnf("Attempting recovery %d/%d for script", p.recoveryAttempts, maxRecoveryAttempts)

	// Clean up the old process just in case it's lingering.
	if p.cmd.Process != nil {
		_ = p.cmd.Process.Kill()
		_, _ = p.cmd.Process.Wait()
	}
	_ = p.stdin.Close()
	_ = p.stdout.Close()

	// Start a new process and replace the old one in the map.
	return h.startProcess(script, p.recoveryAttempts)
}

// isRecoverableError checks if an error indicates a crashed process.
func isRecoverableError(err error) bool {
	if err == io.EOF {
		return true
	}
	// "broken pipe" and "pipe is being closed" are common symptoms of a crashed process.
	return strings.Contains(err.Error(), "broken pipe") || strings.Contains(err.Error(), "pipe is being closed")
}

// ExecuteHook sends already marshaled JSON data to a long-lived hook script and reads its response with a timeout.
// It includes a retry mechanism for crashed scripts.
func (h *HookProcessManager) ExecuteHook(script string, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	for i := 0; i <= maxRecoveryAttempts; i++ {
		h.managerMutex.Lock()
		p, ok := h.processes[script]
		h.managerMutex.Unlock()

		if !ok {
			return nil, fmt.Errorf("hook script '%s' not found or not started", script)
		}

		p.mutex.Lock()

		ctx, cancel := context.WithTimeout(h.ctx, timeout)

		type result struct {
			data  []byte
			error error
		}
		resultChan := make(chan result, 1)

		go func() {
			_, err := p.stdin.Write(append(jsonData, '\n'))
			if err != nil {
				resultChan <- result{nil, err}
				return
			}

			responseLine, err := p.reader.ReadBytes('\n')
			if err != nil {
				resultChan <- result{nil, err}
				return
			}
			resultChan <- result{responseLine, nil}
		}()

		select {
		case <-ctx.Done():
			p.mutex.Unlock()
			cancel() // Clean up the context
			return nil, fmt.Errorf("hook script '%s' timed out after %s", script, timeout)
		case res := <-resultChan:
			p.mutex.Unlock() // Unlock before potentially recovering
			cancel()         // Clean up the context

			if res.error != nil {
				if isRecoverableError(res.error) {
					log.WithError(res.error).Warn("recoverable error detected, attempting to restart process")
					if recoveryErr := h.recoverProcess(script); recoveryErr != nil {
						return nil, recoveryErr // Recovery failed, so we exit.
					}
					continue // Retry the operation in the next loop iteration.
				}
				return nil, res.error // Unrecoverable error.
			}

			// Success
			responseLine := bytes.TrimSuffix(res.data, []byte{'\n'})
			if !json.Valid(responseLine) {
				return nil, fmt.Errorf("script '%s' returned invalid JSON: %s", script, string(responseLine))
			}
			// On success, reset the recovery counter for the process.
			p.recoveryAttempts = 0
			return responseLine, nil
		}
	}

	return nil, fmt.Errorf("hook script '%s' failed after %d retries", script, maxRecoveryAttempts)
}

// stopAll terminates all managed hook processes. This is now an internal method.
func (h *HookProcessManager) stopAll() {
	h.log.Info("Internal stopAll called for hook processes")
	h.managerMutex.Lock()
	defer h.managerMutex.Unlock()

	for _, p := range h.processes {
		p.mutex.Lock()
		_ = p.stdin.Close()
		_ = p.stdout.Close()
		p.mutex.Unlock()
		if p.cmd.Process != nil {
			// The context cancellation should have already signaled the process to stop.
			// Killing is a fallback.
			_ = p.cmd.Process.Kill()
			_, _ = p.cmd.Process.Wait()
		}
	}
}
