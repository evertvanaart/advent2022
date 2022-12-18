package day17

import (
	"advent2022/solutions/common"
)

// For the A part at least, we can allocate the worst-case grid (i.e., if all 2022 rocks were to
// be stacked on top of each other) for faster memory access. After that, we just execute the steps
// as described, keeping track of the maximum rock height at all times. The only optimization worth
// mentioning is that we don't need to actually spawn the rock three full spaces above the highest
// rock; we can just spawn it directly above this highest rock, and perform four gust steps in
// a row; since we know that we cannot collide with other rocks during these four initial
// gusts, we can additionally use a simplified collision detection.

func SolveA(lines []string) common.Solution {
	gridHeight := maxRockNoA * avgRockHeight
	gridSize := gridHeight * gridWidth

	grid := &grid{make([]bool, gridSize), -1}
	gusts := parseGusts(lines[0])

	for rockNo := 0; rockNo < maxRockNoA; rockNo++ {
		rockIndex := rockNo % len(rocks)
		rock := &rocks[rockIndex]
		resolveRock(rock, grid, gusts)
	}

	return common.ToIntSolution(grid.maxHeight + 1)
}
