package pathfinder

// ChebyshevDistance returns the Chebyshev distance between two cells.
//
// Chebyshev distance is a special case of Diagonal distance
// between two cells in a grid where horizontal, vertical, and diagonal moves are allowed,
// and the cost of diagonal moves is the same as the cost of horizontal or vertical moves.
// In this case, the common formula for Diagonal distance:
//
//	D * (dx + dy) + (D2 - 2 * D) * min(dx, dy),
//
//	where dx = |a.X - b.X|, dy = |a.Y - b.Y|,
//	D - cost of horizontal or vertical move, D2 - cost of diagonal move
//
// is simplified to:
//
//	max(dx, dy)
//
// because D = D2 = 1.
func ChebyshevDistance(a, b *Cell) int {
	if a == nil || b == nil {
		return 0
	}

	dx := a.X - b.X
	if dx < 0 {
		dx = -dx
	}

	dy := a.Y - b.Y
	if dy < 0 {
		dy = -dy
	}

	return max(dx, dy)
}
