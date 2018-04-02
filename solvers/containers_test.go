package solvers

import (
	"testing"

	"github.com/klaasjacobdevries/zombies/expects"
)

func TestContainers(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		c   container
		in  []int
		out []int
	}{
		{
			c:   &stack{},
			in:  []int{1},
			out: []int{1},
		}, {
			c:   &stack{},
			in:  []int{1, 2, 3},
			out: []int{3, 2, 1},
		},
		{
			c:   &queue{},
			in:  []int{1},
			out: []int{1},
		}, {
			c:   &queue{},
			in:  []int{1, 2, 3},
			out: []int{1, 2, 3},
		},
		{
			c:   newPriorityQueue(heuristicValuator),
			in:  []int{1},
			out: []int{1},
		}, {
			c:   newPriorityQueue(heuristicValuator),
			in:  []int{1, 2, 3},
			out: []int{3, 2, 1},
		},
		{
			c:   newPriorityQueue(heuristicValuator),
			in:  []int{5, 2, 6, 3, 1, 4},
			out: []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, tc := range testcases {
		expect.Equals(true, tc.c.empty())

		for _, in := range tc.in {
			tc.c.push(newSearchNode(testNode(in)))
		}

		for _, out := range tc.out {
			expect.Equals(false, tc.c.empty())
			expect.Equals(testNode(out), tc.c.pop().current)
		}
		expect.Equals(true, tc.c.empty())
	}
}

func TestPqReplace(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		in    []int
		node  int
		value int
		front int
	}{
		{
			in:    []int{1, 2, 3},
			node:  2,
			value: 4,
			front: 2,
		},
	}

	for _, tc := range testcases {
		c := newPriorityQueue(heuristicValuator)

		for _, i := range tc.in {
			node := testNode(i)
			c.pushWithValue(newSearchNode(node), i)
		}

		c.replace(newSearchNode(testNode(tc.node)), tc.node, tc.value)

		expect.Equals(tc.front, int(c.pop().current.(testNode)))
	}
}
