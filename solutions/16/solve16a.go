package day16

import (
	"advent2022/solutions/common"
	"sort"
)

// First of all, note that there are several nodes in the input with zero flow and exactly two
// connections. We can simplify the resulting graph by eliminating these empty nodes; every node
// stores its non-empty neighbors (i.e., with non-zero flow or more than two neighbors), along
// with their distance, which is the number of intermediate empty nodes plus one. In the actual
// input, this simplification step reduces the number of nodes from 52 to just 16.
//
// Next, we greedily recurse through the graph. "Greedily" in this case means:
// - When we arrive at a valve that's not yet open, we first try the branch where we immediately
//   open this valve; the branch where we leave the valve closed and keep moving is expected to
//   lead to a lower final score in most cases.
// - When we need to select the next node to move to, we try the neighboring nodes in descending
//   order of scoring potential, based on their flow and then distance from the current node.
//
// Before moving away from a node, we check whether there's any point in continuing along the
// current branch, by checking if the maximum score that we can still get is more than the current
// best score. We do so by assuming that we can close all remaining valves in descending order of
// flow and with only one step between each valve, and computing the score we would get from such
// a perfect path. This is obviously an overestimation in almost all cases, but combined with the
// greedy approach to selecting the next action, it means that a lot of low-potential branches can
// be terminated early.
//
// The current method of potential score estimation does make the solution exact (i.e. we will
// always find the best possible score, not an approximation of this score), but it is somewhat
// slow; a smarter and/or faster estimation method might further improve performance. Still,
// the current solution clocks in at around 30ms, which is a time I can live with.

/* --------------------------------- Globals -------------------------------- */

const maxTimeA = 30

var bestScore int = -1

/* ---------------------------------- State --------------------------------- */

type stateA struct {
	score      int
	timeLeft   int
	closedFlow int
	position   int
	opened     []bool
}

func createStateA(maxTime int, startIndex int, nrValves int, totalFlow int) *stateA {
	return &stateA{0, maxTime, totalFlow, startIndex, make([]bool, nrValves)}
}

func (st *stateA) copy() *stateA {
	opened := make([]bool, len(st.opened))
	copy(opened, st.opened)

	return &stateA{st.score, st.timeLeft, st.closedFlow, st.position, opened}
}

/* -------------------------------- Recursion ------------------------------- */

func recurseA(valves valves, st *stateA) {
	// Check if we can stop iteration (i.e., out of time or all valves are open)
	if st.timeLeft <= 0 || st.closedFlow == 0 {
		if st.score > bestScore {
			bestScore = st.score
		}

		return
	}

	// First try opening the current valve if it's still closed
	if !st.opened[st.position] && valves[st.position].flow > 0 {
		newState := st.copy()
		newState.opened[st.position] = true
		newState.score += valves[st.position].flow * (st.timeLeft - 1)
		newState.closedFlow -= valves[st.position].flow
		newState.timeLeft -= 1

		recurseA(valves, newState)
	}

	// Estimate the maximum potential score, and stop iteration if it's lower than the best score
	potentialScore := valves.potentialScore(st.timeLeft, 2, st.opened)
	maxScore := st.score + potentialScore

	if potentialScore > 0 && maxScore < bestScore {
		return
	}

	valve := valves[st.position]
	candidates := make([]*candidate, len(valve.neighbors))

	// For each neighbor, compute the scoring potential based on flow and distance
	for index, neighbor := range valve.neighbors {
		neighborIndex := neighbor.index
		neighborFlow := valves[neighborIndex].flow
		isOpen := st.opened[neighborIndex]

		if neighborFlow == 0 || isOpen {
			candidates[index] = &candidate{neighborIndex, neighbor.distance, st.score}
		} else {
			newScore := st.score + neighborFlow*(st.timeLeft-neighbor.distance-1)
			candidates[index] = &candidate{neighborIndex, neighbor.distance, newScore}
		}
	}

	// Sort by scoring potential, and recurse to each
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	for _, candidate := range candidates {
		newState := st.copy()
		newState.timeLeft -= candidate.distance
		newState.position = candidate.index
		recurseA(valves, newState)
	}
}

func SolveA(lines []string) common.Solution {
	valves, startIndex := createValves(lines)

	totalFlow := valves.totalFlow()
	state := createStateA(maxTimeA, startIndex, len(valves), totalFlow)
	recurseA(valves, state)

	return common.ToIntSolution(bestScore)
}
