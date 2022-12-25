package day22

import (
	"advent2022/solutions/common"
	"fmt"
	"strings"
	"unicode"
)

func noSpace(r rune) bool {
	return !unicode.IsSpace(r)
}

const (
	right = 0
	down  = 1
	left  = 2
	up    = 3
)

type position struct {
	row int
	col int
	dir int
}

func (pos *position) turnLeft() {
	switch pos.dir {
	case right:
		pos.dir = up
	case down:
		pos.dir = right
	case left:
		pos.dir = down
	case up:
		pos.dir = left
	}
}

func (pos *position) turnRight() {
	switch pos.dir {
	case right:
		pos.dir = down
	case down:
		pos.dir = left
	case left:
		pos.dir = up
	case up:
		pos.dir = right
	}
}

func (pos *position) move(g *grid, rowDiff int, colDiff int, steps int) {
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

type grid struct {
	walls     []bool
	rows      int
	cols      int
	rowBounds [][2]int
	colBounds [][2]int
}

func (g *grid) print() {
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.rowBounds[row][0]; col++ {
			fmt.Print(" ")
		}

		for col := g.rowBounds[row][0]; col < g.rowBounds[row][1]; col++ {
			index := row*g.cols + col

			if g.walls[index] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}

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

func createBounds(length int, max int) [][2]int {
	bounds := make([][2]int, length)

	for index := range bounds {
		bounds[index] = [2]int{-1, max}
	}

	return bounds
}

func (g *grid) parseRow(row int, line string) {
	firstIndex := strings.IndexFunc(line, noSpace)
	lastIndex := strings.LastIndexFunc(line, noSpace)
	g.rowBounds[row] = [2]int{firstIndex, lastIndex + 1}

	for col := firstIndex; col <= lastIndex; col++ {
		char := line[col]

		if char == '#' {
			index := row*g.cols + col
			g.walls[index] = true
		}
	}
}

func parseGrid(lines []string) *grid {
	rows := len(lines)
	cols := 0

	for _, line := range lines {
		if len(line) > cols {
			cols = len(line)
		}
	}

	size := rows * cols
	walls := make([]bool, size)
	rowBounds := createBounds(rows, cols)
	colBounds := createBounds(cols, rows)

	grid := &grid{walls, rows, cols, rowBounds, colBounds}

	for row, line := range lines {
		grid.parseRow(row, line)
	}

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if col >= grid.rowBounds[row][0] && col < grid.rowBounds[row][1] {
				grid.colBounds[col][0] = row
				break
			}
		}

		for row := rows - 1; row >= 0; row-- {
			if col >= grid.rowBounds[row][0] && col < grid.rowBounds[row][1] {
				grid.colBounds[col][1] = row + 1
				break
			}
		}
	}

	return grid
}

/* ------------------------------ Instructions ------------------------------ */

type instruction struct {
	steps int
	turn  byte
}

func (i *instruction) apply(pos *position, g *grid) {
	switch pos.dir {
	case right:
		pos.move(g, 0, 1, i.steps)
	case down:
		pos.move(g, 1, 0, i.steps)
	case left:
		pos.move(g, 0, -1, i.steps)
	case up:
		pos.move(g, -1, 0, i.steps)
	}

	if i.turn == 'L' {
		pos.turnLeft()
	} else if i.turn == 'R' {
		pos.turnRight()
	}
}

func parseInstructions(line string) []instruction {
	instructions := []instruction{}
	groupStart := 0

	for index := 0; index < len(line); index++ {
		char := line[index]

		if char == 'L' || char == 'R' {
			steps := common.ToInt(line[groupStart:index])
			instructions = append(instructions, instruction{steps, char})
			groupStart = index + 1
		}
	}

	steps := common.ToInt(line[groupStart:])
	instructions = append(instructions, instruction{steps, 'X'})
	return instructions
}

func SolveA(lines []string) common.Solution {
	instructions := parseInstructions(lines[len(lines)-1])
	grid := parseGrid(lines[:len(lines)-2])

	position := &position{0, grid.rowBounds[0][0], right}

	for _, instruction := range instructions {
		instruction.apply(position, grid)
	}

	result := 1000*(position.row+1) + 4*(position.col+1) + position.dir
	return common.ToIntSolution(result)
}
