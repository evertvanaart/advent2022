package day24

import (
	"advent2022/solutions/common"
	"strings"
)

// Since our solution for the A part is already pretty fast and generic, the B part is simple;
// we essentially just solve the same problem three times, but with different start times, and
// with different start and finish positions. The second and third leg aren't any more compli-
// cated than the first one, so the total runtime is approximately three times that of the
// A part, around 150ms total.

func solve(startTime int, start vector2d, end vector2d, blizzards []*grid, size vector2d) int {
	positions := map[vector2d]bool{start: true}
	time := startTime

	for len(positions) > 0 {
		nextPositions := map[vector2d]bool{}

		for pos := range positions {
			outPositions := step(time, pos, blizzards, size)

			for _, outPosition := range outPositions {
				if outPosition == end {
					return time + 1
				}

				nextPositions[outPosition] = true
			}
		}

		positions = nextPositions
		time += 1
	}

	return -1
}

func SolveB(lines []string) common.Solution {
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
	time := 0

	time = solve(time, source, target, blizzards, gridSize)
	time = solve(time, target, source, blizzards, gridSize)
	time = solve(time, source, target, blizzards, gridSize)

	return common.ToIntSolution(time)
}
