package day09

import (
	"advent2022/solutions/common"
	"strings"
)

// The move() function moves the head and the tail in the given dimension (0 for left or right,
// 1 for up or down) with the given number of steps (positive or negative), and then pulls the
// tail along according to the movement rules. The function returns all new positions visited
// by the tail during this move; adding these tail positions to a set (or a map to true, since
// this is Go) allows us to track the number of positions visited at least once by the tail.

type rope struct {
	head position
	tail position
}

func createRope() *rope {
	return &rope{}
}

// Move the head and the tail, and return all new positions of the tail.
func (rope *rope) move(dim int, steps int) []position {
	rope.head[dim] += steps

	headPos := rope.head[dim]
	tailPos := rope.tail[dim]
	posDiff := headPos - tailPos

	dir := sign(steps)
	absDiff := dir * posDiff

	// If the head is less than two ahead of the tail (in the dimension and direction
	// of the move action), it means the tail did not move, so no new tail positions.
	if absDiff < 2 {
		return []position{}
	}

	newPosCount := absDiff - 1
	newPos := make([]position, newPosCount)
	otherDim := otherDimension(dim)

	// If the difference in the movement dimension is two or more, the tail is pulled
	// into the same row or column as the head in the other (non-movement) dimension.
	rope.tail[otherDim] = rope.head[otherDim]

	// Add all new tail positions
	for i := 0; i < newPosCount; i++ {
		rope.tail[dim] += dir
		newPos[i] = rope.tail
	}

	return newPos
}

func processA(line string, rope *rope, visited map[position]bool) {
	fields := strings.Split(line, " ")
	steps := common.ToInt(fields[1])
	var newPos []position = nil

	switch fields[0] {
	case "U":
		newPos = rope.move(1, -steps)
	case "D":
		newPos = rope.move(1, steps)
	case "L":
		newPos = rope.move(0, -steps)
	case "R":
		newPos = rope.move(0, steps)
	}

	for _, pos := range newPos {
		visited[pos] = true
	}
}

func SolveA(lines []string) common.Solution {
	visited := map[position]bool{{0, 0}: true}
	rope := createRope()

	for _, line := range lines {
		processA(line, rope, visited)
	}

	return common.ToIntSolution(len(visited))
}
