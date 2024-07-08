package dispatcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestCaseDispatcher(t *testing.T) {
	tests := []struct {
		name string
		opts []TestCaseDispatcherOption
		want *TestCaseDispatcher
	}{
		{
			name: "no options",
			opts: nil,
			want: &TestCaseDispatcher{},
		},
		{
			name: "valid options",
			opts: []TestCaseDispatcherOption{
				WithDispatcherPipeSize(3),
				WithDispatcherPoolSize(3),
				WithDispatcherLogger(NewMockLogger(t)),
			},
			want: &TestCaseDispatcher{
				pipeSize: 3,
				poolSize: 3,
				log:      NewMockLogger(t),
			},
		},
		{
			name: "invalid options",
			opts: []TestCaseDispatcherOption{
				WithDispatcherPipeSize(-1),
				WithDispatcherPoolSize(-1),
			},
			want: &TestCaseDispatcher{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewTestCaseDispatcher(context.Background(), test.opts...)
			assert.NotNil(t, got)
			assert.Equal(t, test.want.pipeSize, got.pipeSize)
			assert.Equal(t, test.want.poolSize, got.poolSize)
			assert.Equal(t, test.want.log, got.log)
			if test.want.pipeSize > 0 {
				assert.NotNil(t, got.in)
				assert.NotNil(t, got.out)
			} else {
				assert.Nil(t, got.in)
				assert.Nil(t, got.out)
			}
			if test.want.poolSize > 0 {
				assert.NotNil(t, got.pool)
			} else {
				assert.Nil(t, got.pool)
			}
		})
	}
}

func TestWithDispatcherPipeSize(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "valid pipe size",
			in:   3,
			want: 3,
		},
		{
			name: "invalid pipe size",
			in:   -1,
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := &TestCaseDispatcher{}

			WithDispatcherPipeSize(test.in)(d)
			assert.NotNil(t, d.pipeSize)
			assert.Equal(t, test.want, d.pipeSize)

			if test.want > 0 {
				assert.NotNil(t, d.in)
				assert.NotNil(t, d.out)
			} else {
				assert.Nil(t, d.in)
				assert.Nil(t, d.out)
			}
		})
	}
}

func TestWithDispatcherPoolSize(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "valid pool size",
			in:   3,
			want: 3,
		},
		{
			name: "invalid pool size",
			in:   -1,
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := &TestCaseDispatcher{}

			WithDispatcherPoolSize(test.in)(d)
			assert.NotNil(t, d.poolSize)
			assert.Equal(t, test.want, d.poolSize)
		})
	}
}

func TestWithDispatcherLogger(t *testing.T) {
	l := NewMockLogger(t)

	d := &TestCaseDispatcher{}

	WithDispatcherLogger(l)(d)
	assert.NotNil(t, d.log)
	assert.Equal(t, l, d.log)
}
