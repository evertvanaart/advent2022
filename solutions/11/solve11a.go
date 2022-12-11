package day11

import "advent2022/solutions/common"

// Not much to say about this one; the A part at least is one of those Advent of Code
// challenges where parsing the input and performing the steps according to specification
// is basically the whole challenge, and there's not a lot of room for optimization. The
// adjust function passed to processRound() was added after the fact to support part B.

func SolveA(lines []string) common.Solution {
	monkeys := []*monkey{}

	for lineNo := 0; lineNo < len(lines); lineNo += linesPerMonkey {
		monkeyLines := lines[lineNo : lineNo+linesPerMonkey]
		monkeys = append(monkeys, parseMonkey(monkeyLines))
	}

	adjust := func(value int) int {
		return value / 3
	}

	for round := 0; round < 20; round++ {
		processRound(monkeys, adjust)
	}

	output := determineOutput(monkeys)
	return common.ToInt64Solution(output)
}
