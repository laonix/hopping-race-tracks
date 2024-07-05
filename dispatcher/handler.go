package dispatcher

import (
	"context"

	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/logger"
)

// processor is a function that processes a test case.
type processor func(in *input.TestCase) (string, error)

// testCaseHandler handles the processing of a single test case.
//
// testCaseHandler implements the worker interface and is used by the worker pool to process test cases concurrently.
type testCaseHandler struct {
	// in is the channel for incoming test cases.
	in <-chan *input.TestCase
	// out is the channel for outgoing results.
	out chan<- string

	// processor is the function that processes the test case.
	processor processor

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

// withTestCaseHandlerProcessor sets the processor function for the test case handler.
func withTestCaseHandlerProcessor(processor processor) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.processor = processor
	}
}

// withTestCaseHandlerIn sets the channel for incoming test cases.
func withTestCaseHandlerIn(in <-chan *input.TestCase) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.in = in
	}
}

// withTestCaseHandlerOut sets the channel for outgoing results.
func withTestCaseHandlerOut(out chan<- string) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.out = out
	}
}

// withTestCaseHandlerLogger sets the logger for the test case handler.
func withTestCaseHandlerLogger(log logger.Logger) testCaseHandlerOption {
	return func(h *testCaseHandler) {
		h.log = log
	}
}

// run controls the processing of the incoming test cases.
func (h *testCaseHandler) run(ctx context.Context) {
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

			result, err := h.processor(testCase)
			if err != nil {
				h.log.Error(err, "failed to process test case", "id", testCase.ID)
				continue
			}

			h.out <- result
		}
	}
}
