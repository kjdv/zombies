package solvers

type stack []searchNode

func (s *stack) push(n searchNode) {
	*s = append(*s, n)
}

func (s *stack) pop() searchNode {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *stack) empty() bool {
	return len(*s) == 0
}
