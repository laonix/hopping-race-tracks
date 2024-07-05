package dispatcher

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/pathfinder"
)

// processTestCase processes a single test case and returns the result.
//
// It initializes a new pathfinder with the grid and obstacles from the test case,
// finds the path from the start to the end cell, and returns the string representation of the result.
//
// processTestCase is of processor function type.
func processTestCase(in *input.TestCase) (string, error) {
	if in == nil {
		return "", errors.New("test case must be provided")
	}

	g := pathfinder.NewGrid(in.GridRows, in.GridCols, getObstacles(in.Obstacles)...)
	pf := pathfinder.NewGridPathfinder(g, pathfinder.ChebyshevDistance)

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
