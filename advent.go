package main

import (
	day01 "advent2022/solutions/01"
	day02 "advent2022/solutions/02"
	"advent2022/solutions/common"
	"fmt"
	"time"
)

type solver func([]string) int

func getSolver(task string) solver {
	solvers := map[string]solver{
		"01a": day01.SolveA,
		"01b": day01.SolveB,
		"02a": day02.SolveA,
		"02b": day02.SolveB,
	}

	solver, exists := solvers[task]

	if !exists {
		panic(fmt.Sprintf("Unsupported task '%s'", task))
	}

	return solver
}

func main() {
	task, input := common.ParseArgs()
	day := common.ParseTask(task)
	lines := common.ReadLines(day, input)
	solver := getSolver(task)

	fmt.Printf("Running task '%s' on input '%s'\n", task, input)

	startTime := time.Now()
	solution := solver(lines)
	duration := time.Since(startTime)

	fmt.Printf("Elapsed time: %v\n", duration)
	fmt.Printf("Solution: %v\n", solution)
}
