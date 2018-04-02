package world

import "math/bits"

// exact values don't matter much, as long as win is larger than anything else
// and loose is lower. Don't put them at the maximum for the data types though;
// they might appear in expressions causing overflow
const WIN_UTILITY = 1000
const LOOSE_UTILITY = -1000

func heuristic(s state) int {
	// estimate how 'good' this state is; as in how
	// close it is to getting to a solution.
	// in theory, a heuristic is fine if it always
	// underestimates the true cost

	if s.Wins() {
		return WIN_UTILITY
	}

	val := s.remainingTime
	minCost := 0

	// how much will it minimally cost to take all the in-danger
	// poeple to the safe side
	inDanger := s.safe ^ all
	// if the lantern is on the safe side, someone will have to move over
	if s.this == safe {
		fastest := persons(1 << uint(bits.TrailingZeros(uint(s.safe))))
		inDanger |= fastest
		minCost += fastest.cost()
	}

	minCost += inDanger.cost() // at least as much as the slowest

	// if more than two people have to move accross:
	if bits.OnesCount(uint(inDanger)) > 2 {
		upper := inDanger & (janitor | professor)
		lower := inDanger & (intern | assistant)
		if bits.OnesCount(uint(upper)) == 2 {
			minCost += lower.cost()
		} else {
			minCost += intern.cost()
		}
		minCost += intern.cost() // plus at least another round trip with the lantern
	}

	if minCost > s.remainingTime { // we're doomed
		return LOOSE_UTILITY
	}

	val -= minCost

	return val
}
