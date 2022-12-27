package day22

import (
	"advent2022/solutions/common"
	"fmt"
)

// This was maybe the hardest question in the entire Advent of Code 2022. It's the last question
// I finished this year, and I don't like how this solution turned out. Like with the A part, the
// lack of proper vector-based calculations means that the code is riddled with ugly switch state-
// ments, and in general it feels like I've dramatically overthought the problem. There are several
// places - like the findMirror() function - which feel like they could have been one-liners if I
// could've only gotten my head around the mathematics, but since I was eager to wrap the AoC up
// for the year, I ended up using dumb workarounds more often than not.
//
// The basic idea is that we first create an ordered list of edges, which are the cells immediately
// outside of the cube blueprint, and keep track of the inner corners, i.e. the edges where two
// cube faces meet. Intuitively, we can see that the exit point (where we leave the blueprint)
// and the entry point (where we reappear) will be mirrored in one of these inner corners, i.e.
// the distance from the exit point to one of these inner corners is equal to the distance
// between the entry point and this same inner corner. After finding the correct mirrored
// edge point, we check which of the surrounding cells is non-empty, and use that cell to
// move back onto the blueprint (with the corresponding direction).
//
// This leaves the question of _which_ inner corner to use to find the correct mirrored edge.
// To answer this question, we assign a _reach_ to each inner corner, and specify that an exit
// point X should use corner C to mirror to entry point E if the distance between X and C is
// less than the reach. We compute this reach as (L / 2) - D, where L is the total circum-
// ference of the blueprint (for both sample and real input, this is 14 times the length),
// and D is the distance between the two _other_ inner corners. This is all based on some
// back-of-the-napkin equations, and while it does hold for the sample and main input
// (which both have three inner corners), it won't work for four inner coners.

/* ---------------------------------- Utils --------------------------------- */

func absDiff(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func wrap(index int, length int) int {
	if index < 0 {
		return length - 1
	} else if index >= length {
		return 0
	} else {
		return index
	}
}

/* ---------------------------------- Edge ---------------------------------- */

type edge struct {
	row   int
	col   int
	inner bool
	reach int
}

func findInnerCorners(edges []edge) []int {
	indices := []int{}

	for index := range edges {
		if edges[index].inner {
			indices = append(indices, index)
		}
	}

	return indices
}

func computeDistance(index1 int, index2 int, length int) int {
	distance1 := absDiff(index1, index2) + 1
	distance2 := length - distance1 - 1

	if distance1 < distance2 {
		return distance1
	} else {
		return distance2
	}
}

func computeReach(innerCornerIndices []int, index int, length int) int {
	otherIndex1 := innerCornerIndices[(index+1)%len(innerCornerIndices)]
	otherIndex2 := innerCornerIndices[(index+2)%len(innerCornerIndices)]
	distance := computeDistance(otherIndex1, otherIndex2, length)
	return (length / 2) - distance
}

func findEdge(edges []edge, row int, col int) int {
	for index, edge := range edges {
		if edge.row == row && edge.col == col {
			return index
		}
	}

	panic(fmt.Sprintf("No edge matching (%d, %d)", row, col))
}

func findMirror(edges []edge, startIndex int) int {
	steps := []int{-1, 1}

	for _, step := range steps {
		index := startIndex
		distance := 0

		for ; distance < len(edges); distance++ {
			index = wrap(index+step, len(edges))

			if edges[index].inner {
				break
			}
		}

		if edges[index].reach <= distance {
			continue
		}

		for ; distance >= 0; distance-- {
			index = wrap(index+step, len(edges))
		}

		return index
	}

	return -1
}

/* ---------------------------------- Grid ---------------------------------- */

func (g *grid) trace(start vector2d) []edge {
	startRow := start.row
	startCol := start.col
	currentPos := start
	edges := []edge{}
	length := 0

	for {
		if g.isOpen(currentPos.right()) {
			currentPos = currentPos.right()
		} else if !g.isOpen(currentPos.next()) {
			edges = append(edges, edge{currentPos.row, currentPos.col, true, 0})
			currentPos = currentPos.left()
			length += 2
		} else {
			edges = append(edges, edge{currentPos.row, currentPos.col, false, 0})
			currentPos = currentPos.next()
			length += 1
		}

		if currentPos.row == startRow && currentPos.col == startCol {
			break
		}
	}

	innerCorners := findInnerCorners(edges)
	edges[innerCorners[0]].reach = computeReach(innerCorners, 0, length)
	edges[innerCorners[1]].reach = computeReach(innerCorners, 1, length)
	edges[innerCorners[2]].reach = computeReach(innerCorners, 2, length)
	return edges
}

func (g *grid) isOpen(pos vector2d) bool {
	if pos.row < 0 || pos.row >= g.rows || pos.col < 0 || pos.col >= g.cols {
		return true
	}

	if pos.row < g.colBounds[pos.col][0] || pos.row >= g.colBounds[pos.col][1] {
		return true
	}

	if pos.col < g.rowBounds[pos.row][0] || pos.col >= g.rowBounds[pos.row][1] {
		return true
	}

	return false
}

/* --------------------------------- Vector --------------------------------- */

func (pos *vector2d) moveB(g *grid, steps int, edges []edge) {
	for step := 0; step < steps; step++ {
		nextRow := pos.row
		nextCol := pos.col
		nextDir := pos.dir

		switch pos.dir {
		case down:
			nextRow += 1
		case right:
			nextCol += 1
		case up:
			nextRow -= 1
		case left:
			nextCol -= 1
		}

		if nextRow < 0 || nextRow >= g.rows || nextCol < 0 || nextCol >= g.cols ||
			nextRow < g.colBounds[nextCol][0] || nextRow >= g.colBounds[nextCol][1] ||
			nextCol < g.rowBounds[nextRow][0] || nextCol >= g.rowBounds[nextRow][1] {
			nextRow, nextCol, nextDir = jump(g, nextRow, nextCol, nextDir, edges)
		}

		nextIndex := nextRow*g.cols + nextCol

		if g.walls[nextIndex] {
			break
		}

		pos.row = nextRow
		pos.col = nextCol
		pos.dir = nextDir
	}
}

func jump(g *grid, row int, col int, dir int, edges []edge) (int, int, int) {
	edgeIndex := findEdge(edges, row, col)

	if edges[edgeIndex].inner {
		next := vector2d{row, col, dir}.right()
		return next.row, next.col, next.dir
	}

	mirrorEdge := findMirror(edges, edgeIndex)
	edgeRow := edges[mirrorEdge].row
	edgeCol := edges[mirrorEdge].col

	if !g.isOpen(vector2d{edgeRow - 1, edgeCol, up}) {
		return edgeRow - 1, edgeCol, up
	} else if !g.isOpen(vector2d{edgeRow + 1, edgeCol, down}) {
		return edgeRow + 1, edgeCol, down
	} else if !g.isOpen(vector2d{edgeRow, edgeCol - 1, left}) {
		return edgeRow, edgeCol - 1, left
	} else if !g.isOpen(vector2d{edgeRow, edgeCol + 1, right}) {
		return edgeRow, edgeCol + 1, right
	}

	panic(fmt.Sprintf("No field next to (%d, %d)", edgeRow, edgeCol))
}

func (pos vector2d) next() vector2d {
	switch pos.dir {
	case right:
		return vector2d{pos.row, pos.col + 1, pos.dir}
	case down:
		return vector2d{pos.row + 1, pos.col, pos.dir}
	case left:
		return vector2d{pos.row, pos.col - 1, pos.dir}
	case up:
		return vector2d{pos.row - 1, pos.col, pos.dir}
	}

	panic(fmt.Sprintf("Invalid direction '%d'", pos.dir))
}

func (pos vector2d) right() vector2d {
	switch pos.dir {
	case right:
		return vector2d{pos.row + 1, pos.col, down}
	case down:
		return vector2d{pos.row, pos.col - 1, left}
	case left:
		return vector2d{pos.row - 1, pos.col, up}
	case up:
		return vector2d{pos.row, pos.col + 1, right}
	}

	panic(fmt.Sprintf("Invalid direction '%d'", pos.dir))
}

func (pos vector2d) left() vector2d {
	switch pos.dir {
	case right:
		return vector2d{pos.row - 1, pos.col, up}
	case down:
		return vector2d{pos.row, pos.col + 1, right}
	case left:
		return vector2d{pos.row + 1, pos.col, down}
	case up:
		return vector2d{pos.row, pos.col - 1, left}
	}

	panic(fmt.Sprintf("Invalid direction '%d'", pos.dir))
}

/* ------------------------------- Core logic ------------------------------- */

func (i *instruction) applyB(pos *vector2d, g *grid, edges []edge) {
	pos.moveB(g, i.steps, edges)

	if i.turn == 'L' {
		pos.turnLeft()
	} else if i.turn == 'R' {
		pos.turnRight()
	}
}

func SolveB(lines []string) common.Solution {
	grid := parseGrid(lines[:len(lines)-2])
	edges := grid.trace(vector2d{-1, grid.rowBounds[0][0], right})
	instructions := parseInstructions(lines[len(lines)-1])
	position := &vector2d{0, grid.rowBounds[0][0], right}

	for _, instruction := range instructions {
		instruction.applyB(position, grid, edges)
	}

	result := 1000*(position.row+1) + 4*(position.col+1) + position.dir
	return common.ToIntSolution(result)
}
