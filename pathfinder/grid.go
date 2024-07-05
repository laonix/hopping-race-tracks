package pathfinder

const (
	minimalSpeed = -3
	maximalSpeed = 3
)

// Velocity represents the speed of a hopper.
type Velocity struct {
	// X is the horizontal speed.
	// It can be in the range from -3 to 3 inclusive.
	X int
	// Y is the vertical speed.
	// It can be in the range from -3 to 3 inclusive.
	Y int
}

// Cell represents a cell in a grid.
type Cell struct {
	// X is the horizontal coordinate of the cell.
	X int
	// Y is the vertical coordinate of the cell.
	Y int

	// Open indicates whether the cell has been checked and evaluated
	// during path finding algorithm iteration.
	Open bool
	// Closed indicates whether the cell has been processed
	// during path finding algorithm iteration.
	Closed bool
	// Available indicates whether the cell is available for hopping.
	Available bool

	// GCost is the cost of the path from the start cell to this cell.
	// In the case of Hopping Race game, it is the number of hops from the start cell to this cell.
	GCost int
	// HCost is the heuristic cost of the path from this cell to the finish cell.
	// HCost is calculated by the heuristic function
	// and represents the estimated number of 1-step hops from this cell to the finish cell
	// on a clear grid.
	HCost int
	// FCost is the sum of GCost and HCost.
	FCost int

	// Speed is the speed of the hopper when it reaches this cell.
	Speed Velocity

	// Parent is the cell from which the hopper reached this cell.
	Parent *Cell
}

// Obstacle represents an area in a grid that is not available for hopping.
type Obstacle struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// Grid represents an area in which hoppers can move.
type Grid struct {
	// Cells is a map of cells in the grid.
	Cells map[int]map[int]*Cell

	// Rows is the number of rows in the grid.
	Rows int
	// Cols is the number of columns in the grid.
	Cols int
}

// NewGrid returns a new grid with the given number of rows and columns
// and the specified obstacles.
func NewGrid(rows, cols int, obstacles ...Obstacle) *Grid {
	g := &Grid{
		Cells: make(map[int]map[int]*Cell),
		Rows:  rows,
		Cols:  cols,
	}

	for i := 0; i < rows; i++ {
		g.Cells[i] = make(map[int]*Cell)
		for j := 0; j < cols; j++ {
			g.Cells[i][j] = &Cell{
				X:         j,
				Y:         i,
				Available: true,
			}
		}
	}

	// place obstacles
	for _, o := range obstacles {
		for i := o.Y1; i <= o.Y2; i++ {
			for j := o.X1; j <= o.X2; j++ {
				g.Cells[i][j].Available = false
			}
		}
	}

	return g
}

// GetCell returns the cell at the specified coordinates.
func (g *Grid) GetCell(x, y int) *Cell {
	if x < 0 || x >= g.Cols || y < 0 || y >= g.Rows {
		return nil
	}

	return g.Cells[y][x]
}

// GetNeighbors returns the available neighboring cells of the specified cell.
//
// Neighboring cells are cells that are horizontally, vertically, or diagonally adjacent to the specified cell.
// The speed of the hopper is taken into account when determining the neighbors: the hopper can move in any direction
// keeping its speed and considering a possible velocity change by -1, 0, or 1
// (but gaining the speed not less than -3 and not higher than 3 in each direction).
//
// Cells are considered available if they are not obstacles and have not been closed.
func (g *Grid) GetNeighbors(cell *Cell) []*Cell {
	if cell == nil {
		return nil
	}

	// determine the range of possible speeds at this moment
	fromY := cell.Speed.Y - 1
	if fromY < minimalSpeed {
		fromY = minimalSpeed
	}

	toY := cell.Speed.Y + 1
	if toY > maximalSpeed {
		toY = maximalSpeed
	}

	fromX := cell.Speed.X - 1
	if fromX < minimalSpeed {
		fromX = minimalSpeed
	}

	toX := cell.Speed.X + 1
	if toX > maximalSpeed {
		toX = maximalSpeed
	}

	// find available neighbors
	var neighbors []*Cell

	for i := fromY; i <= toY; i++ {
		for j := fromX; j <= toX; j++ {
			// skip the current cell
			if i == 0 && j == 0 {
				continue
			}

			x := cell.X + j
			y := cell.Y + i

			if c := g.GetCell(x, y); c != nil && c.Available && !c.Closed {
				neighbors = append(neighbors, c)
			}
		}
	}

	return neighbors
}
