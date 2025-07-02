package webhook

import (
	"context"
	"sync"
)

// SimpleQueueWorker provides a basic job queue with a single worker goroutine.
type SimpleQueueWorker struct {
	jobs   chan func()
	wg     sync.WaitGroup
	cancel context.CancelFunc
}

// NewSimpleQueueWorker creates and starts a new queue worker.
func NewSimpleQueueWorker(queueSize int) *SimpleQueueWorker {
	ctx, cancel := context.WithCancel(context.Background())
	qw := &SimpleQueueWorker{
		jobs:   make(chan func(), queueSize),
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
				// Context cancelled, exit immediately.
				return
			}
		}
	}()
}

// Submit adds a job to the queue. It will drop the job if the queue is full.
func (qw *SimpleQueueWorker) Submit(job func()) {
	select {
	case qw.jobs <- job:
		// Job submitted
	default:
		// Queue is full, drop job.
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
