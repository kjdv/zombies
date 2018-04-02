package solvers

type Node interface {
	Next() []Node // subsequent nodes

	Wins() bool // is this node the win condition?

	Cost() int // the cost of reaching this node

	Heuristic() int // heuristic representing an estimation of the remaining cost to win, higher is better
	// todo: feels a little inappropriate to have the heuristic as part of the node, seperate this

	Key() interface{} // should return something that kan uniquely identify this node, for use in comparisons and map keys
}

type searchNode struct {
	current Node
	prev    *searchNode
}

func newSearchNode(start Node) searchNode {
	return searchNode{start, nil}
}

func (s searchNode) appended(head Node) searchNode {
	return searchNode{head, &s}
}

func (s searchNode) heuristic() int {
	return s.current.Heuristic()
}

func (s searchNode) flatten() []Node {
	if s.prev == nil {
		return []Node{s.current}
	}

	return append(s.prev.flatten(), s.current)
}

func (s searchNode) visited(n Node) bool {
	// todo: this is a linear search, consider a hash-set
	if s.current.Key() == n.Key() {
		return true
	} else if s.prev != nil {
		return s.prev.visited(n)
	}
	return false
}
