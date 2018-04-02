package main

import (
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"
	"time"
	"zombies/expects"
	"zombies/solvers"
	"zombies/world"
)

func TestCompareSolvers(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		name   string
		solver solverFunc
	}{
		{"depht first search", solvers.Dephtfirst},
		{"breadth first search", solvers.Breadthfirst},
		{"best first search", solvers.Bestfirst},
		{"Dijkstra's", solvers.Dijkstra},
		{"A*", solvers.Astar},
	}

	builder := strings.Builder{}
	tabber := tabwriter.NewWriter(&builder, 0, 0, 4, ' ', 0)
	fmt.Fprintln(tabber, "\nalgorithm:\tnodes visited:\ttime:")

	for _, tc := range testcases {
		// property of go: dynamic dispatches are partially
		// computed at runtime and cached. For fair comparisons
		// do a 'warm-up' run
		tc.solver(world.InitialState())

		startTime := time.Now()

		start := world.InitialState()
		ok, solution := tc.solver(start)

		endTime := time.Now()

		expect.Equals(true, ok)
		expect.Equals(6, len(solution.Path))

		fmt.Fprintf(tabber, "%v\t%v\t%v\n", tc.name, solution.Visited, endTime.Sub(startTime))
	}

	tabber.Flush()
	t.Log(builder.String())
}
