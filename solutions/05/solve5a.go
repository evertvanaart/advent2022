package day05

import (
	"advent2022/solutions/common"
	"strings"
)

// Since Go doesn't have a standard stack implementation, we'll just create our own. Counting
// the number of containers beforehand and allocating a stack array that can fit all containers
// means that we don't have to worry about resizing the array, although it's not actually that
// much faster than growing and shrinking the array on push and pop; most of the processing
// time is spent on stack initialization and string parsing.

func (stack *stack) push(value byte) {
	stack.top += 1
	stack.values[stack.top] = value
}

func (stack *stack) pop() byte {
	value := stack.values[stack.top]
	stack.top -= 1
	return value
}

func applyA(stacks []*stack, instruction string) {
	fields := strings.Split(instruction, " ")
	count := common.ToInt(fields[1])
	source := common.ToInt(fields[3]) - 1
	target := common.ToInt(fields[5]) - 1

	for i := 0; i < count; i++ {
		value := stacks[source].pop()
		stacks[target].push(value)
	}
}

func SolveA(lines []string) common.Solution {
	separatorIndex := findSeparatorLine(lines)
	instructions := lines[separatorIndex+1:]
	stackLines := lines[:separatorIndex]
	stacks := initializeStacks(stackLines)

	for _, instruction := range instructions {
		applyA(stacks, instruction)
	}

	tops := strings.Join(getTops(stacks), "")
	return common.ToStringSolution(tops)
}
