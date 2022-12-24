package day23

import (
	"advent2022/solutions/common"
)

// Another straightforward simulation task, with not a lot of room for optimization. Most of my
// time was spent debugging due to my own lack of reading comprehension, e.g. with regards to
// how the return value is computed, or how elves that are completely separated from the rest
// do not move at all. Using a set for the positions of the elves is probably less efficient
// than using a regular array, so switching to an array could improve performance, although
// this would mean having to deal with a dynamically expanding grid.

func SolveA(lines []string) common.Solution {
	elves := map[position]bool{}

	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				pos := position{row, col}
				elves[pos] = true
			}
		}
	}

	rules := []rule{
		createRule(offsetNE, offsetN, offsetNW, offsetN),
		createRule(offsetSW, offsetS, offsetSE, offsetS),
		createRule(offsetNW, offsetW, offsetSW, offsetW),
		createRule(offsetSE, offsetE, offsetNE, offsetE),
	}

	for stepNo := 0; stepNo < 10; stepNo++ {
		ruleOffset := stepNo % len(rules)
		process(elves, rules, ruleOffset)
	}

	tl, br := boundingRect(elves)
	size := (br[0] - tl[0] + 1) * (br[1] - tl[1] + 1)
	return common.ToIntSolution(size - len(elves))
}
