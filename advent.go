package main

import (
	day01 "advent2022/solutions/01"
	"advent2022/solutions/common"
	"fmt"
	"os"
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
	args := os.Args[1:]

	if len(args) != 2 {
		common.PrintUsage()
		os.Exit(1)
	}

	task := args[0]
	input := args[1]

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
