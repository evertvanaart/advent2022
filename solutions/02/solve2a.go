package day02

import (
	"advent2022/solutions/common"
	"fmt"
)

// Since there's only nine possible combinations per round, we might as well hardcode all
// possible scores. Based on some very superficial performance testing, this is actually
// faster than the "proper" approach (split the string, compute the hand score based on
// the second field and the round score by comparing both fields) by a factor of four or
// five. We could also use a map instead of a switch, but the switch is slightly faster.

func scoreRoundA(line string) int {
	switch line {
	case "A X":
		return 4 // 3 + 1
	case "A Y":
		return 8 // 6 + 2
	case "A Z":
		return 3 // 0 + 3
	case "B X":
		return 1 // 0 + 1
	case "B Y":
		return 5 // 3 + 2
	case "B Z":
		return 9 // 6 + 3
	case "C X":
		return 7 // 6 + 1
	case "C Y":
		return 2 // 0 + 2
	case "C Z":
		return 6 // 3 + 3
	}

	panic(fmt.Sprintf("Unexpected line '%s'", line))
}

func SolveA(lines []string) common.Solution {
	totalScore := 0

	for _, line := range lines {
		totalScore += scoreRoundA(line)
	}

	return common.ToIntSolution(totalScore)
}
