package world

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/klaasjacobdevries/zombies/solvers"
)

type side uint

const (
	danger side = 0
	safe   side = 1
)

func (s side) opposite() side {
	return s ^ 1
}

func (s side) String() string {
	if s == safe {
		return "safe"
	}
	return "danger"
}

const totalTime = 17

type state struct {
	this          side
	safe          persons // bitmask, these people are safe
	remainingTime int
}

func InitialState() state {
	return state{
		this:          danger,
		safe:          none,
		remainingTime: totalTime,
	}
}

func (s state) Next() []solvers.Node {
	result := []solvers.Node{}

	pos := s.safe
	if s.this == danger {
		pos ^= all
	}

	for _, group := range groups {
		if group.cost() <= s.remainingTime &&
			group&pos == group { // if all members of the group are on this side
			newState := state{
				this:          s.this.opposite(),
				safe:          s.safe ^ group,
				remainingTime: s.remainingTime - group.cost(),
			}

			result = append(result, newState)
		}
	}

	return result
}

func (s state) Wins() bool {
	return s.safe == all
}

func (s state) Cost() int {
	// the cost of reaching this state; i.e. time already spend
	return totalTime - s.remainingTime
}

func (s state) Heuristic() int {
	return heuristic(s)
}

func (s state) Key() interface{} {
	return [2]int{int(s.this), int(s.safe)}
}

func (s state) String() string {
	builder := strings.Builder{}
	tabber := tabwriter.NewWriter(&builder, 0, 0, 4, ' ', 0)

	fmt.Fprintln(tabber, "\nlantern:\t", s.this)
	fmt.Fprintln(tabber, "safe:\t", s.safe)
	fmt.Fprintln(tabber, "not safe:\t", s.safe^all)
	fmt.Fprintln(tabber, "remaining:\t", s.remainingTime)

	tabber.Flush()
	return builder.String()
}
