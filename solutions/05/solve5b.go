package day05

import (
	"advent2022/solutions/common"
	"strings"
)

// For the B part, our custom stack implementation actually works in our favor; rather than having
// to pop N elements from the source stack and push them onto the target stack in reverse order,
// we can simply pop and push slices of length N of the underlying array, allowing us to easily
// move groups of containers from one stack to another while preserving their order.

func (stack *stack) pushN(values []byte) {
	for index, value := range values {
		stack.values[stack.top+index+1] = value
	}

	stack.top += len(values)
}

func (stack *stack) popN(count int) []byte {
	upper := stack.top + 1
	lower := upper - count
	values := stack.values[lower:upper]
	stack.top -= count
	return values
}

func applyB(stacks []*stack, instruction string) {
	fields := strings.Split(instruction, " ")
	count := common.ToInt(fields[1])
	source := common.ToInt(fields[3]) - 1
	target := common.ToInt(fields[5]) - 1

	values := stacks[source].popN(count)
	stacks[target].pushN(values)
}

func SolveB(lines []string) common.Solution {
	separatorIndex := findSeparatorLine(lines)
	instructions := lines[separatorIndex+1:]
	stackLines := lines[:separatorIndex]
	stacks := initializeStacks(stackLines)

	for _, instruction := range instructions {
		applyB(stacks, instruction)
	}

	tops := strings.Join(getTops(stacks), "")
	return common.ToStringSolution(tops)
}
