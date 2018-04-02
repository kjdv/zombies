package main

import (
	"fmt"
	"os"
	"time"
	"zombies/solvers"
	"zombies/world"
)

type solverFunc func(solvers.Node) (bool, solvers.Solution)

func solve(sf solverFunc) bool {

	startTime := time.Now()

	start := world.InitialState()
	ok, solution := sf(start)

	endTime := time.Now()

	if !ok {
		fmt.Print("no solution found!\n")
		return false
	}

	fmt.Printf("solution found, took %v:\n", endTime.Sub(startTime))
	fmt.Print(solution)

	return true
}

func main() {
	if !solve(solvers.Bestfirst) {
		os.Exit(1)
	}
}
