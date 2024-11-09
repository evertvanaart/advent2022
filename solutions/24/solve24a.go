package day24

import (
	"advent2022/solutions/common"
	"strings"
)

// We use a breath-first approach here, where we compute the set of positions that we can reach
// after T minutes for increasing values of T; we then use this set as the starting positions for
// the next step, immediately discarding any positions that are outside of the grid or inside a
// blizzard; for the remaining positions (i.e., the _valid_ positions after T minutes) we compute
// all five next positions (four possible directions, plus one "do nothing" option), and add those
// to the set of possible positions for T+1. Since we are only interested in the _minimum_ time it
// takes to reach the target, we can stop this iteration as soon as we reach the target. Since the
// set of valid positions never grows too large (even in the full input, it only contains around
// 1000 positions by the end), this approach is fairly quick, around 50ms.
//
// The other optimization worth mentioning here is that we don't actually move any of the
// blizzards. Rather, we first divide the blizzards into four grids, one for each direction, and
// whenever we need to check if a position in one of these grids contains a blizzard at time T, we
// subtract T * V from the position (where V is the blizzard's movement per minute), and check if
// the resulting position (wrapped to grid coordinates) contains a blizzard in the original grid.
// This saves us from having to individually update several thousand blizzards every step.

func SolveA(lines []string) common.Solution {
	coreLines := lines[1 : len(lines)-1]

	startCol := strings.Index(lines[0], ".") - 1
	endCol := strings.LastIndex(lines[len(lines)-1], ".") - 1

	blizzards := []*grid{
		parseGrid(coreLines, 'v', vector2d{1, 0}),
		parseGrid(coreLines, '>', vector2d{0, 1}),
		parseGrid(coreLines, '^', vector2d{-1, 0}),
		parseGrid(coreLines, '<', vector2d{0, -1}),
	}

	gridSize := blizzards[0].size
	source = vector2d{-1, startCol}
	target = vector2d{gridSize[0], endCol}

	positions := map[vector2d]bool{source: true}
	time := 0

	for len(positions) > 0 {
		nextPositions := map[vector2d]bool{}

		for pos := range positions {
			outPositions := step(time, pos, blizzards, gridSize)

			for _, outPosition := range outPositions {
				if outPosition == target {
					return common.ToIntSolution(time + 1)
				}

				nextPositions[outPosition] = true
			}
		}

		positions = nextPositions
		time += 1
	}

	return common.ToIntSolution(-1)
}
