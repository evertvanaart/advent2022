package day12

/* -------------------------------- Constants ------------------------------- */

const processedHeightHigh = 1000
const processedHeightLow = -1

/* ---------------------------------- Types --------------------------------- */

type position [2]int
type queue map[position]bool

/* ---------------------------------- Grid ---------------------------------- */

type grid struct {
	values   []int
	rowCount int
	colCount int
}

func (grid *grid) set(pos position, value int) {
	grid.values[pos[0]*grid.colCount+pos[1]] = value
}

func (grid *grid) get(pos position) int {
	return grid.values[pos[0]*grid.colCount+pos[1]]
}

func (grid *grid) contains(pos position) bool {
	return pos[0] >= 0 && pos[0] < grid.rowCount && pos[1] >= 0 && pos[1] < grid.colCount
}

func (grid *grid) neighbors(pos position) []position {
	neighbors := []position{
		{pos[0] - 1, pos[1]},
		{pos[0] + 1, pos[1]},
		{pos[0], pos[1] - 1},
		{pos[0], pos[1] + 1},
	}

	return neighbors
}

/* --------------------------------- Parsing -------------------------------- */

func parse(lines []string) (*grid, position, position) {
	rowCount := len(lines)
	colCount := len(lines[0])
	gridSize := rowCount * colCount
	heights := make([]int, gridSize)
	index := 0

	var start *position = nil
	var end *position = nil

	for row, line := range lines {
		for col, char := range line {
			if start == nil && char == 'S' {
				start = &position{row, col}
				heights[index] = 0
			} else if end == nil && char == 'E' {
				end = &position{row, col}
				heights[index] = 25
			} else {
				height := int(char - 'a')
				heights[index] = height
			}

			index += 1
		}
	}

	return &grid{heights, rowCount, colCount}, *start, *end
}
