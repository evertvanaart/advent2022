package day12

import "advent2022/solutions/common"

// Since we went with a breath-first search in the A part, we can solve the B part simply by
// flipping a few signs and changing the exit condition. We now start at the endpoint, and cannot
// move down by more than one height difference in a single step. Processed nodes are assigned a
// negative height value, preventing them from being visited more than once. We return as soon
// as we find any node with height zero, i.e. with input character 'a'.

func SolveB(lines []string) common.Solution {
	heights, _, endPosition := parse(lines)
	currentQueue := queue{endPosition: true}
	currentDistance := 0

	for {
		nextQueue := queue{}

		for currentPosition := range currentQueue {
			currentHeight := heights.get(currentPosition)

			if currentHeight == 0 {
				return common.ToIntSolution(currentDistance)
			}

			heights.set(currentPosition, processedHeightLow)
			heightLimit := currentHeight - 1

			neighbors := heights.neighbors(currentPosition)

			for _, neighborPosition := range neighbors {
				if !heights.contains(neighborPosition) {
					continue
				}

				neighborHeight := heights.get(neighborPosition)

				if neighborHeight >= heightLimit {
					nextQueue[neighborPosition] = true
				}
			}
		}

		currentQueue = nextQueue
		currentDistance += 1
	}
}
