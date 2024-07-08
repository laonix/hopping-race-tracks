package dispatcher

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/pathfinder"
)

// gridProcessor is an implementation of the Processor interface for the dispatcher.
type gridProcessor struct{}

// NewGridProcessor creates a new processor for the dispatcher.
func NewGridProcessor() Processor {
	return &gridProcessor{}
}

// GetGrid returns a new pathfinder grid initialized with the provided rows, columns, and obstacles.
func (p *gridProcessor) GetGrid(rows, cols int, obstacles ...pathfinder.Obstacle) *pathfinder.Grid {
	return pathfinder.NewGrid(rows, cols, obstacles...)
}

// GetPathfinder returns a new pathfinder initialized with the provided grid and heuristic function.
func (p *gridProcessor) GetPathfinder(g *pathfinder.Grid, distance pathfinder.Heuristic) pathfinder.Pathfinder {
	return pathfinder.NewGridPathfinder(g, distance)
}

// Process processes a single test case and returns the result.
//
// It initializes a new pathfinder with the grid and obstacles from the test case,
// finds the path from the start to the end cell, and returns the string representation of the result.
func (p *gridProcessor) Process(in *input.TestCase) (string, error) {
	if in == nil {
		return "", errors.New("test case must be provided")
	}

	g := p.GetGrid(in.GridRows, in.GridCols, getObstacles(in.Obstacles)...)
	if g == nil {
		return "", errors.New("failed to create grid")
	}

	pf := p.GetPathfinder(g, pathfinder.ChebyshevDistance)

	path, err := pf.FindPath(getCell(in.Start.X, in.Start.Y), getCell(in.End.X, in.End.Y))
	if err != nil {
		return "", errors.Wrap(err, "failed to find path")
	}
	if path != nil {
		return fmt.Sprintf("Test case #%d: Optimal solution takes %d hops.", in.ID, path[len(path)-1].GCost), nil
	} else {
		return fmt.Sprintf("Test case #%d: No solution.", in.ID), nil
	}
}

// getCell returns a new pathfinder cell with the provided coordinates.
func getCell(x, y int) *pathfinder.Cell {
	return &pathfinder.Cell{X: x, Y: y}
}

// getObstacles returns a slice of pathfinder obstacles from the provided input obstacles.
func getObstacles(inputObstacles []input.Obstacle) []pathfinder.Obstacle {
	var obstacles []pathfinder.Obstacle

	for _, o := range inputObstacles {
		obstacles = append(obstacles, pathfinder.Obstacle{
			X1: o.X1,
			X2: o.X2,
			Y1: o.Y1,
			Y2: o.Y2,
		})
	}

	return obstacles
}
