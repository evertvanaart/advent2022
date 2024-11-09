package day16

import (
	"advent2022/solutions/common"
	"sort"
)

// The same basic idea, except we now track two positions instead of one. Since these two
// positions move at different paces (i.e., one may be ready to make a next move while the
// other is still in transit to the next node), we additionally track the number of minutes
// until either of them can make their next choice; if we ever find ourselves in a situation
// where both actors are not yet ready to make a choice, we simply count down the minutes
// until one of them is. When then decide which of the ready actors is allowed to act, and
// perform a recursive step very similar to the one in the A part; the only big difference
// is that we lower the score of a potential next valve if the other actor is already at or
// moving towards that valve, since there's usually not much point in doubling up.
//
// While this is fast - around 60ms on my reference system - it is unfortunately not exact;
// it produces the correct answer for the main input, but for the sample it is actually off
// by one. The culprit here is the potentialScore() function, in particular the timeStep
// parameter. In the A part we set this to 2, since even in the best possible case we'd
// only be able to open one valve every two minutes. This is of course no longer the case
// in the B part, where we can open one valve per minute on average. Leaving the step value
// at two therefore creates the risk of underestimation, which is indeed why it fails for the
// sample input; the path that would lead to the optimal score gets terminated early, because
// the underestimated best possible score is less than the current best score.
//
// Changing the time step to one fixes this underestimation in most cases - although it's
// still not 100% exact - but also dramatically increases the runtime: for the full input,
// it jumps from 60ms to around 2 minutes. I do have the outline of a smarter estimation
// algorithm in mind - which involves pre-computing the minimum distance between pairs of
// nodes, and using this data to more accurately estimate the maximum possible score - but
// this would take a while to implement, and frankly I've already spent way too much time
// on this day. This solution is either fast but inexact, or slow but exact, depending on
// whether you use a 2 or a 1 on line 115.

const maxTimeB = 26

/* -------------------------------- Utilities ------------------------------- */

func other(active int) int {
	if active == 0 {
		return 1
	} else {
		return 0
	}
}

/* ---------------------------------- State --------------------------------- */

type stateB struct {
	score      int
	timeLeft   int
	closedFlow int
	positions  [2]int
	nextChoice [2]int
	opened     []bool
}

func createStateB(maxTime int, startIndex int, nrValves int, totalFlow int) *stateB {
	positions := [2]int{startIndex, startIndex}
	nextChoices := [2]int{0, 0}
	return &stateB{0, maxTime, totalFlow, positions, nextChoices, make([]bool, nrValves)}
}

func (st *stateB) copy() *stateB {
	opened := make([]bool, len(st.opened))
	copy(opened, st.opened)

	positions := [2]int{st.positions[0], st.positions[1]}
	nextChoices := [2]int{st.nextChoice[0], st.nextChoice[1]}

	return &stateB{st.score, st.timeLeft, st.closedFlow, positions, nextChoices, opened}
}

func (st *stateB) determineActive() int {
	if st.nextChoice[0] == 0 {
		return 0
	} else {
		return 1
	}
}

/* -------------------------------- Recursion ------------------------------- */

func recurseB(valves valves, st *stateB) {
	for st.nextChoice[0] > 0 && st.nextChoice[1] > 0 {
		st.nextChoice[0] -= 1
		st.nextChoice[1] -= 1
		st.timeLeft -= 1
	}

	if st.timeLeft <= 0 || st.closedFlow == 0 {
		if st.score > bestScore {
			bestScore = st.score
		}

		return
	}

	active := st.determineActive()
	position := st.positions[active]
	valve := valves[position]

	if !st.opened[position] && valve.flow > 0 {
		newState := st.copy()
		newState.opened[position] = true
		newState.score += valve.flow * (st.timeLeft - 1)
		newState.closedFlow -= valve.flow
		newState.nextChoice[active] = 1

		recurseB(valves, newState)
	}

	potentialScore := valves.potentialScore(st.timeLeft, 2, st.opened)
	maxScore := st.score + potentialScore

	if potentialScore > 0 && maxScore < bestScore {
		return
	}

	candidates := make([]*candidate, len(valve.neighbors))

	for index, neighbor := range valve.neighbors {
		neighborIndex := neighbor.index
		neighborFlow := valves[neighborIndex].flow
		isOpen := st.opened[neighborIndex]
		otherPos := st.positions[other(active)]

		if neighborFlow == 0 || isOpen || neighborIndex == otherPos {
			candidates[index] = &candidate{neighborIndex, neighbor.distance, st.score}
		} else {
			newScore := st.score + neighborFlow*(st.timeLeft-neighbor.distance-1)
			candidates[index] = &candidate{neighborIndex, neighbor.distance, newScore}
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	for _, candidate := range candidates {
		newState := st.copy()
		newState.nextChoice[active] = candidate.distance
		newState.positions[active] = candidate.index
		recurseB(valves, newState)
	}
}

func SolveB(lines []string) common.Solution {
	valves, startIndex := createValves(lines)

	totalFlow := valves.totalFlow()
	state := createStateB(maxTimeB, startIndex, len(valves), totalFlow)
	recurseB(valves, state)

	return common.ToIntSolution(bestScore)
}
