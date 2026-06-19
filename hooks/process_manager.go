package hooks

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/v3/shell"
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
}

// HookProcessManager manages all the long-lived hook processes.
type HookProcessManager struct {
	// processPools holds a channel for each script, which acts as a pool of available processes.
	processPools map[string]chan *ManagedHookProcess
	// allProcesses holds all created processes for cleanup purposes.
	allProcesses []*ManagedHookProcess
	// managerMutex protects the process maps/slices during concurrent operations like recovery and shutdown.
	managerMutex sync.Mutex
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
		allProcesses: make([]*ManagedHookProcess, 0),
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

			for i := 0; i < poolSize; i++ {
				process, err := h.startNativeProcess(script, i)
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

// startNativeProcess boots up a true long-lived OS process.
func (h *HookProcessManager) startNativeProcess(script string, instanceId int) (*ManagedHookProcess, error) {
	log := h.log.WithFields(logrus.Fields{
		"hook_script": script,
		"instance_id": instanceId,
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
		cmd:        cmd,
		stdin:      stdin,
		stdout:     stdout,
		stderr:     stderr,
		reader:     bufio.NewReader(stdout),
		log:        log,
		instanceId: instanceId,
	}

	log.Info("native long-lived process instance started successfully")
	return process, nil
}

// executeHook acts as a dispatcher. It executes one-shot commands directly
// and uses the process pool for long-lived scripts.
func (h *HookProcessManager) executeHook(script HookScript, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	if script.IsOneShot {
		return h.executeOneShotCommand(script, jsonData, timeout, log)
	}
	return h.executePooledProcess(script.Script, jsonData, timeout, log)
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

// executePooledProcess executes a command using the long-lived process pool.
func (h *HookProcessManager) executePooledProcess(script string, jsonData json.RawMessage, timeout time.Duration, log *logrus.Entry) (json.RawMessage, error) {
	pool, ok := h.processPools[script]
	if !ok {
		// this can be the case where user define all one-shot scripts
		// in that case, we don't need to return error
		// we'll simply return original data
		return jsonData, nil
	}

	var process *ManagedHookProcess
	select {
	case process = <-pool:
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
		// This defer block ensures the process is always either returned or recovered.
		defer func() {
			if err != nil {
				log.WithError(err).Error("process execution failed, attempting to recover")
				go h.recoverProcess(script, process)
			} else {
				pool <- process
			}
		}()

		_, err = process.stdin.Write(append(jsonData, '\n'))
		if err != nil {
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
	if oldProcess.cmd != nil {
		_ = oldProcess.stdin.Close()
		_ = oldProcess.stdout.Close()
		_ = oldProcess.stderr.Close()
		if oldProcess.cmd.Process != nil {
			_ = oldProcess.cmd.Process.Kill()
			_, _ = oldProcess.cmd.Process.Wait()
		}
	}

	// Attempt to start a new process.
	newProcess, err := h.startNativeProcess(script, oldProcess.instanceId)
	if err != nil {
		log.WithError(err).Error("failed to recover process; pool size for this script will be reduced")
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

// stopAll terminates all managed hook processes.
func (h *HookProcessManager) stopAll() {
	h.managerMutex.Lock()
	defer h.managerMutex.Unlock()

	for _, p := range h.allProcesses {
		if p.cmd != nil {
			_ = p.stdin.Close()
			_ = p.stdout.Close()
			_ = p.stderr.Close()
			if p.cmd.Process != nil {
				_ = p.cmd.Process.Kill()
				_, _ = p.cmd.Process.Wait()
			}
		}
	}
	// Clear the lists.
	h.allProcesses = make([]*ManagedHookProcess, 0)
	h.processPools = make(map[string]chan *ManagedHookProcess)
}
