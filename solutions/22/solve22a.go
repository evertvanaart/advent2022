package day22

import (
	"advent2022/solutions/common"
)

// This A part is far from elegant, but it works well enough. The basic idea is that while
// parsing the grid, we keep track of the bounds, i.e. for each row we track the range of
// columns with non-empty cells (containing '.' or '#'), and for each column we similarly
// track the range of non-empty rows. Then, whenever a step would cause us to leave these
// bounds, we instead move to the opposite bound of the current row or column. The lack of
// elegance stems from the way the position is updated, which involves far too many switch
// statements; properly splitting the position struct into a position vector and a movement
// vector would allow us to greatly simplify this logic.

/* --------------------------------- Vector --------------------------------- */

func (pos *vector2d) moveA(g *grid, rowDiff int, colDiff int, steps int) {
	for step := 0; step < steps; step++ {
		nextRow := pos.row + rowDiff
		nextCol := pos.col + colDiff

		if rowDiff == 0 {
			nextRow, nextCol = g.wrapHorizontal(nextRow, nextCol)
		} else {
			nextRow, nextCol = g.wrapVertical(nextRow, nextCol)
		}

		nextIndex := nextRow*g.cols + nextCol

		if g.walls[nextIndex] {
			break
		}

		pos.row = nextRow
		pos.col = nextCol
	}
}

/* ---------------------------------- Grid ---------------------------------- */

func (g *grid) wrapHorizontal(row int, col int) (int, int) {
	if col < g.rowBounds[row][0] {
		return row, g.rowBounds[row][1] - 1
	} else if col >= g.rowBounds[row][1] {
		return row, g.rowBounds[row][0]
	}

	return row, col
}

func (g *grid) wrapVertical(row int, col int) (int, int) {
	if row < g.colBounds[col][0] {
		return g.colBounds[col][1] - 1, col
	} else if row >= g.colBounds[col][1] {
		return g.colBounds[col][0], col
	}

	return row, col
}

/* ------------------------------- Core logic ------------------------------- */

func (i *instruction) applyA(pos *vector2d, g *grid) {
	switch pos.dir {
	case right:
		pos.moveA(g, 0, 1, i.steps)
	case down:
		pos.moveA(g, 1, 0, i.steps)
	case left:
		pos.moveA(g, 0, -1, i.steps)
	case up:
		pos.moveA(g, -1, 0, i.steps)
	}

	if i.turn == 'L' {
		pos.turnLeft()
	} else if i.turn == 'R' {
		pos.turnRight()
	}
}

func SolveA(lines []string) common.Solution {
	instructions := parseInstructions(lines[len(lines)-1])
	grid := parseGrid(lines[:len(lines)-2])

	position := &vector2d{0, grid.rowBounds[0][0], right}

	for _, instruction := range instructions {
		instruction.applyA(position, grid)
	}

	result := 1000*(position.row+1) + 4*(position.col+1) + position.dir
	return common.ToIntSolution(result)
}
