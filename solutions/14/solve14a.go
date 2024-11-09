package day14

import "advent2022/solutions/common"

// Another conceptually straightforward task that just takes a while to implement according to
// specification. We first find the minimum and maximum row and column values, which allow us to
// allocate the full grid in one go, rather than use a (usually less performant) sparse array; we
// subtract the minimum row and column from all line coordinates in order to line the grid area up
// with the origin. Once we've created the grid, we simply fill in the rock lines according to the
// input, and then simulate the falling sand (trying the three candidate positions in the given
// order) until we find one that drops off any of the edges.

func simulateA(grid *grid, source *position) bool {
	sandPosition := position{source.row, source.col}

	for {
		candidatePositions := sandPosition.getCandidates()
		moved := false

		for _, candidatePosition := range candidatePositions {
			if !grid.contains(candidatePosition) {
				return false
			}

			if grid.isOpen(candidatePosition) {
				sandPosition.row = candidatePosition.row
				sandPosition.col = candidatePosition.col
				moved = true
				break
			}
		}

		if !moved {
			grid.setSand(&sandPosition)
			return true
		}
	}
}

func SolveA(lines []string) common.Solution {
	allPositions := make([][]*position, len(lines))
	bounds := createBounds(sourceRow, sourceCol)

	for index, line := range lines {
		linePositions := parse(line, bounds)
		allPositions[index] = linePositions
	}

	rowCount := bounds.maxRow - bounds.minRow + 1
	colCount := bounds.maxCol - bounds.minCol + 1
	offset := position{bounds.minRow, bounds.minCol}
	grid := createGrid(rowCount, colCount)

	for _, positions := range allPositions {
		grid.applyLine(positions, &offset)
	}

	sourcePosition := position{sourceRow - offset.row, sourceCol - offset.col}
	shouldContinue := true

	for shouldContinue {
		shouldContinue = simulateA(grid, &sourcePosition)
	}

	sandCount := grid.countSand()
	return common.ToIntSolution(sandCount)
}
