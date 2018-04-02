package solvers

type queue []searchNode

func (q *queue) push(n searchNode) {
	*q = append(*q, n)
}

func (q *queue) pop() searchNode {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
