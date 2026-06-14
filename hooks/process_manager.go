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

// ManagedHookProcess holds the state for a long-lived hook script.
type ManagedHookProcess struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
	reader *bufio.Reader
	log    *logrus.Entry
	// A unique identifier for this process instance, useful for logging.
	instanceId int
}

// HookProcessManager manages all the long-lived hook processes.
type HookProcessManager struct {
	// processPools holds a channel for each script, which acts as a pool of available processes.
	processPools map[string]chan *ManagedHookProcess
	// allProcesses holds all created processes for cleanup purposes.
	allProcesses []*ManagedHookProcess
	// managerMutex protects the process maps/slices during concurrent operations like recovery and shutdown.
	// It is NOT needed during the single-threaded initialization phase.
	managerMutex sync.Mutex
	log          *logrus.Entry
	ctx          context.Context
	startOnce    sync.Once
	startErr     error
}

// NewHookProcessManager creates a new manager.
func NewHookProcessManager(ctx context.Context, log *logrus.Entry) *HookProcessManager {
	return &HookProcessManager{
		processPools: make(map[string]chan *ManagedHookProcess),
		allProcesses: make([]*ManagedHookProcess, 0),
		log:          log,
		ctx:          ctx,
	}
}

// StartHookProcesses starts pools of long-lived processes for each unique script.
func (h *HookProcessManager) StartHookProcesses(scriptsWithPoolSize map[string]int) error {
	h.startOnce.Do(func() {
		// This block is guaranteed to run exactly once.
		// No lock is needed here as no other goroutine can access the maps yet.
		for script, poolSize := range scriptsWithPoolSize {
			if poolSize <= 0 {
				poolSize = 1
			}
			h.log.Infof("starting process pool for script '%s' with size %d", script, poolSize)
			h.processPools[script] = make(chan *ManagedHookProcess, poolSize)

			for i := 0; i < poolSize; i++ {
				process, err := h.startProcess(script, i)
				if err != nil {
					h.startErr = fmt.Errorf("failed to start instance %d for script '%s': %w", i, script, err)
					return // Exit the Do func immediately on error.
				}
				h.processPools[script] <- process
				h.allProcesses = append(h.allProcesses, process)
			}
		}

		// Start the shutdown listener only after a fully successful initialization.
		go func() {
			<-h.ctx.Done()
			h.log.Info("main context cancelled, stopping all long-lived hook processes")
			h.stopAll()
		}()
	})

	// If initialization failed, we must clean up any processes that were started.
	if h.startErr != nil {
		h.stopAll()
	}

	return h.startErr
}

// startProcess boots up a single long-lived hook process instance.
func (h *HookProcessManager) startProcess(script string, instanceId int) (*ManagedHookProcess, error) {
	log := h.log.WithFields(logrus.Fields{
		"hook_script": script,
		"instance_id": instanceId,
	})
	log.Info("starting new hook process instance")

	parts, err := shell.Fields(script, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse script command '%s': %w", script, err)
	}
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty script command provided")
	}

	cmd := exec.CommandContext(h.ctx, parts[0], parts[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	// for mostly logging from scripts for better debug
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start script: %w", err)
	}

	// Goroutine to log stderr for this process.
	go func() {
		scanner := bufio.NewScanner(stderr)
		log := log.WithField("pipe", "stderr")
		for scanner.Scan() {
			log.Infoln(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.WithError(err).Error("error reading from hook script stderr")
		}
	}()

	process := &ManagedHookProcess{
		cmd:        cmd,
		stdin:      stdin,
		stdout:     stdout,
		stderr:     stderr,
		reader:     bufio.NewReader(stdout),
		log:        log,
		instanceId: instanceId,
	}

	log.Info("hook process instance started successfully")
	return process, nil
}

// recoverProcess attempts to restart a failed hook process and returns it to the pool.
func (h *HookProcessManager) recoverProcess(script string, oldProcess *ManagedHookProcess) {
	h.managerMutex.Lock()
	defer h.managerMutex.Unlock()

	// Before recovering, check if the main context is done. If so, the server is shutting down.
	if h.ctx.Err() != nil {
		h.log.Warnf("server is shutting down, skipping recovery for process instance %d", oldProcess.instanceId)
		return
	}

	log := h.log.WithField("hook_script", script)
	log.Warnf("attempting to recover failed process instance %d", oldProcess.instanceId)

	// Clean up the old process resources.
	_ = oldProcess.stdin.Close()
	_ = oldProcess.stdout.Close()
	_ = oldProcess.stderr.Close()
	if oldProcess.cmd.Process != nil {
		_ = oldProcess.cmd.Process.Kill()
		_, _ = oldProcess.cmd.Process.Wait()
	}

	// Attempt to start a new process.
	newProcess, err := h.startProcess(script, oldProcess.instanceId)
	if err != nil {
		log.WithError(err).Error("failed to recover process; pool size for this script will be reduced")
		// The pool for this script will now have one less process.
		// We could implement more complex logic here, like retrying the recovery.
		return
	}

	// Replace the old process with the new one in the global list.
	for i, p := range h.allProcesses {
		if p == oldProcess {
			h.allProcesses[i] = newProcess
			break
		}
	}

	// Return the newly recovered process to the pool.
	h.processPools[script] <- newProcess
	log.Info("successfully recovered process instance")
}

// isRecoverableError checks if an error indicates a crashed process.
func isRecoverableError(err error) bool {
	if err == io.EOF {
		return true
	}
	return strings.Contains(err.Error(), "broken pipe") || strings.Contains(err.Error(), "pipe is being closed")
}

// ExecuteHook gets a process from the pool, sends it data, and reads the response.
func (h *HookProcessManager) ExecuteHook(script string, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	pool, ok := h.processPools[script]
	if !ok {
		return nil, fmt.Errorf("no process pool found for hook script '%s'", script)
	}

	// Get a process from the pool, waiting if none are available.
	var process *ManagedHookProcess
	select {
	case process = <-pool:
		// Got a process.
	case <-h.ctx.Done():
		return nil, fmt.Errorf("server is shutting down; could not get process for script '%s'", script)
	}

	// This context is for the request's timeout.
	ctx, cancel := context.WithTimeout(h.ctx, timeout)
	defer cancel()

	type result struct {
		data  []byte
		error error
	}
	resultChan := make(chan result, 1)

	// This worker goroutine is responsible for the entire lifecycle of the process for this request.
	// It will always either return the process to the pool or recover it.
	go func() {
		_, err := process.stdin.Write(append(jsonData, '\n'))
		if err != nil {
			if isRecoverableError(err) {
				log.WithError(err).Warnf("recoverable error in script '%s' (instance %d)", script, process.instanceId)
				go h.recoverProcess(script, process)
			} else {
				// For non-recoverable write errors, we also recover.
				go h.recoverProcess(script, process)
			}
			resultChan <- result{nil, err}
			return
		}

		responseLine, err := process.reader.ReadBytes('\n')
		if err != nil {
			if isRecoverableError(err) {
				log.WithError(err).Warnf("recoverable error in script '%s' (instance %d)", script, process.instanceId)
				go h.recoverProcess(script, process)
			} else {
				// For non-recoverable read errors, we also recover.
				go h.recoverProcess(script, process)
			}
			resultChan <- result{nil, err}
			return
		}

		// Success! Return the process to the pool.
		pool <- process
		resultChan <- result{responseLine, nil}
	}()

	select {
	case <-ctx.Done():
		// The request timed out. The worker goroutine is still running in the background.
		// It will eventually finish and handle returning or recovering the process.
		// We just return the timeout error to the caller.
		return nil, fmt.Errorf("hook script '%s' (instance %d) timed out after %s", script, process.instanceId, timeout)

	case res := <-resultChan:
		// The worker finished before the timeout.
		if res.error != nil {
			return nil, fmt.Errorf("process for script '%s' failed: %w", script, res.error)
		}

		responseLine := bytes.TrimSuffix(res.data, []byte{'\n'})
		if !json.Valid(responseLine) {
			return nil, fmt.Errorf("script '%s' returned invalid JSON: %s", script, string(responseLine))
		}
		return responseLine, nil
	}
}

// stopAll terminates all managed hook processes.
func (h *HookProcessManager) stopAll() {
	h.managerMutex.Lock()
	defer h.managerMutex.Unlock()

	for _, p := range h.allProcesses {
		_ = p.stdin.Close()
		_ = p.stdout.Close()
		_ = p.stderr.Close()
		if p.cmd.Process != nil {
			_ = p.cmd.Process.Kill()
			_, _ = p.cmd.Process.Wait()
		}
	}
	// Clear the lists.
	h.allProcesses = make([]*ManagedHookProcess, 0)
	h.processPools = make(map[string]chan *ManagedHookProcess)
}
