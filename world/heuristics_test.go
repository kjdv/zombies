package world

import (
	"testing"

	"github.com/kjdv/zombies/expects"
)

func TestValueHeuristic(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		state state
		val   int
	}{
		{
			state: state{danger, all ^ professor, 11},
			val:   1,
		}, {
			state: state{safe, all, 10},
			val:   WIN_UTILITY,
		}, {
			state: state{danger, none, 5},
			val:   LOOSE_UTILITY,
		}, {
			state: state{safe, none, 17},
			val:   4,
		}, {
			state: state{safe, none, 11},
			val:   LOOSE_UTILITY,
		}, {
			state: state{danger, professor, 10},
			val:   3,
		}, {
			state: state{safe, all ^ professor, 11},
			val:   0,
		},
	}

	for _, tc := range testcases {
		actual := heuristic(tc.state)
		expect.Equals(tc.val, actual)
	}
}
