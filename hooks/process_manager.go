package hooks

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/shell"
)

const (
	// ProcessStateHealthy indicates the process is running and available.
	ProcessStateHealthy int32 = 0
	// ProcessStateRecovering indicates the process has been found to be unhealthy and is being recovered.
	ProcessStateRecovering int32 = 1
)

// ManagedHookProcess holds the state for a true long-lived hook script.
type ManagedHookProcess struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
	reader *bufio.Reader
	log    *logrus.Entry
	// A unique identifier for this process instance, useful for logging.
	instanceId int
	// The script this process is an instance of.
	script string
	// The current state of the process (Healthy, Recovering).
	state atomic.Int32
	// How many times this process slot has been recovered.
	recoveryCount int
}

// HookProcessManager manages all the long-lived hook processes.
type HookProcessManager struct {
	// processPools holds a channel for each script, which acts as a pool of available processes.
	processPools map[string]chan *ManagedHookProcess
	// allProcesses maps a script string to its list of running processes.
	allProcesses map[string][]*ManagedHookProcess
	// managerMutex protects the process maps/slices during concurrent operations like recovery and shutdown.
	managerMutex sync.RWMutex
	log          *logrus.Entry
	ctx          context.Context
	startOnce    sync.Once
	startErr     error
	httpClient   *http.Client
}

// NewHookProcessManager creates a new manager.
// The provided context should be the main application context to manage the lifecycle of the hook processes.
func NewHookProcessManager(ctx context.Context, log *logrus.Entry) *HookProcessManager {
	return &HookProcessManager{
		processPools: make(map[string]chan *ManagedHookProcess),
		allProcesses: make(map[string][]*ManagedHookProcess),
		log:          log,
		ctx:          ctx,
		httpClient:   &http.Client{},
	}
}

// StartHookProcesses starts pools of long-lived processes for each unique script.
// It is protected by a sync.Once to ensure it only runs once.
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

			scriptProcesses := make([]*ManagedHookProcess, 0, poolSize)
			for i := 0; i < poolSize; i++ {
				process, err := h.startNativeProcess(script, i, 0)
				if err != nil {
					h.startErr = fmt.Errorf("failed to start instance %d for script '%s': %w", i, script, err)
					return // Exit the Do func immediately on error.
				}
				h.processPools[script] <- process
				scriptProcesses = append(scriptProcesses, process)
			}
			h.allProcesses[script] = scriptProcesses

			// Launch monitors only after the initial state is fully consistent.
			for _, process := range scriptProcesses {
				go h.monitorProcess(process)
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

// startNativeProcess boots up a true long-lived OS process.
func (h *HookProcessManager) startNativeProcess(script string, instanceId int, recoveryCount int) (*ManagedHookProcess, error) {
	log := h.log.WithFields(logrus.Fields{
		"hook_script":    script,
		"instance_id":    instanceId,
		"recovery_count": recoveryCount,
	})
	log.Info("starting native long-lived process instance")

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
			log.Debugln(scanner.Text())
		}
	}()

	process := &ManagedHookProcess{
		cmd:           cmd,
		stdin:         stdin,
		stdout:        stdout,
		stderr:        stderr,
		reader:        bufio.NewReader(stdout),
		log:           log,
		instanceId:    instanceId,
		script:        script,
		recoveryCount: recoveryCount,
	}

	log.Info("native long-lived process instance started successfully")
	return process, nil
}

// monitorProcess waits for a process to exit and triggers recovery.
// This is the definitive way to detect a crashed process.
func (h *HookProcessManager) monitorProcess(p *ManagedHookProcess) {
	err := p.cmd.Wait()
	// Wait() will block until the process exits. When it returns, the process is dead.
	// We must now trigger recovery.

	p.log.WithError(err).Warn("process has exited unexpectedly. Triggering recovery.")

	// The existing recoverProcess function is already safe for concurrent calls.
	h.recoverProcess(p.script, p)
}

// executeHook acts as a dispatcher. It executes one-shot commands directly
// and uses the process pool for long-lived scripts.
func (h *HookProcessManager) executeHook(script HookScript, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	if script.IsOneShot {
		return h.executeOneShotCommand(script, jsonData, timeout, log)
	}
	return h.executePooledProcess(script.Script, jsonData, timeout, log)
}

// executeOneShotCommand executes a command like 'curl' directly.
func (h *HookProcessManager) executeOneShotCommand(script HookScript, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	if strings.HasPrefix(script.Script, HookCommandHttpRequest) {
		return h.executeHttpRequest(script, jsonData, timeout, log)
	}

	log.Infof("executing one-shot command: %s", script.Script)

	if timeout == 0 {
		timeout = 5 * time.Minute
	}
	ctx, cancel := context.WithTimeout(h.ctx, timeout)
	defer cancel()

	parts, err := shell.Fields(script.Script, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse one-shot script command '%s': %w", script.Script, err)
	}

	cmd := exec.CommandContext(ctx, parts[0], parts[1:]...)
	cmd.Stdin = bytes.NewReader(jsonData)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("one-shot script execution failed: %w, output: %s", err, string(output))
	}

	responseLine := bytes.TrimSuffix(output, []byte{'\n'})
	if len(responseLine) > 0 && !json.Valid(responseLine) {
		return nil, fmt.Errorf("one-shot script '%s' returned invalid JSON: %s", script.Script, string(responseLine))
	}

	return responseLine, nil
}

func (h *HookProcessManager) executeHttpRequest(script HookScript, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	log.Infof("executing %s: %s", HookCommandHttpRequest, script.Script)

	if timeout == 0 {
		timeout = 5 * time.Minute
	}
	ctx, cancel := context.WithTimeout(h.ctx, timeout)
	defer cancel()

	parts, err := shell.Fields(script.Script, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s script command '%s': %w", HookCommandHttpRequest, script.Script, err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", parts[1], bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request execution failed: %w", err)
	}
	defer resp.Body.Close()

	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %w", err)
	}

	responseLine := bytes.TrimSuffix(output, []byte{'\n'})
	if len(responseLine) > 0 && !json.Valid(responseLine) {
		return nil, fmt.Errorf("%s '%s' returned invalid JSON: %s", HookCommandHttpRequest, script.Script, string(responseLine))
	}

	return responseLine, nil
}

// executePooledProcess executes a command using the long-lived process pool.
func (h *HookProcessManager) executePooledProcess(script string, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	h.managerMutex.RLock()
	pool, ok := h.processPools[script]
	h.managerMutex.RUnlock()

	if !ok {
		h.log.Warnf("no process pool found for long-lived script '%s', it will be skipped", script)
		return jsonData, nil
	}

	var process *ManagedHookProcess
	select {
	case process = <-pool:
	// It's possible to get a process that has just crashed, before the monitor triggers.
	// The isRecoverableError check below will handle this.
	case <-h.ctx.Done():
		return nil, fmt.Errorf("server is shutting down; could not get process for script '%s'", script)
	}

	if timeout == 0 {
		timeout = 5 * time.Minute
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
	go func() {
		var err error
		defer func() {
			if err != nil {
				// If an error occurs, we assume the process is compromised.
				// We DO NOT return it to the pool. The monitor will handle recovery.
				log.WithError(err).Warn("Request failed and process will be abandoned. The process will be recovered by its monitor if it has crashed.")
			} else {
				// Success, return the process to the pool.
				pool <- process
			}
		}()

		if _, err = process.stdin.Write(append(jsonData, '\n')); err != nil {
			resultChan <- result{nil, err}
			return
		}

		var responseLine []byte
		responseLine, err = process.reader.ReadBytes('\n')
		if err != nil {
			resultChan <- result{nil, err}
			return
		}

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
		if len(responseLine) > 0 && !json.Valid(responseLine) {
			return nil, fmt.Errorf("script '%s' returned invalid JSON: %s", script, string(responseLine))
		}
		return responseLine, nil
	}
}

// recoverProcess rebuilds the entire process pool for a script when one instance fails.
func (h *HookProcessManager) recoverProcess(script string, oldProcess *ManagedHookProcess) {
	// 1. Atomically check and mark the process. This is a fast, non-blocking operation.
	if !oldProcess.state.CompareAndSwap(ProcessStateHealthy, ProcessStateRecovering) {
		h.log.Debugf("process instance %d for script '%s' is already being recovered. Skipping.", oldProcess.instanceId, script)
		return
	}

	// 2. Check for server shutdown before sleeping.
	if h.ctx.Err() != nil {
		h.log.Warnf("server is shutting down, skipping recovery for process instance %d", oldProcess.instanceId)
		return
	}

	log := h.log.WithField("hook_script", script)
	log.Warnf("process instance %d crashed. Waiting 5s before attempting pool rebuild...", oldProcess.instanceId)

	// 3. Perform the sleep BEFORE acquiring the global lock.
	time.Sleep(5 * time.Second)

	// 4. NOW, acquire the global lock to perform the disruptive pool rebuild.
	h.managerMutex.Lock()

	// 5. After sleeping and getting the lock, we must re-validate.
	if h.ctx.Err() != nil {
		h.log.Warnf("server is shutting down, aborting recovery for process instance %d", oldProcess.instanceId)
		h.managerMutex.Unlock()
		return
	}

	var isStillPresent = false
	if processes, ok := h.allProcesses[script]; ok {
		for _, p := range processes {
			if p == oldProcess {
				isStillPresent = true
				break
			}
		}
	}
	if !isStillPresent {
		log.Warnf("process for script '%s' was already removed during sleep. Aborting rebuild.", script)
		h.managerMutex.Unlock()
		return
	}

	// --- Proceed with the rebuild logic ---
	newRecoveryCount := oldProcess.recoveryCount + 1
	log.Warnf("rebuilding process pool for script '%s' (rebuild attempt %d)", script, newRecoveryCount)

	pool, ok := h.processPools[script]
	if !ok {
		log.Warnf("pool for script '%s' disappeared during recovery. Aborting.", script)
		h.managerMutex.Unlock()
		return
	}
	poolSize := cap(pool)

	newPool := make(chan *ManagedHookProcess, poolSize)
	newLiveProcesses := make([]*ManagedHookProcess, 0, poolSize)

	// Kill old processes
	if oldProcesses, ok := h.allProcesses[script]; ok {
		for _, p := range oldProcesses {
			if p.cmd != nil && p.cmd.Process != nil {
				_ = p.cmd.Process.Kill()
			}
		}
	}

	// Create new processes, but DO NOT start monitors yet.
	for i := 0; i < poolSize; i++ {
		process, err := h.startNativeProcess(script, i, newRecoveryCount)
		if err != nil {
			log.WithError(err).Errorf("failed to start instance %d during pool rebuild. Pool size will be reduced.", i)
			continue
		}
		newPool <- process
		newLiveProcesses = append(newLiveProcesses, process)
	}

	// Update the master lists
	h.allProcesses[script] = newLiveProcesses
	h.processPools[script] = newPool

	// Release the lock BEFORE launching the new monitors.
	h.managerMutex.Unlock()

	// --- Post-Lock Operations ---

	// NOW, with the lock released, it is safe to start the monitors.
	for _, p := range newLiveProcesses {
		go h.monitorProcess(p)
	}

	log.Infof("successfully rebuilt process pool for script '%s' with %d instances (now at %d rebuilds).", script, len(newLiveProcesses), newRecoveryCount)
}

// stopAll terminates all managed hook processes.
func (h *HookProcessManager) stopAll() {
	h.managerMutex.Lock()
	defer h.managerMutex.Unlock()

	for _, processes := range h.allProcesses {
		for _, p := range processes {
			if p.cmd != nil {
				_ = p.stdin.Close()
				_ = p.stdout.Close()
				_ = p.stderr.Close()
				if p.cmd.Process != nil {
					_ = p.cmd.Process.Kill()
					// We don't need to wait here. The `cmd.Wait()` in the monitor will still
					// return, and its call to `recoverProcess` will be safely aborted
					// by the `h.ctx.Err() != nil` check.
				}
			}
		}
	}
	// Clear the maps.
	h.allProcesses = make(map[string][]*ManagedHookProcess)
	h.processPools = make(map[string]chan *ManagedHookProcess)
}

// isRecoverableError checks if an error indicates a crashed process that needs recovery.
func (h *HookProcessManager) isRecoverableError(err error) bool {
	if errors.Is(err, io.EOF) {
		// EOF means the process closed its stdout stream, likely by exiting,
		// before sending a complete newline-terminated response.
		return true
	}
	if err != nil {
		// "broken pipe" and "pipe is being closed" are common symptoms of a crashed process
		// when we try to write to its stdin.
		errMsg := err.Error()
		return strings.Contains(errMsg, "broken pipe") || strings.Contains(errMsg, "pipe is being closed")
	}
	return false
}
