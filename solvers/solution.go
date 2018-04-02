package solvers

import (
	"fmt"
	"strings"
)

type Solution struct {
	Path    []Node
	Visited int // the number of nodes that needed to be visited to find this
}

func (s Solution) String() string {
	builder := strings.Builder{}

	for idx, node := range s.Path {
		fmt.Fprintln(&builder, "Step ", idx)
		fmt.Fprintln(&builder, node)
	}

	fmt.Fprintln(&builder, "Visited: ", s.Visited)
	return builder.String()
}
