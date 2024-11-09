package day22

import (
	"advent2022/solutions/common"
	"strings"
	"unicode"
)

/* -------------------------------- Constants ------------------------------- */

const (
	right = 0
	down  = 1
	left  = 2
	up    = 3
)

/* -------------------------------- Utilities ------------------------------- */

func noSpace(r rune) bool {
	return !unicode.IsSpace(r)
}

/* --------------------------------- Vector --------------------------------- */

type vector2d struct {
	row int
	col int
	dir int
}

func (pos *vector2d) turnLeft() {
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

func (pos *vector2d) turnRight() {
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

/* ---------------------------------- Grid ---------------------------------- */

type grid struct {
	walls     []bool
	rows      int
	cols      int
	rowBounds [][2]int
	colBounds [][2]int
}

/* --------------------------------- Parsing -------------------------------- */

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
