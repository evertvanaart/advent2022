package day19

import (
	"advent2022/solutions/common"
)

// The final approach of the A part is fast enough that it scales to 32 minutes without any
// issues; processing the first three blueprints takes around 500ms on my reference machine.

func SolveB(lines []string) common.Solution {
	maxTime = 32
	product := 1

	for _, line := range lines[:3] {
		blueprint := parse(line)
		score := solveBlueprint(blueprint)
		product *= score
	}

	return common.ToIntSolution(product)
}
