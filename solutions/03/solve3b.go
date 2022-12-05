package day03

import (
	"advent2022/solutions/common"
	"fmt"
)

// Exact same idea, but now with three strings instead of two. Create a character set for the
// first two strings, then check which of the characters in the third string is in both sets.

func solveGroup(lines []string) int {
	characters1 := createCharacterArray(lines[0])
	characters2 := createCharacterArray(lines[1])

	for _, c := range lines[2] {
		if characters1[c] && characters2[c] {
			return getPriority(c)
		}
	}

	panic(fmt.Sprintf("Failed to find common value: %v", lines))
}

func SolveB(lines []string) common.Solution {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		sum += solveGroup(lines[i : i+3])
	}

	return common.ToIntSolution(sum)
}
