# zombies
Solve the zombie puzzle from https://www.youtube.com/watch?v=7yDmGnA8Hw0, comparing various algorithms.

To see the solution
```
$ go build
$ ./zombies
...
```

To see how the various algorithms perform:
```
$ go test -v
=== RUN   TestCompareSolvers
--- PASS: TestCompareSolvers (0.00s)
	compare_test.go:52: 
		algorithm:              nodes visited:    time:
		depht first search      85                45.284µs
		breadth first search    115               67.061µs
		best first search       10                10.818µs
		Dijkstra's              26                30.544µs
		A*                      10                15.212µs
		
PASS
```
