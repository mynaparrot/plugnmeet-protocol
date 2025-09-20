package webhook

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

// SimpleQueueWorker provides a basic job queue with a single worker goroutine.
type SimpleQueueWorker struct {
	jobs   chan func()
	wg     sync.WaitGroup
	logger *logrus.Entry
	cancel context.CancelFunc
}

// NewSimpleQueueWorker creates and starts a new queue worker.
func NewSimpleQueueWorker(ctx context.Context, queueSize int, logger *logrus.Entry) *SimpleQueueWorker {
	ctx, cancel := context.WithCancel(ctx)
	qw := &SimpleQueueWorker{
		jobs:   make(chan func(), queueSize),
		logger: logger,
		cancel: cancel,
	}
	qw.start(ctx)
	return qw
}

// start launches the worker goroutine.
func (qw *SimpleQueueWorker) start(ctx context.Context) {
	qw.wg.Add(1)
	go func() {
		defer qw.wg.Done()
		for {
			select {
			case job, ok := <-qw.jobs:
				if !ok {
					// Channel closed, exit.
					return
				}
				job()
			case <-ctx.Done():
				// Context canceled, exit immediately.
				return
			}
		}
	}()
}

// Submit adds a job to the queue. It will drop the job if the queue is full.
func (qw *SimpleQueueWorker) Submit(job func()) {
	// First, check if the worker is still running.
	// This prevents a panic if Submit is called after StopGracefully or Kill.
	if qw.isStopped() {
		return
	}

	select {
	case qw.jobs <- job:
		// Job submitted
	default:
		// The Queue is full, drop the job to prevent blocking.
		qw.logger.Warnln("webhook queue is full, dropping job")
	}
}

// StopGracefully waits for all queued jobs to be processed before stopping the worker.
func (qw *SimpleQueueWorker) StopGracefully() {
	close(qw.jobs)
	qw.wg.Wait()
}

// Kill stops the worker immediately, dropping any unprocessed jobs in the queue.
func (qw *SimpleQueueWorker) Kill() {
	qw.cancel()
	qw.wg.Wait()
}

// isStopped checks if the worker's context has been cancelled.
func (qw *SimpleQueueWorker) isStopped() bool {
	return qw.cancel == nil || qw.jobs == nil
}
