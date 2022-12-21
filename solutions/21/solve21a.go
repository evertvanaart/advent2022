package day21

import (
	"advent2022/solutions/common"
)

// Another fairly straightforward solution; first parse the monkeys according to the syntax, and
// store them all in a map; additionally, create a dependency map, with monkey name M as the key,
// and a list of monkeys depending on M (i.e., using M as one of their values) as the value. Next,
// determine the list of starting monkeys, i.e. monkeys that have a number instead of an equation,
// and check which monkeys depend on two starting monkeys. We then iterate through this list of
// ready monkeys; for each, we solve its equation, and using the dependency map, let all its
// dependencies know that one of their values is done; if this causes a dependency to become
// ready, we add it to the list. We can stop this iteration as soon as we find the root monkey.

func SolveA(lines []string) common.Solution {
	monkeyMap := map[string]*monkey{}
	dependencyMap := map[string][]string{}
	startingMonkeys := []string{}

	for _, line := range lines {
		monkey := parseMonkey(line)
		monkeyMap[monkey.name] = monkey

		if monkey.state == done {
			startingMonkeys = append(startingMonkeys, monkey.name)
		} else {
			addDepedencies(dependencyMap, monkey.name, monkey.vars)
		}
	}

	readyMonkeys := []string{}

	for _, startingMonkey := range startingMonkeys {
		monkey := monkeyMap[startingMonkey]
		ready := monkey.onReady(monkeyMap, dependencyMap)
		readyMonkeys = append(readyMonkeys, ready...)
	}

	for index := 0; index < len(readyMonkeys); index++ {
		monkey := monkeyMap[readyMonkeys[index]]
		value1 := monkeyMap[monkey.vars[0]].number
		value2 := monkeyMap[monkey.vars[1]].number

		monkey.solve(value1, value2)

		if monkey.name == "root" {
			return common.ToIntSolution(monkey.number)
		}

		ready := monkey.onReady(monkeyMap, dependencyMap)
		readyMonkeys = append(readyMonkeys, ready...)
	}

	return common.ToIntSolution(-1)
}
