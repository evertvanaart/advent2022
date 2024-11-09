package day12

import "advent2022/solutions/common"

// A breath-first search approach; we keep track of a set of active nodes (the queue), and each
// step we check which not-yet-processed nodes can be accessed from any of the active nodes, given
// the "at most one step up" rule; this set of accessible, not-yet-visited nodes becomes the queue
// in the next step. The current distance traveled is simply the number of such steps executed so
// far. Once we've processed a node we set its height to some value larger than the maximum height,
// which will automatically prevent future steps from returning to this node. Since we can only
// move in orthogonal steps of length 1, there's never any points in coming back to a node
// we've already visited, meaning we can return as soon as we reach the end.
//
// While this is still pretty fast, a guided depth-first approach would probably be more
// efficient, i.e. try to move towards the endpoint first, and only try routes that require
// moving away from the endpoint if there are no better options; this approach would essentially
// produce something similar to A*. On the other hand, this breath-first approach can be applied
// to the B part with minimal changes, which is not true for a depth-first approach.

func SolveA(lines []string) common.Solution {
	heights, startPosition, endPosition := parse(lines)
	currentQueue := queue{startPosition: true}
	currentDistance := 0

	for {
		nextQueue := queue{}

		for currentPosition := range currentQueue {
			if currentPosition == endPosition {
				return common.ToIntSolution(currentDistance)
			}

			currentHeight := heights.get(currentPosition)
			heights.set(currentPosition, processedHeightHigh)
			heightLimit := currentHeight + 1

			neighbors := heights.neighbors(currentPosition)

			for _, neighborPosition := range neighbors {
				if !heights.contains(neighborPosition) {
					continue
				}

				neighborHeight := heights.get(neighborPosition)

				if neighborHeight <= heightLimit {
					nextQueue[neighborPosition] = true
				}
			}
		}

		currentQueue = nextQueue
		currentDistance += 1
	}
}
