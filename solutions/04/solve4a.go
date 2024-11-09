package day04

import (
	"advent2022/solutions/common"
	"strings"
)

// Not a lot of room for optimization here, simply split the line into its individual fields,
// parse those fields to integers, and then check whether one of the two resulting ranges is
// fully inside the other one.

func (sections *sections) contains(other *sections) bool {
	return other.start >= sections.start && other.end <= sections.end
}

func fullyContains(line string) bool {
	fields := strings.Split(line, ",")
	sections1 := toSections(fields[0])
	sections2 := toSections(fields[1])

	return sections1.contains(sections2) || sections2.contains(sections1)
}

func SolveA(lines []string) common.Solution {
	count := 0

	for _, line := range lines {
		if fullyContains(line) {
			count += 1
		}
	}

	return common.ToIntSolution(count)
}
