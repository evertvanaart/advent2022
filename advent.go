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
	day15 "advent2022/solutions/15"
	day16 "advent2022/solutions/16"
	day17 "advent2022/solutions/17"
	day18 "advent2022/solutions/18"
	day19 "advent2022/solutions/19"
	day20 "advent2022/solutions/20"
	day21 "advent2022/solutions/21"
	day22 "advent2022/solutions/22"
	day23 "advent2022/solutions/23"
	day24 "advent2022/solutions/24"
	day25 "advent2022/solutions/25"
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
		"15a": day15.SolveA,
		"15b": day15.SolveB,
		"16a": day16.SolveA,
		"16b": day16.SolveB,
		"17a": day17.SolveA,
		"17b": day17.SolveB,
		"18a": day18.SolveA,
		"18b": day18.SolveB,
		"19a": day19.SolveA,
		"19b": day19.SolveB,
		"20a": day20.SolveA,
		"20b": day20.SolveB,
		"21a": day21.SolveA,
		"21b": day21.SolveB,
		"22a": day22.SolveA,
		"22b": day22.SolveB,
		"23a": day23.SolveA,
		"23b": day23.SolveB,
		"24a": day24.SolveA,
		"24b": day24.SolveB,
		"25a": day25.SolveA,
	}

	solver, exists := solvers[task]

	if !exists {
		panic(fmt.Sprintf("Unsupported task '%s'", task))
	}

	return solver
}

const profileRuns = 20

func main() {
	task, input, flag := common.ParseArgs()

	day := common.ParseTask(task)
	lines := common.ReadLines(day, input)
	solver := getSolver(task)

	if flag == "--profile" {
		fmt.Printf("Profiling task '%s' using input '%s'\n", task, input)
		startTime := time.Now()

		for i := 0; i < profileRuns; i++ {
			solver(lines)
		}

		duration := time.Since(startTime)
		average := duration.Microseconds() / profileRuns
		fmt.Printf("Average: %d microseconds\n", average)
	} else {
		fmt.Printf("Running task '%s' on input '%s'\n", task, input)

		startTime := time.Now()
		solution := solver(lines)
		duration := time.Since(startTime)

		fmt.Printf("Elapsed time: %v\n", duration)
		fmt.Printf("Solution: %s\n", solution.String())
	}
}
