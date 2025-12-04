package main

import (
	"fmt"
	"log"
	"os"
)

type SolutionTest struct {
	FileName string
	Expected []string
}

type Solver struct {
	Solution Solution
	FileName string
	Test     SolutionTest
}

var solvers = []Solver{
	{Day01{}, "inputs/day01.txt", SolutionTest{"inputs/day01.example.txt", []string{"3"}}},
}

const errStr = "\033[0m[\033[1;31mERROR\033[0m]"

func main() {
	fmt.Println("\033[1;33m[Advent of Code 2025]\033[0m")

OUTER:
	for _, solver := range solvers {

		fmt.Printf("\n\033[0;34m== [\033[1m%s\033[0;34m] ==\033[0m\n", solver.Solution.Name())

		prefix := fmt.Sprintf("\033[1;33m[%s]\033[0m", solver.Solution.Name())
		logger := log.New(os.Stdout, prefix+" ", 0)

		// Test solution
		testLogger := log.New(os.Stderr, fmt.Sprintf("%s [TEST] ", prefix), 0)
		if err := solver.Solution.Init(solver.Test.FileName, testLogger); err != nil {
			logger.Printf("%s Could not init test: %s", errStr, err)
			continue
		}

		testSolutions, err := solver.Solution.Solve()
		if err != nil {
			logger.Printf("%s Could not solve test: %s", errStr, err)
			continue
		}

		if len(testSolutions) != len(solver.Test.Expected) {
			logger.Printf("Unexpected length of test solution: expected %d but got %d", len(solver.Test.Expected), len(testSolutions))
			continue
		}

		for i, result := range testSolutions {
			if result != solver.Test.Expected[i] {
				logger.Printf("Test %d failed: Expected \"%s\" but got \"%s\"", i+1, solver.Test.Expected[i], result)
				continue OUTER
			}
		}

		// Run actual solution
		if err := solver.Solution.Init(solver.FileName, log.New(os.Stderr, prefix+" ", 0)); err != nil {
			logger.Printf("%s Could not initialize: %s", errStr, err)
			continue
		}

		solutions, err := solver.Solution.Solve()
		if err != nil {
			logger.Printf("%s Could not solve: %s", errStr, err)
			continue
		}

		for i, res := range solutions {
			logger.Printf("Part %d: %s", i+1, res)
		}
	}
}
