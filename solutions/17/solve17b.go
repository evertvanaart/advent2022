package day17

import (
	"advent2022/solutions/common"
)

// Intuitively, we can assume that the input will start to loop at some point; after all, there
// are only five possible rocks, and around 10,000 unique gusts, so we should encounter a loop
// (i.e., when both the selected rock and gust index are the same as they were at the start of
// a previous rock) after at most 50,000 rocks. If we assume that this loop will continue, i.e.
// every X rocks the maximum height will continue to increase by Y, we can skip ahead to the
// last partial loop before reaching 1,000,000,000,000, saving ourselves a lot of time.
//
// Unfortunately, skipping ahead like this on the very first loop we encounter produces the
// wrong answer, at least for the main input. This is probably because the state of the grid
// may differ between loops, e.g. the first rock after the loop may dissappear down a hole that
// wasn't there the first time. As an additional check, we therefore confirm that the position
// of the next block (relative to the maximum height) is the same as in the previous loop; if
// not, we keep looking for a loop that properly repeats itself. This still isn't foolproof
// I believe, but it at least produces the correct answer for this input.

func solveFromLoop(grid *grid, gusts *gusts,
	rockNoOffset int, loopRockCount int, loopHeight int) common.Solution {
	loops := ((maxRockNoB - rockNoOffset) / loopRockCount) - 1
	rockNo := rockNoOffset + loops*loopRockCount
	height := loopHeight * loops

	for ; rockNo < maxRockNoB; rockNo++ {
		rockIndex := rockNo % len(rocks)
		rock := &rocks[rockIndex]
		resolveRock(rock, grid, gusts)
	}

	return common.ToIntSolution(grid.maxHeight + height + 1)
}

type loopInfo struct {
	rockNo int
	height int
	pos    position
}

func (info loopInfo) positionMatches(origin position, height int) bool {
	return info.pos[0] == origin[0] && info.pos[1] == origin[1]-height
}

func SolveB(lines []string) common.Solution {
	// Overestimate the worst possible grid height
	gridHeight := 2 * len(rocks) * len(lines[0]) * avgRockHeight

	gridSize := gridHeight * gridWidth
	grid := &grid{make([]bool, gridSize), -1}
	gusts := parseGusts(lines[0])
	loopMap := map[[2]int]loopInfo{}

	for rockNo := 0; rockNo < maxRockNoB; rockNo++ {
		currentHeight := grid.maxHeight
		rockIndex := rockNo % len(rocks)
		loopKey := [2]int{rockIndex, gusts.index}
		prevInfo, hasLooped := loopMap[loopKey]

		rock := &rocks[rockIndex]
		origin := resolveRock(rock, grid, gusts)

		if hasLooped && prevInfo.positionMatches(*origin, grid.maxHeight) {
			loopRockCount := rockNo - prevInfo.rockNo
			loopHeight := currentHeight - prevInfo.height
			return solveFromLoop(grid, gusts, rockNo+1, loopRockCount, loopHeight)
		}

		pos := position{origin[0], origin[1] - grid.maxHeight}
		loopMap[loopKey] = loopInfo{rockNo, currentHeight, pos}
	}

	return common.ToIntSolution(grid.maxHeight + 1)
}
