package main

import (
	day01 "advent2022/solutions/01"
	"advent2022/solutions/common"
	"fmt"
	"time"
)

type solver func([]string) int

func getSolver(task string) solver {
	switch task {
	case "01a":
		return day01.SolveA
	case "01b":
		return day01.SolveB
	default:
		panic(fmt.Sprintf("Unsupported task '%s'", task))
	}
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
