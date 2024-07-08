package dispatcher

import (
	"context"

	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/logger"
)

// TestCaseDispatcher controls the processing of test cases in a concurrent manner.
type TestCaseDispatcher struct {
	// in is the channel for incoming test cases.
	in chan *input.TestCase
	// out is the channel for outgoing results.
	out chan string
	// pipeSize is the size of the channels in and out.
	pipeSize int

	// pool is the worker pool that processes the test cases.
	pool *workerPool
	// poolSize is the number of workers in the pool.
	poolSize int

	log logger.Logger
}

// TestCaseDispatcherOption provides a way to configure the TestCaseDispatcher.
type TestCaseDispatcherOption func(d *TestCaseDispatcher)

// NewTestCaseDispatcher creates a new TestCaseDispatcher with the provided options.
//
// A newly created dispatcher starts the worker pool with the specified number of workers
// and the channels for incoming test cases and outgoing results.
func NewTestCaseDispatcher(ctx context.Context, opts ...TestCaseDispatcherOption) *TestCaseDispatcher {
	d := &TestCaseDispatcher{}

	for _, opt := range opts {
		opt(d)
	}

	d.startHandlers(ctx)

	return d
}

// WithDispatcherPipeSize sets the size of the channels in and out.
//
// Channels are created right after the pipe size is set.
func WithDispatcherPipeSize(pipeSize int) TestCaseDispatcherOption {
	return func(d *TestCaseDispatcher) {
		if pipeSize > 0 {
			d.pipeSize = pipeSize
		} else {
			d.pipeSize = 0
		}

		if d.pipeSize > 0 {
			d.in = make(chan *input.TestCase, d.pipeSize)
			d.out = make(chan string, d.pipeSize)
		}
	}
}

// WithDispatcherPoolSize sets the number of workers in the pool.
func WithDispatcherPoolSize(poolSize int) TestCaseDispatcherOption {
	return func(d *TestCaseDispatcher) {
		if poolSize > 0 {
			d.poolSize = poolSize
		} else {
			d.poolSize = 0
		}
	}
}

// WithDispatcherLogger sets the logger for the dispatcher.
func WithDispatcherLogger(log logger.Logger) TestCaseDispatcherOption {
	return func(d *TestCaseDispatcher) {
		d.log = log
	}
}

// Dispatch sends the test case to the input channel for processing.
func (d *TestCaseDispatcher) Dispatch(testCase *input.TestCase) {
	d.in <- testCase
}

// Results returns the channel for outgoing results.
func (d *TestCaseDispatcher) Results() <-chan string {
	return d.out
}

// Stop stops the dispatcher and the worker pool.
func (d *TestCaseDispatcher) Stop(ctx context.Context) {
	close(d.in)
	close(d.out)

	d.pool.stop()

	if err := d.pool.wait(ctx); err != nil {
		d.log.Warn("wait for pool close completed", "error", err)
		return
	}
}

// startHandlers starts the worker pool with the specified number of workers.
func (d *TestCaseDispatcher) startHandlers(ctx context.Context) {
	d.pool = startPool(
		ctx,
		d.poolSize,
		withHandlerIn(d.in),
		withHandlerOut(d.out),
		withHandlerProcessor(NewGridProcessor()),
		withHandlerLogger(d.log),
	)
}

// startPool creates and starts a new worker pool with the specified number of workers.
//
// The workers are created with the provided options and added to the pool.
func startPool(ctx context.Context, count int, opts ...testCaseHandlerOption) *workerPool {
	if count <= 0 {
		return nil
	}

	pool := newWorkerPool(ctx)

	var handlers []Worker
	for i := 0; i < count; i++ {
		handlers = append(handlers, newTestCaseHandler(opts...))
	}

	pool.startWorkers(handlers...)

	return pool
}
