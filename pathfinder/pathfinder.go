package pathfinder

import (
	"container/heap"

	"github.com/pkg/errors"
)

// Heuristic is a function that estimates the cost of moving from cell a to cell b.
type Heuristic func(a, b *Cell) int

// GridPathfinder finds the shortest path between to cells in a given grid.
type GridPathfinder struct {
	grid      *Grid
	heuristic Heuristic
}

// NewGridPathfinder returns a new grid pathfinder with the given grid and heuristic function.
func NewGridPathfinder(grid *Grid, h Heuristic) *GridPathfinder {
	return &GridPathfinder{
		grid:      grid,
		heuristic: h,
	}
}

// FindPath returns the shortest path from the start cell to the end cell.
//
// The path is calculated using the A* algorithm.
func (pf *GridPathfinder) FindPath(start, finish *Cell) ([]*Cell, error) {
	if start == nil || finish == nil {
		return nil, errors.New("start and finish cells must be provided")
	}

	// get start point
	s := pf.grid.GetCell(start.X, start.Y)
	if s == nil {
		return nil, errors.New("start cell is out of grid")
	}

	// get finish point
	f := pf.grid.GetCell(finish.X, finish.Y)
	if f == nil {
		return nil, errors.New("finish cell is out of grid")
	}
	if !f.Available {
		return nil, errors.New("finish cell is not available")
	}

	// initialize the start cell
	s.GCost = 0
	s.HCost = pf.heuristic(s, f)
	s.FCost = s.GCost + s.HCost
	s.Speed = Velocity{X: 0, Y: 0}

	// initialize the open cells priority queue
	open := &priorityQueue{}
	heap.Init(open)
	heap.Push(open, s)

	for open.Len() > 0 {
		// get the cell with the lowest priority (e.g., the cell with the lowest FCost)
		// and mark it as closed
		current := heap.Pop(open).(*Cell)
		current.Open = false
		current.Closed = true

		// if the finish cell is reached, reconstruct the path and return it
		if current == f {
			return reconstructPath(current), nil
		}

		// calculate the cost of moving to the neighbors of the current cell;
		// in the case of Hopping Race game, the cost is the number of hops,
		// and it remains the same for all neighbors
		gCost := current.GCost + 1

		// evaluate neighbors of the current cell and push them to the open cells priority queue
		for _, neighbor := range pf.grid.GetNeighbors(current) {
			// prepare the neighbor cell for re-evaluation
			// if a new path to it is shorter than the previous one
			if gCost <= neighbor.GCost {
				if neighbor.Open {
					heap.Remove(open, open.GetIndex(neighbor))
				}
				neighbor.Open = false
				neighbor.Closed = false
			}

			// evaluate not visited cell
			if !neighbor.Open && !neighbor.Closed {
				neighbor.GCost = gCost
				neighbor.HCost = pf.heuristic(neighbor, f)
				neighbor.FCost = neighbor.GCost + neighbor.HCost
				neighbor.Speed = Velocity{X: neighbor.X - current.X, Y: neighbor.Y - current.Y}
				neighbor.Parent = current
				neighbor.Open = true
				heap.Push(open, neighbor)
			}
		}
	}

	// No path found
	return nil, nil
}

// reconstructPath returns the path from the start cell to the given cell.
func reconstructPath(cell *Cell) []*Cell {
	path := make([]*Cell, 0)

	for cell != nil {
		path = append(path, cell)
		cell = cell.Parent
	}

	return reversePath(path)
}

// reversePath reverses the given path.
func reversePath(path []*Cell) []*Cell {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}

	return path
}
