package day13

import "advent2022/solutions/common"

// We parse each line to a structure containing elements - which can be integers or nested
// lists - and then perform pair-wise comparisons between lines. This is probably overkill for
// the A part; looking at the input, most pairs are decided on their first or second value, so
// most of the time spent parsing the full lines is wasted. For the A part at least, scanning
// both lines from left to right and stopping when we encounter a difference probably would
// have been a lot faster, although maybe also more difficult to implement. Still, having
// access to the fully parsed lines will make the B part easier, so we'll keep it as is.

func compare(left string, right string) bool {
	leftElement := parseElement(left)
	rightElement := parseElement(right)
	result := leftElement.compare(rightElement)
	return result < 0
}

func SolveA(lines []string) common.Solution {
	index := 1
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		correctOrder := compare(lines[i], lines[i+1])

		if correctOrder {
			sum += index
		}

		index += 1
	}

	return common.ToIntSolution(sum)
}
