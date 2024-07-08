package dispatcher

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/pathfinder"
)

func TestNewGridProcessor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "new grid processor",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			NewGridProcessor()
		})
	}
}

func TestGridProcessor_GetGrid(t *testing.T) {
	type input struct {
		rows      int
		cols      int
		obstacles []pathfinder.Obstacle
	}

	tests := []struct {
		name string
		in   input
		want *pathfinder.Grid
	}{
		{
			name: "valid grid",
			in: input{
				rows: 3,
				cols: 3,
				obstacles: []pathfinder.Obstacle{
					{X1: 1, Y1: 1, X2: 1, Y2: 1},
				},
			},
			want: &pathfinder.Grid{
				Rows: 3,
				Cols: 3,
				Cells: map[int]map[int]*pathfinder.Cell{
					0: {
						0: {X: 0, Y: 0, Available: true},
						1: {X: 1, Y: 0, Available: true},
						2: {X: 2, Y: 0, Available: true},
					},
					1: {
						0: {X: 0, Y: 1, Available: true},
						1: {X: 1, Y: 1, Available: false},
						2: {X: 2, Y: 1, Available: true},
					},
					2: {
						0: {X: 0, Y: 2, Available: true},
						1: {X: 1, Y: 2, Available: true},
						2: {X: 2, Y: 2, Available: true},
					},
				},
			},
		},
		{
			name: "empty grid",
			in: input{
				rows: 0,
				cols: 0,
			},
			want: nil,
		},
		{
			name: "invalid input",
			in: input{
				rows: -1,
				cols: -1,
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewGridProcessor()

			got := p.GetGrid(test.in.rows, test.in.cols, test.in.obstacles...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGridProcessor_GetPathfinder(t *testing.T) {
	type input struct {
		grid *pathfinder.Grid
		h    pathfinder.Heuristic
	}

	grid := &pathfinder.Grid{
		Rows: 3,
		Cols: 3,
		Cells: map[int]map[int]*pathfinder.Cell{
			0: {
				0: {X: 0, Y: 0, Available: true},
				1: {X: 1, Y: 0, Available: true},
				2: {X: 2, Y: 0, Available: true},
			},
			1: {
				0: {X: 0, Y: 1, Available: true},
				1: {X: 1, Y: 1, Available: false},
				2: {X: 2, Y: 1, Available: true},
			},
			2: {
				0: {X: 0, Y: 2, Available: true},
				1: {X: 1, Y: 2, Available: true},
				2: {X: 2, Y: 2, Available: true},
			},
		},
	}

	tests := []struct {
		name string
		in   input
		want pathfinder.Pathfinder
	}{
		{
			name: "valid pathfinder",
			in: input{
				grid: grid,
				h:    pathfinder.ChebyshevDistance,
			},
			want: &pathfinder.GridPathfinder{
				Grid:      grid,
				Heuristic: pathfinder.ChebyshevDistance,
			},
		},
		{
			name: "invalid grid",
			in: input{
				grid: nil,
				h:    pathfinder.ChebyshevDistance,
			},
			want: nil,
		},
		{
			name: "invalid heuristic",
			in: input{
				grid: grid,
				h:    nil,
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewGridProcessor()

			got := p.GetPathfinder(test.in.grid, test.in.h)
			if test.want != nil {
				assert.NotNil(t, got)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestGridProcessor_Process(t *testing.T) {
	tests := []struct {
		name string
		in   *input.TestCase
		want string
		err  error
	}{
		{
			name: "valid path",
			in: &input.TestCase{
				ID:       1,
				GridRows: 3,
				GridCols: 3,
				Start:    input.CellCoordinates{X: 0, Y: 0},
				End:      input.CellCoordinates{X: 2, Y: 2},
			},
			want: "Test case #1: Optimal solution takes 2 hops.",
			err:  nil,
		},
		{
			name: "no path",
			in: &input.TestCase{
				ID:       1,
				GridRows: 3,
				GridCols: 3,
				Start:    input.CellCoordinates{X: 0, Y: 0},
				End:      input.CellCoordinates{X: 2, Y: 2},
				Obstacles: []input.Obstacle{
					{X1: 1, Y1: 0, X2: 1, Y2: 2},
					{X1: 0, Y1: 1, X2: 2, Y2: 1},
				},
			},
			want: "Test case #1: No solution.",
			err:  nil,
		},
		{
			name: "invalid input",
			in:   nil,
			want: "",
			err:  errors.New("test case must be provided"),
		},
		{
			name: "failed to create grid",
			in: &input.TestCase{
				ID: 1,
			},
			want: "",
			err:  errors.New("failed to create grid"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewGridProcessor()

			got, err := p.Process(test.in)
			assert.Equal(t, test.want, got)
			if test.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
