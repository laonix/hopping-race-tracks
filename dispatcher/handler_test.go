package dispatcher

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/laonix/hopping-race-tracks/input"
)

func TestNewTestCaseHandler(t *testing.T) {
	tests := []struct {
		name string
		opts []testCaseHandlerOption
		want *testCaseHandler
	}{
		{
			name: "no options",
			opts: nil,
			want: &testCaseHandler{},
		},
		{
			name: "valid options",
			opts: []testCaseHandlerOption{
				withHandlerIn(make(chan *input.TestCase, 3)),
				withHandlerOut(make(chan string, 3)),
				withHandlerLogger(NewMockLogger(t)),
				withHandlerProcessor(NewMockProcessor(t)),
			},
			want: &testCaseHandler{
				in:        make(chan *input.TestCase, 3),
				out:       make(chan string, 3),
				log:       NewMockLogger(t),
				processor: NewMockProcessor(t),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := newTestCaseHandler(test.opts...)
			assert.NotNil(t, got)
			if test.want.in != nil {
				assert.NotNil(t, got.in)
			} else {
				assert.Nil(t, got.in)
			}
			if test.want.out != nil {
				assert.NotNil(t, got.out)
			} else {
				assert.Nil(t, got.out)
			}
			assert.Equal(t, test.want.log, got.log)
			assert.Equal(t, test.want.processor, got.processor)
		})
	}
}

func TestTestCaseHandler_Run(t *testing.T) {
	tests := []struct {
		name string
		on   func(processor *MockProcessor, logger *MockLogger)
		ctx  context.Context
		in   chan *input.TestCase
		out  chan string
		tc   *input.TestCase
		want string
		err  error
	}{
		{
			name: "valid test case",
			on: func(processor *MockProcessor, logger *MockLogger) {
				processor.EXPECT().Process(mock.AnythingOfType("*input.TestCase")).Return("result", nil)
			},
			ctx:  context.Background(),
			in:   make(chan *input.TestCase, 1),
			out:  make(chan string, 1),
			tc:   &input.TestCase{ID: 1},
			want: "result",
			err:  nil,
		},
		{
			name: "processor error",
			on: func(processor *MockProcessor, logger *MockLogger) {
				processor.EXPECT().Process(mock.AnythingOfType("*input.TestCase")).Return("", assert.AnError)
				logger.EXPECT().Error(assert.AnError, "failed to process test case", "id", 1)
			},
			ctx:  context.Background(),
			in:   make(chan *input.TestCase, 1),
			out:  make(chan string, 1),
			tc:   &input.TestCase{ID: 1},
			want: "",
			err:  assert.AnError,
		},
		{
			name: "context done",
			on: func(processor *MockProcessor, logger *MockLogger) {
				logger.EXPECT().Debug("stopping test case handler")

			},
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			in:   make(chan *input.TestCase, 1),
			out:  make(chan string, 1),
			tc:   &input.TestCase{ID: 1},
			want: "",
			err:  context.Canceled,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewMockProcessor(t)
			l := NewMockLogger(t)
			if test.on != nil {
				test.on(p, l)
			}

			h := newTestCaseHandler(
				withHandlerIn(test.in),
				withHandlerOut(test.out),
				withHandlerLogger(l),
				withHandlerProcessor(p),
			)

			go h.Run(test.ctx)
			test.in <- test.tc

			for {
				select {
				case <-time.After(1 * time.Millisecond):
					if test.err == nil {
						t.Errorf("expected result '%s', got nothing", test.want)
					}
					return
				case got := <-test.out:
					assert.Equal(t, test.want, got)
					return
				}
			}
		})
	}
}

func TestTestCaseHandler_Run_Closed_input_channel(t *testing.T) {
	p := NewMockProcessor(t)
	l := NewMockLogger(t)

	l.EXPECT().Debug("test case channel is closed")

	h := newTestCaseHandler(
		withHandlerIn(func() chan *input.TestCase {
			in := make(chan *input.TestCase, 1)
			close(in)
			return in
		}(),
		),
		withHandlerOut(make(chan string, 1)),
		withHandlerLogger(l),
		withHandlerProcessor(p),
	)

	go h.Run(context.Background())
	<-time.After(1 * time.Millisecond)
}
