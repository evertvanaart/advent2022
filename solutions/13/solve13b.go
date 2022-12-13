package day13

import (
	"advent2022/solutions/common"
	"fmt"
	"sort"
)

// Since we already implemented the full comparison logic in the A part, the B part is mostly
// a matter of using that comparison logic as a sort comparator. For finding the markers in the
// sorted list, I did not feel like implementing the logic needed to a deep comparison between
// elements, so I instead opted to just convert them all to strings (having already written the
// String() functions for debugging) and simply compare those strings. Definitely a lot of
// room for improvement on this one, since it is one of the slowest solutions so far.

var marker1 element = parseElement("[[2]]")
var marker2 element = parseElement("[[6]]")

func indexOf(elementStrings []string, targetElementString string) int {
	for index, elementString := range elementStrings {
		if elementString == targetElementString {
			return index + 1
		}
	}

	return -1
}

func SolveB(lines []string) common.Solution {
	lineElements := []element{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parsedElement := parseElement(line)
		lineElements = append(lineElements, parsedElement)
	}

	lineElements = append(lineElements, marker1)
	lineElements = append(lineElements, marker2)

	sort.Slice(lineElements, func(i, j int) bool {
		element1 := lineElements[i]
		element2 := lineElements[j]
		result := element1.compare(element2)
		return result < 0
	})

	lineStrings := make([]string, len(lineElements))

	for index, lineElement := range lineElements {
		lineStrings[index] = fmt.Sprint(lineElement)
	}

	marker1Index := indexOf(lineStrings, fmt.Sprint(marker1))
	marker2Index := indexOf(lineStrings, fmt.Sprint(marker2))
	return common.ToIntSolution(marker1Index * marker2Index)
}
