package pathfinder

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewGridPathfinder(t *testing.T) {
	tests := []struct {
		name      string
		grid      *Grid
		heuristic Heuristic
		want      Pathfinder
	}{
		{
			name:      "valid grid and heuristic",
			grid:      NewGrid(3, 3),
			heuristic: func(a, b *Cell) int { return 0 },
			want: &GridPathfinder{
				Grid:      NewGrid(3, 3),
				Heuristic: func(a, b *Cell) int { return 0 },
			},
		},
		{
			name:      "nil grid",
			grid:      nil,
			heuristic: func(a, b *Cell) int { return 0 },
			want:      nil,
		},
		{
			name:      "nil heuristic",
			grid:      NewGrid(3, 3),
			heuristic: nil,
			want:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewGridPathfinder(test.grid, test.heuristic)
			if test.want != nil {
				assert.NotNil(t, got)
				assert.Equal(t, test.want.(*GridPathfinder).Grid, got.(*GridPathfinder).Grid)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestGridPathfinder_FindPath(t *testing.T) {
	tests := []struct {
		name   string
		pf     *GridPathfinder
		start  *Cell
		finish *Cell
		want   []*Cell
		err    error
	}{
		{
			name: "valid path found",
			pf: &GridPathfinder{
				Grid:      NewGrid(3, 3),
				Heuristic: ChebyshevDistance,
			},
			start:  &Cell{X: 0, Y: 0},
			finish: &Cell{X: 2, Y: 2},
			want: []*Cell{
				{
					X:         0,
					Y:         0,
					Available: true,
					GCost:     0,
					HCost:     2,
					FCost:     2,
					Speed:     Velocity{X: 0, Y: 0},
					Closed:    true,
					Parent:    nil,
				},
				{
					X:         1,
					Y:         1,
					Available: true,
					GCost:     1,
					HCost:     1,
					FCost:     2,
					Speed:     Velocity{X: 1, Y: 1},
					Closed:    true,
					Parent: &Cell{
						X:         0,
						Y:         0,
						Available: true,
						GCost:     0,
						HCost:     2,
						FCost:     2,
						Speed:     Velocity{X: 0, Y: 0},
						Closed:    true,
						Parent:    nil,
					},
				},
				{
					X:         2,
					Y:         2,
					Available: true,
					GCost:     2,
					HCost:     0,
					FCost:     2,
					Speed:     Velocity{X: 1, Y: 1},
					Closed:    true,
					Parent: &Cell{
						X:         1,
						Y:         1,
						Available: true,
						GCost:     1,
						HCost:     1,
						FCost:     2,
						Speed:     Velocity{X: 1, Y: 1},
						Closed:    true,
						Parent: &Cell{
							X:         0,
							Y:         0,
							Available: true,
							GCost:     0,
							HCost:     2,
							FCost:     2,
							Speed:     Velocity{X: 0, Y: 0},
							Closed:    true,
							Parent:    nil,
						},
					},
				},
			},
			err: nil,
		},
		{
			name: "no path found",
			pf: &GridPathfinder{
				Grid: NewGrid(3, 3, []Obstacle{
					{X1: 1, Y1: 0, X2: 1, Y2: 2},
					{X1: 0, Y1: 1, X2: 2, Y2: 1},
				}...),
				Heuristic: ChebyshevDistance,
			},
			start:  &Cell{X: 0, Y: 0},
			finish: &Cell{X: 2, Y: 2},
			want:   nil,
			err:    nil,
		},
		{
			name: "nil input cell",
			pf: &GridPathfinder{
				Grid:      NewGrid(3, 3),
				Heuristic: ChebyshevDistance,
			},
			start:  nil,
			finish: &Cell{X: 2, Y: 2},
			want:   nil,
			err:    errors.New("start and finish cells must be provided"),
		},
		{
			name: "start cell out of grid",
			pf: &GridPathfinder{
				Grid:      NewGrid(3, 3),
				Heuristic: ChebyshevDistance,
			},
			start:  &Cell{X: 3, Y: 3},
			finish: &Cell{X: 2, Y: 2},
			want:   nil,
			err:    errors.New("start cell is out of grid"),
		},
		{
			name: "finish cell out of grid",
			pf: &GridPathfinder{
				Grid:      NewGrid(3, 3),
				Heuristic: ChebyshevDistance,
			},
			start:  &Cell{X: 0, Y: 0},
			finish: &Cell{X: 3, Y: 3},
			want:   nil,
			err:    errors.New("finish cell is out of grid"),
		},
		{
			name: "finish cell not available",
			pf: &GridPathfinder{
				Grid: NewGrid(3, 3, []Obstacle{
					{X1: 2, Y1: 2, X2: 2, Y2: 2},
				}...),
				Heuristic: ChebyshevDistance,
			},
			start:  &Cell{X: 0, Y: 0},
			finish: &Cell{X: 2, Y: 2},
			want:   nil,
			err:    errors.New("finish cell is not available"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.pf.FindPath(test.start, test.finish)

			if test.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want, got)
		})
	}
}
