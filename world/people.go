package world

import (
	"fmt"
	"math/bits"
)

type persons uint // bitmask

const (
	intern    persons = 1
	assistant persons = 2
	janitor   persons = 4
	professor persons = 8

	all  persons = 0xf
	none persons = 0x0
)

func (p persons) cost() int {
	// as if-else instead of switch makes this inlineable
	// check with 'go build -gcflags -m'
	if p&professor != none {
		return 10
	} else if p&janitor != none {
		return 5
	} else if p&assistant != none {
		return 2
	} else if p&intern != none {
		return 1
	}
	return 0
}

func (p persons) expand() []persons {
	result := make([]persons, 0, bits.OnesCount(uint(p)))

	for i := intern; i <= professor; i <<= 1 {
		if p&i != 0 {
			result = append(result, i)
		}
	}
	return result
}

func (p persons) String() string {
	ps := []string{}
	for _, person := range p.expand() {
		var s string
		switch person {
		case intern:
			s = "intern"
		case assistant:
			s = "assistant"
		case janitor:
			s = "janitor"
		case professor:
			s = "professor"
		}

		ps = append(ps, s)
	}
	return fmt.Sprint(ps)
}

// all possible permutions or 1 or 2 persons that can cross over
// computed once at startup
var groups []persons

func init() {
	groups = make([]persons, 0, 10)

	for p := none; p <= all; p++ {
		count := bits.OnesCount(uint(p))
		if count > 0 && count <= 2 {
			groups = append(groups, p)
		}
	}
}
