package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {
	type input struct {
		rows, cols int
		obstacles  []Obstacle
	}

	tests := []struct {
		name string
		in   input
		want *Grid
	}{
		{
			name: "2x2 grid with no obstacles",
			in: input{
				rows: 2,
				cols: 2,
			},
			want: &Grid{
				Rows: 2,
				Cols: 2,
				Cells: map[int]map[int]*Cell{
					0: {
						0: {X: 0, Y: 0, Available: true},
						1: {X: 1, Y: 0, Available: true},
					},
					1: {
						0: {X: 0, Y: 1, Available: true},
						1: {X: 1, Y: 1, Available: true},
					},
				},
			},
		},
		{
			name: "3x3 grid with obstacles",
			in: input{
				rows: 3,
				cols: 3,
				obstacles: []Obstacle{
					{X1: 1, Y1: 0, X2: 1, Y2: 2},
					{X1: 0, Y1: 1, X2: 2, Y2: 1},
				},
			},
			want: &Grid{
				Rows: 3,
				Cols: 3,
				Cells: map[int]map[int]*Cell{
					0: {
						0: {X: 0, Y: 0, Available: true},
						1: {X: 1, Y: 0, Available: false},
						2: {X: 2, Y: 0, Available: true},
					},
					1: {
						0: {X: 0, Y: 1, Available: false},
						1: {X: 1, Y: 1, Available: false},
						2: {X: 2, Y: 1, Available: false},
					},
					2: {
						0: {X: 0, Y: 2, Available: true},
						1: {X: 1, Y: 2, Available: false},
						2: {X: 2, Y: 2, Available: true},
					},
				},
			},
		},
		{
			name: "invalid input",
			in: input{
				rows: 0,
				cols: 1,
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewGrid(test.in.rows, test.in.cols, test.in.obstacles...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGrid_GetCell(t *testing.T) {
	grid := NewGrid(3, 3)
	cell := grid.GetCell(1, 1)
	assert.Equal(t, &Cell{X: 1, Y: 1, Available: true}, cell)

	cell = grid.GetCell(3, 3)
	assert.Nil(t, cell)
}

func TestGrid_GetNeighbors(t *testing.T) {
	type input struct {
		x, y int
	}

	tests := []struct {
		name string
		grid *Grid
		cell input
		want []*Cell
	}{
		{
			name: "valid neighbors",
			grid: func() *Grid {
				grid := NewGrid(5, 5, []Obstacle{
					{X1: 1, Y1: 2, X2: 4, Y2: 3},
				}...)
				grid.GetCell(4, 0).Speed = Velocity{X: 0, Y: 0}
				grid.GetCell(4, 1).Closed = true
				grid.GetCell(3, 1).Speed = Velocity{X: -1, Y: 1}
				return grid
			}(),
			cell: input{3, 1},
			want: []*Cell{
				{X: 2, Y: 1, Available: true},
				{X: 1, Y: 1, Available: true},
			},
		},
		{
			name: "invalid input",
			grid: NewGrid(3, 3),
			cell: input{3, 3},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.grid.GetNeighbors(test.grid.GetCell(test.cell.x, test.cell.y))
			assert.ElementsMatch(t, test.want, got)
		})
	}
}
