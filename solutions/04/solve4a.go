package day04

import "strings"

// Not a lot of room for optimization here, simply split the line into its individual fields,
// parse those fields to integers, and then check whether one of the two resulting ranges is
// fully inside the other one.

func fullyContains(line string) bool {
	fields := strings.Split(line, ",")
	sections1 := toSections(fields[0])
	sections2 := toSections(fields[1])

	if sections1.start >= sections2.start {
		return sections1.end <= sections2.end
	} else {
		return sections2.end <= sections1.end
	}
}

func SolveA(lines []string) int {
	count := 0

	for _, line := range lines {
		if fullyContains(line) {
			count += 1
		}
	}

	return count
}
