package solvers

import "sort"

type pqItem struct {
	item  searchNode
	value int
}

type priorityQueue struct {
	items    []pqItem
	valuator func(n searchNode) int
}

func newPriorityQueue(valuator func(n searchNode) int) *priorityQueue {
	return &priorityQueue{[]pqItem{}, valuator}
}

func (pq *priorityQueue) push(item searchNode) {
	val := pq.valuator(item)
	pq.pushWithValue(item, val)
}

func (pq *priorityQueue) pushWithValue(item searchNode, value int) {
	i := sort.Search(len(pq.items), func(idx int) bool {
		return pq.items[idx].value < value
	})

	new := pqItem{item, value}
	pq.items = append(pq.items[:i], append([]pqItem{new}, pq.items[i:]...)...)

}

func (pq *priorityQueue) pop() searchNode {
	item := pq.items[0]
	pq.items = pq.items[1:]
	return item.item
}

func (pq *priorityQueue) empty() bool {
	return len(pq.items) == 0
}

func (pq *priorityQueue) replace(n searchNode, oldValue int, newValue int) {
	// remove
	i := sort.Search(len(pq.items), func(idx int) bool {
		item := pq.items[idx]
		return item.item.current.Key() == n.current.Key() && item.value == oldValue
	})
	if i < 0 || i >= len(pq.items) {
		panic("can only replace existing values")
	}

	pq.items = append(pq.items[:i], pq.items[i+1:]...)

	// and re-insert
	pq.pushWithValue(n, newValue)
}

func heuristicValuator(n searchNode) int {
	return n.current.Heuristic()
}

func costValuator(n searchNode) int {
	// the priority queue prioritizes the highest value, so invert the cost
	return -n.current.Cost()
}
