package world

import (
	"testing"
	"zombies/expects"
	"zombies/solvers"
)

func TestNextStates(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		initial state
		expect  []solvers.Node
	}{
		{
			state{danger, all ^ intern, 1},
			[]solvers.Node{
				state{
					this:          safe,
					safe:          all,
					remainingTime: 0,
				},
			},
		}, {
			state{danger, none, 3},
			[]solvers.Node{
				state{
					this:          safe,
					safe:          intern,
					remainingTime: 2,
				},
				state{
					this:          safe,
					safe:          assistant,
					remainingTime: 1,
				},
				state{
					this:          safe,
					safe:          intern | assistant,
					remainingTime: 1,
				},
			},
		}, {
			InitialState(),
			[]solvers.Node{
				state{
					this:          safe,
					safe:          intern,
					remainingTime: 16,
				},
				state{
					this:          safe,
					safe:          assistant,
					remainingTime: 15,
				},
				state{
					this:          safe,
					safe:          intern | assistant,
					remainingTime: 15,
				},
				state{
					this:          safe,
					safe:          janitor,
					remainingTime: 12,
				},

				state{
					this:          safe,
					safe:          intern | janitor,
					remainingTime: 12,
				},
				state{
					this:          safe,
					safe:          assistant | janitor,
					remainingTime: 12,
				},
				state{
					this:          safe,
					safe:          professor,
					remainingTime: 7,
				},
				state{
					this:          safe,
					safe:          intern | professor,
					remainingTime: 7,
				},
				state{
					this:          safe,
					safe:          assistant | professor,
					remainingTime: 7,
				},
				state{
					this:          safe,
					safe:          janitor | professor,
					remainingTime: 7,
				},
			},
		},
	}

	for _, tc := range testcases {
		actual := tc.initial.Next()
		expect.Equals(tc.expect, actual)
	}
}

func TestWins(t *testing.T) {
	expect := expects.New(t)

	testcases := []struct {
		state  state
		expect bool
	}{
		{
			InitialState(),
			false,
		}, {
			state{safe, all, 0},
			true,
		},
	}

	for _, tc := range testcases {
		actual := tc.state.Wins()
		expect.Equals(tc.expect, actual)
	}
}
