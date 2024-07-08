package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		in   priorityQueue
		want int
	}{
		{
			name: "empty",
			in:   priorityQueue{},
			want: 0,
		},
		{
			name: "not empty",
			in: priorityQueue{
				&queueEntry{index: 0},
				&queueEntry{index: 0},
				&queueEntry{index: 0},
			},
			want: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.in.Len()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	tests := []struct {
		name string
		in   priorityQueue
		i, j int
		want bool
	}{
		{
			name: "less FCost",
			in: priorityQueue{
				&queueEntry{
					index: 0,
					cell: &Cell{
						GCost: 0,
						HCost: 3,
						FCost: 3,
					},
				},
				&queueEntry{
					index: 1,
					cell: &Cell{
						GCost: 1,
						HCost: 4,
						FCost: 5,
					},
				},
			},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "equal FCost, less HCost",
			in: priorityQueue{
				&queueEntry{
					index: 0,
					cell: &Cell{
						GCost: 2,
						HCost: 2,
						FCost: 4,
					},
				},
				&queueEntry{
					index: 1,
					cell: &Cell{
						GCost: 1,
						HCost: 3,
						FCost: 4,
					},
				},
			},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "invalid indexes",
			in: priorityQueue{
				&queueEntry{
					index: 0,
					cell: &Cell{
						GCost: 0,
						HCost: 3,
						FCost: 3,
					},
				},
				&queueEntry{
					index: 1,
					cell: &Cell{
						GCost: 1,
						HCost: 4,
						FCost: 5,
					},
				},
			},
			i:    1,
			j:    2,
			want: false,
		},
		{
			name: "nil entries",
			in: priorityQueue{
				&queueEntry{
					index: 0,
					cell: &Cell{
						GCost: 0,
						HCost: 3,
						FCost: 3,
					},
				},
				nil,
			},
			i:    0,
			j:    1,
			want: false,
		},
		{
			name: "nil cells",
			in: priorityQueue{
				&queueEntry{
					index: 0,
					cell:  nil,
				},
				&queueEntry{
					index: 1,
					cell: &Cell{
						GCost: 1,
						HCost: 4,
						FCost: 5,
					},
				},
			},
			i:    0,
			j:    1,
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.in.Less(test.i, test.j)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	entryI := &queueEntry{index: 0}
	entryJ := &queueEntry{index: 1}

	queue := priorityQueue{entryI, entryJ}

	// no effect if indexes are invalid
	queue.Swap(1, 2)
	assert.True(t, queue[0] == entryI)
	assert.True(t, queue[1] == entryJ)

	// valid swap
	queue.Swap(0, 1)
	assert.True(t, queue[0] == entryJ)
	assert.Equal(t, 0, queue[0].index)
	assert.True(t, queue[1] == entryI)
	assert.Equal(t, 1, queue[1].index)

	// one of the entries is nil
	queue = priorityQueue{&queueEntry{index: 0}, nil}
	queue.Swap(0, 1)
	assert.Nil(t, queue[0])
	assert.Equal(t, 1, queue[1].index)
}

func TestPriorityQueue_Push(t *testing.T) {
	queue := priorityQueue{}

	cell := &Cell{}
	queue.Push(cell)

	assert.Len(t, queue, 1)
	assert.Equal(t, cell, queue[0].cell)
	assert.Equal(t, 0, queue[0].index)

	// push nil
	queue.Push(nil)
	assert.Len(t, queue, 1)
}

func TestPriorityQueue_Pop(t *testing.T) {
	cell := &Cell{}
	queue := priorityQueue{}

	// pop from empty queue
	assert.Nil(t, queue.Pop())

	// pop from non-empty queue
	queue.Push(cell)

	got := queue.Pop()
	assert.True(t, cell == got)
	assert.Len(t, queue, 0)
}

func TestPriorityQueue_GetIndex(t *testing.T) {
	inQueue := &Cell{}
	outQueue := &Cell{}

	queue := priorityQueue{}
	queue.Push(inQueue)

	assert.Equal(t, 0, queue.GetIndex(inQueue))
	assert.Equal(t, -1, queue.GetIndex(outQueue))
	assert.Equal(t, -1, queue.GetIndex(nil))
}
