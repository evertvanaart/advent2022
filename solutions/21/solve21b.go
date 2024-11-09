package day21

import (
	"advent2022/solutions/common"
)

// The initial part is mostly the same as the A part, expect that we 1) change the operand of the
// root monkey to "="; 2) tag the human with a special state and don't add it to the starting list,
// and 3) don't return when we've found the root monkey. This produces a partially solved monkey
// map, i.e. all monkeys that do not depend on the human (directly or indirectly) will be done.
//
// From the input and the dependency map, we can observe that there is a single path leading from
// the root to the human, i.e. only one monkey M1 uses the human, then only one monkey M2 uses M1,
// and so on, all the way to the top. This observation makes things a lot easier, since we do not
// have to consider cases where both sides of a monkey's equation depend on the human.
//
// Using this observation, we start from the root monkey in our partially solved monkey map,
// and move down the tree graph; at each step, we solve the inverse of the equation to obtain the
// value of the unknown side, using the result and the value of the known side. We then use this
// value as the result in the equation of the monkey representing the unknown side, and keep
// repeating this until we've arrived at the human.

func (m *monkey) reverseRight(lh int, result int) int {
	switch m.operand {
	case "+":
		return result - lh
	case "-":
		return lh - result
	case "*":
		return result / lh
	case "/":
		return lh / result
	case "=":
		return lh
	}

	return -1
}

func (m *monkey) reverseLeft(rh int, result int) int {
	switch m.operand {
	case "+":
		return result - rh
	case "-":
		return result + rh
	case "*":
		return result / rh
	case "/":
		return result * rh
	case "=":
		return rh
	}

	return -1
}

func (m *monkey) reverse(result int, monkeyMap map[string]*monkey) (int, string) {
	lhMonkey := monkeyMap[m.vars[0]]
	rhMonkey := monkeyMap[m.vars[1]]

	if lhMonkey.state == done && rhMonkey.state != done {
		return m.reverseRight(lhMonkey.number, result), rhMonkey.name
	} else if lhMonkey.state != done && rhMonkey.state == done {
		return m.reverseLeft(rhMonkey.number, result), lhMonkey.name
	}

	panic("Expected exactly one monkey to be ready")
}

func SolveB(lines []string) common.Solution {
	monkeyMap := map[string]*monkey{}
	dependencyMap := map[string][]string{}
	startingMonkeys := []string{}

	for _, line := range lines {
		monkey := parseMonkey(line)
		monkeyMap[monkey.name] = monkey

		if monkey.name == "root" {
			monkey.operand = "="
		} else if monkey.name == "humn" {
			monkey.state = human
		}

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

		ready := monkey.onReady(monkeyMap, dependencyMap)
		readyMonkeys = append(readyMonkeys, ready...)
	}

	currentMonkey := "root"
	currentResult := 0

	for currentMonkey != "humn" {
		monkey := monkeyMap[currentMonkey]
		currentResult, currentMonkey = monkey.reverse(currentResult, monkeyMap)
	}

	return common.ToIntSolution(currentResult)
}
