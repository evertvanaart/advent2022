package day04

import "strings"

// Very similar; the only thing worth noting is that checking for overlap between two ranges
// is much easier if we start from the negative case, i.e. check if the ranges do _not_ overlap.

func overlaps(line string) bool {
	fields := strings.Split(line, ",")
	sections1 := toSections(fields[0])
	sections2 := toSections(fields[1])

	return !(sections1.end < sections2.start || sections2.end < sections1.start)
}

func SolveB(lines []string) int {
	count := 0

	for _, line := range lines {
		if overlaps(line) {
			count += 1
		}
	}

	return count
}
