package solvers

type container interface {
	push(n searchNode)
	pop() searchNode
	empty() bool
}

// solver, generic search algorihm. Use different actual algorithms by switching
// container implementations
func solve(c container, start Node) (bool, Solution) {
	c.push(newSearchNode(start))

	visited := 0 // track how many nodes we evaluated to find a solution
	for !c.empty() {
		snode := c.pop()
		node := snode.current
		visited += 1

		if node.Wins() {
			return true, Solution{snode.flatten(), visited}
		}

		for _, next := range node.Next() {
			if !snode.visited(next) { // don't evaluate the same state twice
				c.push(snode.appended(next))
			}
		}
	}

	return false, Solution{}
}

func Dephtfirst(start Node) (bool, Solution) {
	// depth first search is by using a stack as container
	return solve(&stack{}, start)
}

func Breadthfirst(start Node) (bool, Solution) {
	// breadth first uses a queue
	return solve(&queue{}, start)
}

func Bestfirst(start Node) (bool, Solution) {
	// Best first uses a priority queue, using a heuristic
	return solve(newPriorityQueue(heuristicValuator), start)
}

// it is possible to express Dijkstra's algorithm and A* using the simple solver above
// by just having a priority queue + overloaded push
// one minor difference is that the simple solve is greedy and returns on the first
// solution found; we don't care
type advancedContainer struct { // todo: better name
	pq   priorityQueue
	dist map[interface{}]int // track values / cost distances
}

func newAdvancedContainer(valuator func(searchNode) int) *advancedContainer {
	return &advancedContainer{
		pq:   *newPriorityQueue(valuator),
		dist: make(map[interface{}]int),
	}
}

func (ac *advancedContainer) push(node searchNode) {
	k := node.current.Key()
	d, found := ac.dist[k]
	value := ac.pq.valuator(node)
	if !found { // we haven't seen this node before
		ac.pq.pushWithValue(node, value)
		ac.dist[k] = value
	} else if value > d { // we have seen this node, but re-evaluate
		ac.pq.replace(node, d, value)
		ac.dist[k] = value
	}
}

func (ac *advancedContainer) pop() searchNode {
	return ac.pq.pop()
}

func (ac *advancedContainer) empty() bool {
	return ac.pq.empty()
}

func Dijkstra(start Node) (bool, Solution) {
	// Dijkstra searches by just taking the lowest cost
	return solve(newAdvancedContainer(costValuator), start)
}

func Astar(start Node) (bool, Solution) {
	// A* searches by taking the lowest cost + the lowest estimated remainng cost
	// In theory, same results but (most cases) less nodes visited
	valuator := func(n searchNode) int {
		return costValuator(n) + heuristicValuator(n)
	}

	return solve(newAdvancedContainer(valuator), start)
}
