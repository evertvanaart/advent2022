package day08

import "math"

/* -------------------------------- textGrid -------------------------------- */

type textGrid struct {
	lines    []string
	rowCount int
	colCount int
}

func createTextGrid(lines []string) *textGrid {
	rowCount := len(lines)
	colCount := len(lines[0])
	return &textGrid{lines, rowCount, colCount}
}

func (grid *textGrid) get(iter *gridIterator) byte {
	return grid.lines[iter.row][iter.col]
}

func (grid *textGrid) contains(iter *gridIterator) bool {
	return iter.row >= 0 && iter.row < grid.rowCount && iter.col >= 0 && iter.col < grid.colCount
}

/* --------------------------------- intGrid -------------------------------- */

type intGrid struct {
	values   []int
	rowCount int
	colCount int
}

func createIntGrid(rowCount int, colCount int) *intGrid {
	return &intGrid{make([]int, rowCount*colCount), rowCount, colCount}
}

func (grid *intGrid) initialize(value int) {
	for index := range grid.values {
		grid.values[index] = value
	}
}

func (grid *intGrid) set(iter *gridIterator, value int) {
	grid.values[iter.row*grid.colCount+iter.col] = value
}

func (grid *intGrid) increment(iter *gridIterator) {
	grid.values[iter.row*grid.colCount+iter.col] += 1
}

func (grid *intGrid) multiply(iter *gridIterator, value int) {
	grid.values[iter.row*grid.colCount+iter.col] *= value
}

func (grid *intGrid) countPositive() int {
	count := 0

	for _, value := range grid.values {
		if value > 0 {
			count += 1
		}
	}

	return count
}

func (grid *intGrid) findMax() int {
	max := math.MinInt

	for _, value := range grid.values {
		if value > max {
			max = value
		}
	}

	return max
}

/* ------------------------------ gridIterator ------------------------------ */

const (
	dirUp    int = 0
	dirDown  int = 1
	dirLeft  int = 2
	dirRight int = 3
)

type gridIterator struct {
	row int
	col int
	dir int
}

func createGridIterator(row int, col int, dir int) *gridIterator {
	return &gridIterator{row, col, dir}
}

func (iter *gridIterator) step() {
	switch iter.dir {
	case dirUp:
		iter.row -= 1
	case dirDown:
		iter.row += 1
	case dirLeft:
		iter.col -= 1
	case dirRight:
		iter.col += 1
	}
}
