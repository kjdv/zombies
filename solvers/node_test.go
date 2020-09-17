package solvers

import (
	"testing"

	"github.com/kjdv/zombies/expects"
)

type testNode int

func (tn testNode) Next() []Node {
	return nil
}

func (tn testNode) Wins() bool {
	return false
}

func (tn testNode) Cost() int {
	return 0
}

func (tn testNode) Heuristic() int {
	return int(tn)
}

func (tn testNode) Key() interface{} {
	return tn
}

func TestSearchNode(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		input     []int
		heuristic int
	}{
		{
			input:     []int{1, 2, 3},
			heuristic: 3,
		}, {
			input:     []int{2},
			heuristic: 2,
		},
	}

	for _, tc := range testcases {
		node := newSearchNode(testNode(tc.input[0]))

		for idx := 1; idx < len(tc.input); idx++ {
			node = node.appended(testNode(tc.input[idx]))
		}

		expect.Equals(tc.heuristic, node.heuristic())

		out := node.flatten()
		outInts := []int{}
		for _, i := range out {
			outInts = append(outInts, int(i.(testNode)))
		}

		expect.Equals(tc.input, outInts)
	}

}

func TestSearchNodeVisited(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		start   []int
		new     int
		visited bool
	}{
		{
			start:   []int{1},
			new:     2,
			visited: false,
		}, {
			start:   []int{1},
			new:     1,
			visited: true,
		}, {
			start:   []int{1, 2, 3, 4},
			new:     2,
			visited: true,
		}, {
			start:   []int{1, 2, 3, 4},
			new:     5,
			visited: false,
		},
	}

	for _, tc := range testcases {
		node := newSearchNode(testNode(tc.start[0]))

		for idx := 1; idx < len(tc.start); idx++ {
			node = node.appended(testNode(tc.start[idx]))
		}

		expect.Equals(tc.visited, node.visited(testNode(tc.new)))
	}
}
