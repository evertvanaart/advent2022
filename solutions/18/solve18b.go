package day18

import (
	"advent2022/solutions/common"
)

// No particularly clever ideas here; just fill in the rocks in the grid, and start flooding
// it with water. When we encounter a cell that still contains air, we change it to water, and
// add all six neighbors to the buffer. When the cell contains rock, we increase the number of
// exposed sides. Since rock cells can only be added to the buffer while flooding an air cell,
// each rock side will get added to the buffer X times, where X is the number of flooded (and
// thus, externally exposed) cells next to this rock cell, i.e. its number of externally ex-
// posed faces. The widen() call was originally added in the A part so that we wouldn't have
// to deal with edge cells, but it actually helps us in a different way here; adding a single-
// cell margin around the core area means that all rock spaces can be flooded from all sides.

const (
	air   = 0
	rock  = 1
	water = 2
)

func flood(g *grid) int {
	exposedSides := 0
	buffer := []position{{0, 0, 0}}
	index := 0

	for index < len(buffer) {
		pos := buffer[index]
		index += 1

		if !g.contains(pos) {
			continue
		}

		posIndex := g.index(pos)
		value := g.values[posIndex]

		if value == water {
			continue
		} else if value == rock {
			exposedSides += 1
			continue
		}

		g.values[posIndex] = water

		for _, offset := range offsets {
			neighborPos := addPositions(pos, offset)
			buffer = append(buffer, neighborPos)
		}
	}

	return exposedSides
}

func SolveB(lines []string) common.Solution {
	positions := make([]position, len(lines))
	bounds := createBounds()

	for index, line := range lines {
		positions[index] = parseLine(line, bounds)
	}

	bounds.widen()
	grid := createGrid(bounds)

	for _, pos := range positions {
		posIndex := grid.index(pos)
		grid.values[posIndex] = rock
	}

	exposedSides := flood(grid)
	return common.ToIntSolution(exposedSides)
}
