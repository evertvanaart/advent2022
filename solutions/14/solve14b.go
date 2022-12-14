package day14

import "advent2022/solutions/common"

// Very simular to the A part; once we've parsed all lines and computed the grid bounds, we
// simply add the bottom line and extend the bounds accordingly. Because of the way the sand
// falls, the horizontal width of the final sand pyramid can never be more than twice its
// height; we use this fact to extend the horizontal (column) bounds, allowing us to still
// use a fixed-size grid without fear of sand falling off. After that, it's simply a matter
// of running the simulation until a piece of sand settles in the source position.
//
// While this is a straightforward extension of the solution to the A part, it is pretty slow,
// around 20ms on my reference system. An alternate solution might involve starting from the
// full sand pyramid, and then subtracting all rock areas and empty spaces below rock areas;
// for example, a horizontal rock line of width 5 in the middle of the pyramid will probably
// have three empty (air) spaces on the row below it, and one on the row below that, i.e. an
// inverse air pyramid will naturally form below horizontal lines. Figuring out how this
// extends to vertical rock lines and intersecting lines seems like a lot of work though.

func simulateB(grid *grid, source *position) bool {
	sandPosition := position{source.row, source.col}

	for {
		candidatePositions := sandPosition.getCandidates()
		moved := false

		for _, candidatePosition := range candidatePositions {
			if grid.isOpen(candidatePosition) {
				sandPosition.row = candidatePosition.row
				sandPosition.col = candidatePosition.col
				moved = true
				break
			}
		}

		if !moved {
			grid.setSand(&sandPosition)
			return sandPosition != *source
		}
	}
}

func createBottom(bounds *bounds) []*position {
	bottomRow := bounds.maxRow + 2
	maxPyramidHeight := bottomRow - sourceRow
	bottomStart := &position{bottomRow, sourceCol - maxPyramidHeight}
	bottomEnd := &position{bottomRow, sourceCol + maxPyramidHeight}

	bounds.update(bottomStart)
	bounds.update(bottomEnd)

	return []*position{bottomStart, bottomEnd}
}

func SolveB(lines []string) common.Solution {
	allPositions := make([][]*position, len(lines))
	bounds := createBounds(sourceRow, sourceCol)

	for index, line := range lines {
		linePositions := parse(line, bounds)
		allPositions[index] = linePositions
	}

	bottom := createBottom(bounds)
	allPositions = append(allPositions, bottom)

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
		shouldContinue = simulateB(grid, &sourcePosition)
	}

	sandCount := grid.countSand()
	return common.ToIntSolution(sandCount)
}
