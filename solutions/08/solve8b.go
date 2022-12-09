package day08

import "advent2022/solutions/common"

// The naive approach would be to iterate over every tree, and for each tree determine the viewing
// distance in all four directions; however, assuming a square grid with N trees to a side, this
// would result in a complexity of O(N^3), since we'd need to check up to 2N trees for every tree.
//
// To keep things down to O(N^2) complexity, we again sweep through all rows and columns in both
// directions, keeping track of the number of steps we've taken since the start (or end) of the
// row or column. In addition, we keep track of a "last seen" array, which contains for every
// possible tree length the step value at which we last saw a tree of that length _or higher_.
// Thus, the viewing distance for one tree in one direction is equal to the difference between
// that tree's step value (i.e., its distance to the edge in the given direction) minus the
// current "last seen" step value for that tree's length. We multiply the tree's total score
// by this viewing distance, and after processing all rows and columns in all four directions,
// we can obtain the maximum value from the score grid. Note that this naturally covers the
// literal edge cases; scores for trees at the edge of the grid will always be multiplied
// by zero (and thus set to zero) for at least one direction.

func sweepB(grid *textGrid, scores *intGrid, iter *gridIterator) {
	lastSeen := make([]int, 10)
	step := 0

	for ; grid.contains(iter); iter.step() {
		heightIndex := int(grid.get(iter) - '0')
		score := step - lastSeen[heightIndex]
		scores.multiply(iter, score)

		for index := 0; index <= heightIndex; index++ {
			lastSeen[index] = step
		}

		step += 1
	}
}

func SolveB(lines []string) common.Solution {
	grid := createTextGrid(lines)
	scores := createIntGrid(grid.rowCount, grid.colCount)
	scores.initialize(1)

	for row := 0; row < grid.rowCount; row++ {
		sweepB(grid, scores, createGridIterator(row, 0, dirRight))
		sweepB(grid, scores, createGridIterator(row, grid.colCount-1, dirLeft))
	}

	for col := 0; col < grid.colCount; col++ {
		sweepB(grid, scores, createGridIterator(0, col, dirDown))
		sweepB(grid, scores, createGridIterator(grid.rowCount-1, col, dirUp))
	}

	maxScore := scores.findMax()
	return common.ToIntSolution(maxScore)
}
