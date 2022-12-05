package day05

import (
	"fmt"
	"strings"
)

type stack struct {
	values []byte
}

func (stack *stack) add(value byte) {
	fmt.Printf("Adding '%s'\n", string(value))
	stack.values = append(stack.values, value)
}

func (stack *stack) pop() byte {
	if len(stack.values) == 0 {
		panic("Unexpected empty stack")
	}

	value := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]
	return value
}

func (stack *stack) String() string {
	fields := make([]string, len(stack.values))

	for _, value := range stack.values {
		field := fmt.Sprintf("[%s]", string(value))
		fields = append(fields, field)
	}

	return fmt.Sprintf("[%s]", strings.Join(fields, ", "))
}

func findSeparatorLine(lines []string) int {
	for index, line := range lines {
		if len(line) == 0 {
			return index
		}
	}

	panic("Failed to find separator line")
}

// todo: makes assumptions about input format
func findStackIndices(line string) []int {
	fields := strings.Split(line, " ")
	indices := []int{}

	for index, field := range fields {
		if len(field) > 0 {
			indices = append(indices, index)
		}
	}

	return indices
}

func initializeStackLine(stacks []*stack, indices []int, line string) {
	for stackIndex, lineIndex := range indices {
		if line[lineIndex] != ' ' {
			stacks[stackIndex].add(line[lineIndex])
		}
	}
}

func initializeStacks(stacks []*stack, indices []int, lines []string) {
	for index := range stacks {
		stacks[index] = &stack{[]byte{}}
	}

	for i := len(lines) - 1; i >= 0; i-- {
		initializeStackLine(stacks, indices, lines[i])
	}
}

func SolveA(lines []string) int {
	separatorIndex := findSeparatorLine(lines)
	// instructions := lines[separatorIndex+1:]
	stackLines := lines[:separatorIndex-1]
	stacksNos := lines[separatorIndex-1]

	stackIndices := findStackIndices(stacksNos)
	fmt.Println(stackIndices)
	stacks := make([]*stack, len(stackIndices))

	initializeStacks(stacks, stackIndices, stackLines)
	fmt.Println(stacks)

	return 0
}
