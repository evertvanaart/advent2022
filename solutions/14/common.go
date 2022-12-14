package day14

import (
	"advent2022/solutions/common"
	"fmt"
	"strings"
)

/* -------------------------------- Utilities ------------------------------- */

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/* -------------------------------- Constants ------------------------------- */

const (
	air  = 0
	sand = 1
	rock = 2
)

const (
	sourceRow = 0
	sourceCol = 500
)

/* -------------------------------- Position -------------------------------- */

type position struct {
	row int
	col int
}

func parsePosition(field string) *position {
	values := strings.Split(field, ",")
	col := common.ToInt(values[0])
	row := common.ToInt(values[1])
	return &position{row, col}
}

func (pos *position) getCandidates() []*position {
	return []*position{
		{pos.row + 1, pos.col},
		{pos.row + 1, pos.col - 1},
		{pos.row + 1, pos.col + 1},
	}
}

/* --------------------------------- Bounds --------------------------------- */

type bounds struct {
	minRow int
	minCol int
	maxRow int
	maxCol int
}

func createBounds(sourceRow int, sourceCol int) *bounds {
	return &bounds{sourceRow, sourceCol, sourceRow, sourceCol}
}

func (bounds *bounds) update(pos *position) {
	bounds.minRow = min(bounds.minRow, pos.row)
	bounds.maxRow = max(bounds.maxRow, pos.row)
	bounds.minCol = min(bounds.minCol, pos.col)
	bounds.maxCol = max(bounds.maxCol, pos.col)
}

/* ---------------------------------- Grid ---------------------------------- */
type grid struct {
	values   []int
	rowCount int
	colCount int
}

func createGrid(rowCount int, colCount int) *grid {
	gridSize := rowCount * colCount
	values := make([]int, gridSize)
	return &grid{values, rowCount, colCount}
}

func (grid *grid) setSand(pos *position) {
	index := pos.row*grid.colCount + pos.col
	grid.values[index] = sand
}

func (grid *grid) isOpen(pos *position) bool {
	index := pos.row*grid.colCount + pos.col
	return grid.values[index] == air
}

func (grid *grid) contains(pos *position) bool {
	return pos.row >= 0 && pos.row < grid.rowCount && pos.col >= 0 && pos.col < grid.colCount
}

func (grid *grid) countSand() int {
	count := 0

	for _, value := range grid.values {
		if value == sand {
			count += 1
		}
	}

	return count
}

func (grid *grid) applyPair(start *position, end *position, offset *position) {
	minRow := min(start.row, end.row) - offset.row
	maxRow := max(start.row, end.row) - offset.row
	minCol := min(start.col, end.col) - offset.col
	maxCol := max(start.col, end.col) - offset.col

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			index := row*grid.colCount + col
			grid.values[index] = rock
		}
	}
}

func (grid *grid) applyLine(positions []*position, offset *position) {
	for i := 0; i < len(positions)-1; i++ {
		grid.applyPair(positions[i], positions[i+1], offset)
	}
}

func (grid *grid) print() {
	charMap := map[int]string{0: ".", rock: "#", sand: "O"}

	for row := 0; row < grid.rowCount; row++ {
		fields := make([]string, grid.colCount)

		for col := 0; col < grid.colCount; col++ {
			index := row*grid.colCount + col
			fields[col] = charMap[grid.values[index]]
		}

		fmt.Println(strings.Join(fields, ""))
	}
}

/* --------------------------------- Parsing -------------------------------- */

func parse(line string, bounds *bounds) []*position {
	fields := strings.Split(line, " -> ")
	positions := make([]*position, len(fields))

	for index, field := range fields {
		position := parsePosition(field)
		positions[index] = position
		bounds.update(position)
	}

	return positions
}
