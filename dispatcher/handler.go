package dispatcher

import (
	"context"

	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/logger"
	"github.com/laonix/hopping-race-tracks/pathfinder"
)

// Processor is an interface for processing test cases.
type Processor interface {
	// GetGrid returns a new pathfinder grid initialized with the provided rows, columns, and obstacles.
	GetGrid(rows, cols int, obstacles ...pathfinder.Obstacle) *pathfinder.Grid

	// GetPathfinder returns a new pathfinder initialized with the provided grid and heuristic function.
	GetPathfinder(g *pathfinder.Grid, distance pathfinder.Heuristic) pathfinder.Pathfinder

	// Process processes the provided test case and returns the result.
	Process(*input.TestCase) (string, error)
}

// testCaseHandler handles the processing of a single test case.
//
// testCaseHandler implements the worker interface and is used by the worker pool to process test cases concurrently.
type testCaseHandler struct {
	// in is the channel for incoming test cases.
	in <-chan *input.TestCase
	// out is the channel for outgoing results.
	out chan<- string

	// processor is the function that processes the test case.
	processor Processor

	log logger.Logger
}

// testCaseHandlerOption provides a way to configure the testCaseHandler.
type testCaseHandlerOption func(h *testCaseHandler)

// newTestCaseHandler creates a new testCaseHandler with the provided options.
func newTestCaseHandler(opts ...testCaseHandlerOption) *testCaseHandler {
	h := &testCaseHandler{}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

// withHandlerProcessor sets the processor function for the test case handler.
func withHandlerProcessor(processor Processor) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.processor = processor
	}
}

// withHandlerIn sets the channel for incoming test cases.
func withHandlerIn(in <-chan *input.TestCase) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.in = in
	}
}

// withHandlerOut sets the channel for outgoing results.
func withHandlerOut(out chan<- string) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.out = out
	}
}

// withHandlerLogger sets the logger for the test case handler.
func withHandlerLogger(log logger.Logger) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.log = log
	}
}

// Run controls the processing of the incoming test cases.
func (h *testCaseHandler) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			h.log.Debug("stopping test case handler")
			return

		case testCase, ok := <-h.in:
			if !ok {
				h.log.Debug("test case channel is closed")
				return
			}

			select {
			case <-ctx.Done():
				h.log.Debug("stopping test case handler")
				return
			default:
				result, err := h.processor.Process(testCase)
				if err != nil {
					h.log.Error(err, "failed to process test case", "id", testCase.ID)
					continue
				}

				h.out <- result
			}
		}
	}
}
