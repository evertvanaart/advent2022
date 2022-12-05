package main

import (
	day01 "advent2022/solutions/01"
	day02 "advent2022/solutions/02"
	day03 "advent2022/solutions/03"
	day04 "advent2022/solutions/04"
	day05 "advent2022/solutions/05"
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
		"03a": day03.SolveA,
		"03b": day03.SolveB,
		"04a": day04.SolveA,
		"04b": day04.SolveB,
		"05a": day05.SolveA,
		"05b": day05.SolveB,
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
