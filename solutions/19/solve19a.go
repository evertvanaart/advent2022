package day19

import (
	"advent2022/solutions/common"
)

// My first attempt used a naive recursion approach, where I'd decide what to do next at every
// time step (i.e., either build one of the available robots, or do nothing to wait for more
// resources). This obviously lead to an explosion in possible states after ten steps or so,
// and so in order to keep runtime down I tacked on several optimizations, including a kind
// of memoization to decide whether the current state (based on current resources and income)
// was better than others observed so far. The resulting solution was pretty fast, and produced
// the correct result for the sample input, but the results for the actual input was wrong. With
// no way to debug the already overcomplicated logic, I ended up binning this approach.
//
// This second approach is much simpler; instead of deciding what to do at every point in time,
// we decide which robot we are going to build next (limited to those options for which we have
// positive income for all required resources), and then wait until we've gathered the resources
// needed to build this robot. As an additional rule, we specify that we stop building robots of
// a certain type if the income of the corresponding resource is already equal to the maximum
// cost (i.e., if no robot costs more than 3 ore, there is no point in building a fourth ore
// robot, since we can at most spend 3 ore per minute).
//
// With just this approach, the solution is already pretty fast, around 450ms. To optimize this
// further, we can use an approach similar to the one used for Day 16, where we first take the
// greedy route (which in this case means first trying to build geode robots, then obsidian,
// then clay, and then ore); once we have a best score, we can terminate other branches early
// if the estimated best possible score is less than the current actual best score. We use a
// gross overestimation - essentially assuming that we can keep building an additional geode
// robot every single minute - but even then, this check drops runtime from around 450ms to
// just under 200ms. As with Day 16, improving this estimation method (without allowing it
// to underestimate the score) could improve runtime further, but between the aborted first
// attempt and this one, I think I've spent enough time on this problem.

func SolveA(lines []string) common.Solution {
	sum := 0

	for index, line := range lines {
		blueprint := parse(line)
		score := solveBlueprint(blueprint)
		sum += (index + 1) * score
	}

	return common.ToIntSolution(sum)
}
