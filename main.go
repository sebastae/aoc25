package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type SolverPart struct {
	Solve         func(lines []string, logger *DebugLogger) (int, error)
	ExampleResult int
}

type Solver struct {
	Parts       []SolverPart
	InputFile   string
	ExampleFile string
}

var solvers = map[string]Solver{
	"Day01": {Parts: []SolverPart{{SolveDay1Part1, 3}, {SolveDay1Part2, 6}}, InputFile: "inputs/day01.txt", ExampleFile: "inputs/day01.example.txt"},
	"Day02": {Parts: []SolverPart{{SolveDay2Part1, 1227775554}, {SolveDay2Part2, 4174379265}}, InputFile: "inputs/day02.txt", ExampleFile: "inputs/day02.example.txt"},
	"Day03": {Parts: []SolverPart{{Day03.SolvePart1, 357}, {Day03.SolvePart2, 3121910778619}}, InputFile: "inputs/day03.txt", ExampleFile: "inputs/day03.example.txt"},
}

const errStr = "\033[0m[\033[1;31mERROR\033[0m]"

type SolverResult struct {
	Result   int
	Pass     bool
	Duration time.Duration
}

func RunSolver(solver *Solver, name string) ([]*SolverResult, error) {
	fmt.Printf("\n\033[0;34m== [\033[1m%s\033[0;34m] ==\033[0m\n", name)

	prefix := fmt.Sprintf("[\033[1;34m%s\033[0m]", name)
	logger := NewDebugLogger(os.Stdout, prefix+" ", 0, false)

	if solver == nil {
		return nil, fmt.Errorf("solver is nil")
	}

	// Run parts
	testLogger := NewDebugLogger(os.Stderr, fmt.Sprintf("%s [TEST] ", prefix), 0, true)
	exampleLines, err := ReadFile(solver.ExampleFile)
	if err != nil {
		return nil, fmt.Errorf("could not read example file: %w", err)
	}

	lines, err := ReadFile(solver.InputFile)
	if err != nil {
		return nil, fmt.Errorf("could not read input file: %w", err)
	}

	results := make([]*SolverResult, 0, len(solver.Parts))
	for i, part := range solver.Parts {
		result := SolverResult{}
		// Run example / test
		res, err := part.Solve(exampleLines, testLogger)
		if err != nil {
			testLogger.Printf("\033[0;31mERROR\033[0m: error running example for part %d: %s", i+1, err)
			results = append(results, nil)
			continue
		} else {
			result.Pass = res == part.ExampleResult
			if result.Pass {
				testLogger.Printf("\033[0;32mSUCCESS\033[0m: part %d example completed with result { %d }", i+1, res)
			} else {
				testLogger.Printf("\033[0;33mFAIL\033[0m: part %d expected example result %d, got %d", i+1, part.ExampleResult, res)
			}
		}

		// Run actual solution
		startTime := time.Now()
		res, err = part.Solve(lines, logger)
		endTime := time.Now()
		dur := endTime.Sub(startTime)

		if err != nil {
			logger.Printf("\033[0;31mERROR\033[0m: part %d returned error after %s: %s", i+1, dur, err)
			results = append(results, nil)
			continue
		}

		logger.Printf("\033[0;32mSUCCESS\033[0m: part %d returned with result { %d } after %s", i+1, res, dur)

		result.Result = res
		result.Duration = dur
		results = append(results, &result)
	}

	return results, nil
}

func GetResultsIcon(results []*SolverResult) string {
	var anyFail, anyErr bool
	for _, res := range results {
		if res == nil {
			anyErr = true
		} else if !res.Pass {
			anyFail = true
		}
	}

	if anyErr {
		return "❌"
	} else if anyFail {
		return "⚠️"
	} else {
		return "✅"
	}
}

func main() {
	fmt.Println("\033[1;33m[Advent of Code 2025]\033[0m")
	results := make(map[string][]*SolverResult)

	if len(os.Args) > 1 {
		for _, name := range os.Args[1:] {
			solver, ok := solvers[name]
			if !ok {
				log.Default().Fatalf("Could not run solver: no such solver defined: \"%s\"", os.Args[1])
			}

			res, err := RunSolver(&solver, name)
			if err != nil {

			}

			results[name] = res
		}
	} else {
		for name, solver := range solvers {
			res, err := RunSolver(&solver, name)
			if err != nil {

			}

			results[name] = res
		}
	}

	// Print summary
	fmt.Printf("\n[Summary]\n")
	for name, partResults := range results {
		fmt.Printf("[\033[0;34m%s\033[0m][%s]\n", name, GetResultsIcon(partResults))
		for i, res := range partResults {
			resStr := "Failed"
			if res != nil {
				passStr := "\033[0;32mPASS\033[0m"
				if !res.Pass {
					passStr = "\033[0;33mFAIL\033[0m"
				}
				resStr = fmt.Sprintf("[%s] Completed with result { %d } in %s", passStr, res.Result, res.Duration.String())
			}
			fmt.Printf("  [%s][Part %d] %s\n", GetResultsIcon([]*SolverResult{res}), i+1, resStr)
		}

	}
}
