package day09

import (
	"advent2022/solutions/common"
	"strings"
)

// The longer rope means that we can no longer resolve all steps of a single instruction in one
// go; we'll have to resolve the rope step by step and node by node, i.e. move the head, then loop
// through the other nodes, and move them if needed (i.e. if the distance between the current node
// and the preceding node is more than one in either direction). We stop this iteration as soon as
// we find any node that didn't move, and we only need to add the new tail position to the output
// map if we did not take this early return (i.e. all nodes moved).

const length = 10

type longrope struct {
	nodes [length]position
}

func (rope *longrope) resolve(index int) bool {
	diff0 := rope.nodes[index-1][0] - rope.nodes[index][0]
	diff1 := rope.nodes[index-1][1] - rope.nodes[index][1]
	sign0 := sign(diff0)
	sign1 := sign(diff1)
	abs0 := sign0 * diff0
	abs1 := sign1 * diff1

	if abs0 == 2 && abs1 == 2 {
		rope.nodes[index][0] += sign0
		rope.nodes[index][1] += sign1
		return true
	}

	if abs0 == 2 {
		rope.nodes[index][0] += sign0
		rope.nodes[index][1] += diff1
		return true
	}

	if abs1 == 2 {
		rope.nodes[index][1] += sign1
		rope.nodes[index][0] += diff0
		return true
	}

	return false
}

func (rope *longrope) move(dim int, step int) bool {
	rope.nodes[0][dim] += step

	for index := 1; index < length; index++ {
		moved := rope.resolve(index)

		if !moved {
			return false
		}
	}

	return true
}

func (rope *longrope) moveN(dim int, step int, steps int) []position {
	positions := []position{}

	for i := 0; i < steps; i++ {
		tailMoved := rope.move(dim, step)

		if tailMoved {
			positions = append(positions, rope.nodes[length-1])
		}
	}

	return positions
}

func processB(line string, rope *longrope, visited map[position]bool) {
	fields := strings.Split(line, " ")
	steps := common.ToInt(fields[1])
	var positions []position = nil

	switch fields[0] {
	case "U":
		positions = rope.moveN(1, -1, steps)
	case "D":
		positions = rope.moveN(1, 1, steps)
	case "L":
		positions = rope.moveN(0, -1, steps)
	case "R":
		positions = rope.moveN(0, 1, steps)
	}

	for _, pos := range positions {
		visited[pos] = true
	}
}

func SolveB(lines []string) common.Solution {
	visited := map[position]bool{{0, 0}: true}
	longrope := &longrope{}

	for _, line := range lines {
		processB(line, longrope, visited)
	}

	return common.ToIntSolution(len(visited))
}
