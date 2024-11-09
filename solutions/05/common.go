package day05

import "strings"

/* -------------------------------------------------------------------------- */
/*                                Stack Common                                */
/* -------------------------------------------------------------------------- */

type stack struct {
	values []byte
	top    int
}

func createStack(capacity int) *stack {
	values := make([]byte, capacity)
	return &stack{values, -1}
}

func (stack *stack) peek() byte {
	return stack.values[stack.top]
}

func (stack *stack) String() string {
	fields := make([]string, len(stack.values))

	for index, value := range stack.values {
		fields[index] = string(value)
	}

	return strings.Join(fields, "")
}

/* -------------------------------------------------------------------------- */
/*                             Stack Manipulation                             */
/* -------------------------------------------------------------------------- */

func getTops(stacks []*stack) []string {
	tops := []string{}

	for _, stack := range stacks {
		tops = append(tops, string(stack.peek()))
	}

	return tops
}

func initializeStack(lines []string, index int, capacity int) *stack {
	stack := createStack(capacity)

	// Move backwards through the stack lines (skipping the last one, which contains the numbers),
	// and add the character at the specified line index to the stack, unless it's a space.
	for i := len(lines) - 2; i >= 0; i-- {
		character := lines[i][index]

		if character == ' ' {
			break
		}

		stack.push(lines[i][index])
	}

	return stack
}

func initializeStacks(lines []string) []*stack {
	containers := countContainers(lines[:len(lines)-1])
	numbersLine := lines[len(lines)-1]
	stacks := []*stack{}

	// Assuming that stacks are always seperated by three spaces, and that both the container
	// contents and the stack number always consist of a single ASCII character.
	for index := 1; index < len(numbersLine); index += 4 {
		stack := initializeStack(lines, index, containers)
		stacks = append(stacks, stack)
	}

	return stacks
}

/* -------------------------------------------------------------------------- */
/*                                   Parsing                                  */
/* -------------------------------------------------------------------------- */

// Find the index of the empty line separating the initial stacks from the instructions.
func findSeparatorLine(lines []string) int {
	for index, line := range lines {
		if len(line) == 0 {
			return index
		}
	}

	panic("Failed to find separator line")
}

// Count the number of containers in the initial stack setup.
func countContainers(lines []string) int {
	count := 0

	for _, line := range lines {
		// Each container contains exactly one '[' character.
		count += strings.Count(line, "[")
	}

	return count
}
