package day19

import (
	"fmt"
	"strconv"
	"strings"
)

/* --------------------------------- Globals -------------------------------- */

var maxTime = 24

const (
	ore      = 0
	clay     = 1
	obsidian = 2
	geode    = 3
)

/* --------------------------------- Parsing -------------------------------- */

func getMaxCosts(costsPerRobot [4]costs) costs {
	max := costs{0, 0, 0}

	for robot := ore; robot <= geode; robot++ {
		for resource := ore; resource <= obsidian; resource++ {
			if costsPerRobot[robot][resource] > max[resource] {
				max[resource] = costsPerRobot[robot][resource]
			}
		}
	}

	return max
}

func parseNumbers(str string) []int {
	fields := strings.Split(str, " ")
	numbers := []int{}

	for _, field := range fields {
		number, err := strconv.Atoi(field)

		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func parse(line string) *blueprint {
	numbers := parseNumbers(line)

	oreCosts := costs{numbers[0], 0, 0}
	clayCosts := costs{numbers[1], 0, 0}
	obsidianCosts := costs{numbers[2], numbers[3], 0}
	geodeCosts := costs{numbers[4], 0, numbers[5]}

	costsPerRobot := [4]costs{oreCosts, clayCosts, obsidianCosts, geodeCosts}
	maxCosts := getMaxCosts(costsPerRobot)

	return &blueprint{costsPerRobot, maxCosts}
}

/* -------------------------------- Blueprint ------------------------------- */

type costs [3]int
type blueprint struct {
	costsPerRobot [4]costs
	maxCosts      costs
}

/* -------------------------------- Resources ------------------------------- */

type resources [4]int

func (rs *resources) gain(income *resources) {
	rs[ore] += income[ore]
	rs[clay] += income[clay]
	rs[obsidian] += income[obsidian]
	rs[geode] += income[geode]
}

func (rs *resources) pay(costs costs) {
	rs[ore] -= costs[ore]
	rs[clay] -= costs[clay]
	rs[obsidian] -= costs[obsidian]
}

func (rs *resources) canAfford(costs costs) bool {
	return rs[ore] >= costs[ore] &&
		rs[clay] >= costs[clay] &&
		rs[obsidian] >= costs[obsidian]
}

func (rs *resources) isPositive(costs costs) bool {
	return (costs[ore] == 0 || rs[ore] > 0) &&
		(costs[clay] == 0 || rs[clay] > 0) &&
		(costs[obsidian] == 0 || rs[obsidian] > 0)
}

func (rs *resources) copy() *resources {
	return &resources{rs[0], rs[1], rs[2], rs[3]}
}

/* ---------------------------------- State --------------------------------- */

type state struct {
	time      int
	buildNext int
	current   *resources
	income    *resources
}

func initialState(buildFirst int) *state {
	current := &resources{0, 0, 0, 0}
	income := &resources{1, 0, 0, 0}
	return &state{0, buildFirst, current, income}
}

func (st *state) String() string {
	return fmt.Sprintf("{time: %d, buildNext: %d, current: %v, income: %v}",
		st.time, st.buildNext, st.current, st.income)
}

func (st *state) copy(buildNext int) *state {
	return &state{st.time, buildNext, st.current.copy(), st.income.copy()}
}

func estimate(currentGeodes int, geodeIncome int, timeLeft int) int {
	current := currentGeodes
	income := geodeIncome

	for time := 0; time < timeLeft; time++ {
		current += income
		income += 1
	}

	return current
}

/* -------------------------------- Recursion ------------------------------- */

func recurse(st *state, bp *blueprint, bestScore int) int {
	nextCosts := bp.costsPerRobot[st.buildNext]

	// Wait until we can build the next robot, or time runs out
	for !st.current.canAfford(nextCosts) && st.time < maxTime {
		st.current.gain(st.income)
		st.time += 1
	}

	if st.time >= maxTime {
		return st.current[geode]
	}

	// Check if there's still a way to beat the current best score
	bestPossibleScore := estimate(st.current[geode], st.income[geode], maxTime-st.time)

	if bestPossibleScore < bestScore {
		return bestScore
	}

	// Build the next robot
	st.current.gain(st.income)
	st.current.pay(bp.costsPerRobot[st.buildNext])
	st.income[st.buildNext] += 1
	st.time += 1

	newBestScore := bestScore

	// Try building robots in reverse order
	for resource := geode; resource >= ore; resource-- {
		if st.income.isPositive(bp.costsPerRobot[resource]) {
			// Stop building robots if we've reached the maximum resource cost
			if resource != geode && st.income[resource] >= bp.maxCosts[resource] {
				continue
			}

			newState := st.copy(resource)
			score := recurse(newState, bp, newBestScore)

			if score > bestScore {
				newBestScore = score
			}
		}
	}

	return newBestScore
}

func solveBlueprint(bp *blueprint) int {
	best1 := recurse(initialState(clay), bp, 0)
	best2 := recurse(initialState(ore), bp, 0)

	if best1 > best2 {
		return best1
	} else {
		return best2
	}
}
