package day23

import (
	"advent2022/solutions/common"
)

// Almost identical to the A part, just track whether any of the elves actually moved during each
// step, and stop as soon as we find a step during which nobody moved. There's probably room for
// optimization here; as it stands, this is the slowest solution so far, with a runtime of almost
// one second. As mentioned in the A part, using a one-dimensional array instead of a set would
// be a possible improvement; even after almost one thousand steps, the grid size is less than
// 200 on either side, so a dense array could easily fit in memory.

func SolveB(lines []string) common.Solution {
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

	stepNo := 0

	for {
		ruleOffset := stepNo % len(rules)
		moved := process(elves, rules, ruleOffset)
		stepNo += 1

		if !moved {
			break
		}
	}

	return common.ToIntSolution(stepNo)
}
