package day08

import "advent2022/solutions/common"

// This was the first entry this year for which I didn't get the correct answer on my first try.
// I initially stopped counting as soon as I encountered a tree shorter than the previous one,
// failing to consider that later trees might still be larger. This incorrect approach does
// produce the correct result for the sample input, but fails on the actual input.
//
// With that admission out of the way, the solution I ended up with is still fairly simple.
// For each row and column, start at the edge, and work your way to the other side, tracking
// the height of the largest tree encountered so far; whenever encountering a larger tree,
// increment the visibility score of the current field, and after processing all rows and
// columns in both directions, count the number of fields with a positive visibility score.
//
// There might a faster solution to this (one that does not involve checking every field up to
// four times), but the current solution is still pretty fast. Two optimizations worth mentioning:
//
// 1) Keeping the input as strings and using the raw byte values saves a lot of time otherwise
//    spent parsing everything to integers; this works since the byte values of the numeric
//    characters have the exact same ordering as the numeric values they represent.
//
// 2) Returning early when we find a tree of length 9 actually saves quite a bit of time,
//    reducing running time by around 50% in some basic performance tests.

func sweepA(grid *textGrid, visible *intGrid, iter *gridIterator) {
	var currentMax byte = 0

	for ; grid.contains(iter); iter.step() {
		value := grid.get(iter)

		if value > currentMax {
			visible.increment(iter)
			currentMax = value

			if currentMax == '9' {
				return
			}
		}
	}
}

func SolveA(lines []string) common.Solution {
	grid := createTextGrid(lines)
	visible := createIntGrid(grid.rowCount, grid.colCount)

	for row := 0; row < grid.rowCount; row++ {
		sweepA(grid, visible, createGridIterator(row, 0, dirRight))
		sweepA(grid, visible, createGridIterator(row, grid.colCount-1, dirLeft))
	}

	for col := 0; col < grid.colCount; col++ {
		sweepA(grid, visible, createGridIterator(0, col, dirDown))
		sweepA(grid, visible, createGridIterator(grid.rowCount-1, col, dirUp))
	}

	visibleCount := visible.countPositive()
	return common.ToIntSolution(visibleCount)
}
