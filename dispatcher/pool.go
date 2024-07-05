package dispatcher

import (
	"context"
	"sync"
)

// worker is an interface for a worker that processes tasks.
type worker interface {
	run(ctx context.Context)
}

// workerPool is a pool of workers that processes tasks concurrently.
type workerPool struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// newWorkerPool creates a new worker pool with the provided context.
func newWorkerPool(ctx context.Context) *workerPool {
	p := &workerPool{}

	p.ctx, p.cancel = context.WithCancel(ctx)

	return p
}

// startWorkers starts the provided workers in the pool.
func (p *workerPool) startWorkers(workers ...worker) {
	p.wg.Add(len(workers))

	for _, w := range workers {
		go func(ctx context.Context, w worker) {
			defer p.wg.Done()

			w.run(ctx)
		}(p.ctx, w)
	}
}

// wait waits for all workers in the pool to finish processing tasks.
func (p *workerPool) wait(ctx context.Context) error {
	endCh := make(chan struct{})

	go func() {
		p.wg.Wait()

		endCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-endCh:
		return nil
	}
}

// stop stops the worker pool.
func (p *workerPool) stop() {
	p.cancel()
}
