package day02

import (
	"advent2022/solutions/common"
	"fmt"
)

// Still hardcoding the scores, just different values than in the A part.

func scoreRoundB(line string) int {
	switch line {
	case "A X":
		return 3 // 0 + 3
	case "A Y":
		return 4 // 3 + 1
	case "A Z":
		return 8 // 6 + 2
	case "B X":
		return 1 // 0 + 1
	case "B Y":
		return 5 // 3 + 2
	case "B Z":
		return 9 // 6 + 3
	case "C X":
		return 2 // 0 + 2
	case "C Y":
		return 6 // 3 + 3
	case "C Z":
		return 7 // 6 + 1
	}

	panic(fmt.Sprintf("Unexpected line '%s'", line))
}

func SolveB(lines []string) common.Solution {
	totalScore := 0

	for _, line := range lines {
		totalScore += scoreRoundB(line)
	}

	return common.ToIntSolution(totalScore)
}
