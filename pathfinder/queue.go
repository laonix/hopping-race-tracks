package pathfinder

// priorityQueue implements heap.Interface
// to hold grid cells ordered by their costs descending.
type priorityQueue []*queueEntry

// queueEntry is a wrapper for a cell in the priority queue.
type queueEntry struct {
	cell  *Cell
	index int
}

// Len returns the number of elements in the queue.
func (pq *priorityQueue) Len() int {
	return len(*pq)
}

// Less reports whether the Cell under index i
// is less than the Cell under index j.
// The comparison is based on FCost of the cells
// (in case of equality, comparison continues with HCost).
func (pq *priorityQueue) Less(i, j int) bool {
	if i < 0 || j < 0 || i >= len(*pq) || j >= len(*pq) {
		return false
	}

	if (*pq)[i] == nil || (*pq)[j] == nil ||
		(*pq)[i].cell == nil || (*pq)[j].cell == nil {
		return false
	}

	if (*pq)[i].cell.FCost == (*pq)[j].cell.FCost {
		return (*pq)[i].cell.HCost < (*pq)[j].cell.HCost
	}

	return (*pq)[i].cell.FCost < (*pq)[j].cell.FCost
}

// Swap swaps the elements under indexes i and j.
func (pq *priorityQueue) Swap(i, j int) {
	if i < 0 || j < 0 || i >= len(*pq) || j >= len(*pq) {
		return
	}

	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	if (*pq)[i] != nil {
		(*pq)[i].index = i
	}
	if (*pq)[j] != nil {
		(*pq)[j].index = j
	}
}

// Push adds a new element to the queue.
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)

	cell, ok := x.(*Cell)
	if !ok {
		return
	}

	entry := &queueEntry{cell: cell}
	entry.index = n

	*pq = append(*pq, entry)
}

// Pop removes the element with the highest priority from the queue.
func (pq *priorityQueue) Pop() interface{} {
	if len(*pq) == 0 {
		return nil
	}

	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1
	*pq = old[:n-1]
	return entry.cell
}

// GetIndex returns the index of the given cell in the queue.
//
// If the cell is not found, -1 is returned.
func (pq *priorityQueue) GetIndex(cell *Cell) int {
	for _, entry := range *pq {
		if entry.cell == cell {
			return entry.index
		}
	}

	return -1
}
