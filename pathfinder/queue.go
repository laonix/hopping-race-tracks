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
// (in case of equality, comparison continues with HCost and finally with GCost).
func (pq *priorityQueue) Less(i, j int) bool {
	if (*pq)[i].cell == nil || (*pq)[j].cell == nil {
		return false
	}

	if (*pq)[i].cell.FCost == (*pq)[j].cell.FCost {
		if (*pq)[i].cell.HCost == (*pq)[j].cell.HCost {
			return (*pq)[i].cell.GCost < (*pq)[j].cell.GCost
		}

		return (*pq)[i].cell.HCost < (*pq)[j].cell.HCost
	}

	return (*pq)[i].cell.FCost < (*pq)[j].cell.FCost
}

// Swap swaps the elements under indexes i and j.
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

// Push adds a new element to the queue.
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	cell := x.(*Cell)
	entry := &queueEntry{cell: cell}
	entry.index = n
	*pq = append(*pq, entry)
}

// Pop removes the element with the highest priority from the queue.
func (pq *priorityQueue) Pop() interface{} {
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
	for i, entry := range *pq {
		if entry.cell == cell {
			return i
		}
	}

	return -1
}
