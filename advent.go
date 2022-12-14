package main

import (
	day01 "advent2022/solutions/01"
	day02 "advent2022/solutions/02"
	day03 "advent2022/solutions/03"
	day04 "advent2022/solutions/04"
	day05 "advent2022/solutions/05"
	day06 "advent2022/solutions/06"
	day07 "advent2022/solutions/07"
	day08 "advent2022/solutions/08"
	day09 "advent2022/solutions/09"
	day10 "advent2022/solutions/10"
	day11 "advent2022/solutions/11"
	day12 "advent2022/solutions/12"
	day13 "advent2022/solutions/13"
	day14 "advent2022/solutions/14"
	"advent2022/solutions/common"
	"fmt"
	"time"
)

type solver func([]string) common.Solution

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
		"06a": day06.SolveA,
		"06b": day06.SolveB,
		"07a": day07.SolveA,
		"07b": day07.SolveB,
		"08a": day08.SolveA,
		"08b": day08.SolveB,
		"09a": day09.SolveA,
		"09b": day09.SolveB,
		"10a": day10.SolveA,
		"10b": day10.SolveB,
		"11a": day11.SolveA,
		"11b": day11.SolveB,
		"12a": day12.SolveA,
		"12b": day12.SolveB,
		"13a": day13.SolveA,
		"13b": day13.SolveB,
		"14a": day14.SolveA,
		"14b": day14.SolveB,
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
	fmt.Printf("Solution: %s\n", solution.String())
}
