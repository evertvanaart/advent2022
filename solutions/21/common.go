package day21

import (
	"fmt"
	"strconv"
	"strings"
)

/* -------------------------------- Constants ------------------------------- */

const (
	done      = -1
	noneReady = 0
	oneReady  = 1
	bothReady = 2
	human     = 3
)

/* --------------------------------- Monkeys -------------------------------- */

type monkey struct {
	name    string
	number  int
	state   int
	vars    []string
	operand string
}

func (m *monkey) String() string {
	if m.state == done && m.vars == nil {
		return fmt.Sprintf("%s: %v", m.name, m.number)
	} else if m.state == done {
		return fmt.Sprintf("%s: %s %s %s = %d",
			m.name, m.vars[0], m.operand, m.vars[1], m.number)
	} else {
		return fmt.Sprintf("%s: %s %s %s = ?",
			m.name, m.vars[0], m.operand, m.vars[1])
	}
}

func (m *monkey) solve(lh int, rh int) {
	result := 0

	switch m.operand {
	case "+":
		result = lh + rh
	case "-":
		result = lh - rh
	case "*":
		result = lh * rh
	case "/":
		result = lh / rh
	}

	m.number = result
	m.state = done
}

func (m *monkey) onReady(monkeyMap map[string]*monkey,
	dependencyMap map[string][]string) []string {
	readyMonkeys := []string{}

	for _, dependency := range dependencyMap[m.name] {
		monkeyMap[dependency].state += 1

		if monkeyMap[dependency].state == bothReady {
			readyMonkeys = append(readyMonkeys, dependency)
		}
	}

	return readyMonkeys
}

func parseMonkey(line string) *monkey {
	fields := strings.Split(line, ": ")
	value, err := strconv.Atoi(fields[1])

	if err == nil {
		return &monkey{fields[0], value, -1, nil, ""}
	}

	subfields := strings.Split(fields[1], " ")
	vars := []string{subfields[0], subfields[2]}
	return &monkey{fields[0], -1, 0, vars, subfields[1]}
}

/* ------------------------------ Dependencies ------------------------------ */

func addDepedencies(dependencyMap map[string][]string, name string, vars []string) {
	for _, varName := range vars {
		current, exists := dependencyMap[varName]

		if exists {
			dependencyMap[varName] = append(current, name)
		} else {
			dependencyMap[varName] = []string{name}
		}
	}
}
