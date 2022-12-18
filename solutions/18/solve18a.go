package day18

import (
	"advent2022/solutions/common"
)

// We keep track of a 3D integer grid (which we can construct once we know the minimum and
// maximum values in all three dimensions); a negative value means that we've placed a rock
// at that location. Whenever we place a rock, all surrounding neighbors that do not yet have
// a rock increment their "pending" value; this represents the number of rocks next to this
// currently empty space. Then, if we later encounter an instruction to place a rock in this
// space, the total number of exposed sides will decrease with this "pending" value. The
// number of exposed sides is increased whenever we increase the "pending" value. In this
// way, we track the total number of exposed sides as we are adding rocks one by one, and
// we avoid having to iterate through the entire space after adding all rocks.

func processA(pos position, pending *grid) int {
	posIndex := pending.index(pos)
	sideDiff := 0

	if pending.values[posIndex] > 0 {
		sideDiff -= pending.values[posIndex]
	}

	pending.values[posIndex] = -1

	for _, offset := range offsets {
		neighborPos := addPositions(pos, offset)
		neighborIndex := pending.index(neighborPos)

		if pending.values[neighborIndex] >= 0 {
			pending.values[neighborIndex] += 1
			sideDiff += 1
		}
	}

	return sideDiff
}

func SolveA(lines []string) common.Solution {
	positions := make([]position, len(lines))
	bounds := createBounds()

	for index, line := range lines {
		positions[index] = parseLine(line, bounds)
	}

	bounds.widen()
	pending := createGrid(bounds)
	exposedSides := 0

	for _, pos := range positions {
		exposedSides += processA(pos, pending)
	}

	return common.ToIntSolution(exposedSides)
}
