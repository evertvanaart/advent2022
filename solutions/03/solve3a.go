package day03

import (
	"advent2022/solutions/common"
	"fmt"
)

// Straightforward approach: Create a "set" from the characters in one of the two compartments,
// then for each character in the other compartment, check if it exists in the set of the first
// compartment; if so, convert the character to its priority value. Since we know that the input
// characters are limited to ASCII values, we can use an array with fixed length for our "set"
// (instead of a e.g. a map), which is a lot faster.

func solveLine(line string) int {
	halfLength := len(line) / 2
	compartment1 := line[:halfLength]
	compartment2 := line[halfLength:]

	characters1 := createCharacterArray(compartment1)

	for _, c := range compartment2 {
		if characters1[c] {
			return getPriority(c)
		}
	}

	panic(fmt.Sprintf("Failed to find common value: %s, %s", compartment1, compartment2))
}

func SolveA(lines []string) common.Solution {
	sum := 0

	for _, line := range lines {
		sum += solveLine(line)
	}

	return common.ToIntSolution(sum)
}
