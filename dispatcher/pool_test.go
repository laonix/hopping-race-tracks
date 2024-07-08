package dispatcher

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewWorkerPool(t *testing.T) {
	tests := []struct {
		name string
		in   context.Context
	}{
		{
			name: "valid context",
			in:   context.Background(),
		},
		{
			name: "nil context",
			in:   nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := newWorkerPool(test.in)
			assert.NotNil(t, got)
		})
	}
}

func TestWorkerPool_StartWorkers(t *testing.T) {
	tests := []struct {
		name    string
		workers func() []Worker
	}{
		{
			name: "valid workers",
			workers: func() []Worker {
				var workers []Worker
				for i := 0; i < 3; i++ {
					w := NewMockWorker(t)
					w.EXPECT().Run(mock.AnythingOfType("*context.cancelCtx"))
					workers = append(workers, w)
				}
				return workers
			},
		},
		{
			name: "empty workers",
			workers: func() []Worker {
				return nil
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := newWorkerPool(context.Background())
			p.startWorkers(test.workers()...)
			time.Sleep(1 * time.Millisecond) // make sure workers have started
		})
	}
}

func TestWorkerPool_Wait(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	w := NewMockWorker(t)
	w.EXPECT().Run(mock.AnythingOfType("*context.cancelCtx"))

	p := newWorkerPool(ctx)
	p.startWorkers(w)

	err := p.wait(context.Background())
	assert.NoError(t, err)
}

func TestWorkerPool_Wait_Cancel(t *testing.T) {
	w := NewMockWorker(t)
	w.EXPECT().Run(mock.AnythingOfType("*context.cancelCtx")).After(5 * time.Millisecond)

	p := newWorkerPool(context.Background())
	p.startWorkers(w)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	err := p.wait(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}
