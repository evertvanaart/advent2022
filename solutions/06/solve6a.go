package day06

import "advent2022/solutions/common"

// For additional realism and/or challenge, I'm pretending that the characters really do arrive
// one at a time, i.e. the solution logic doesn't have access to the full input string, and the
// buffer must use constant memory. This solution uses a naive FIFO buffer that stores the last
// four characters received, and checks if they are all unique. There is some room for improvement
// here - shifting the buffer one-by-one is not very efficient, and the double for-loop in the
// allUnique() function obviously isn't great - but I'll address those concerns in the B part;
// for a buffer of just four characters, this naive approach is good enough.

type bufferA struct {
	values   []rune
	capacity int
	received int
}

func createBufferA(capacity int) *bufferA {
	return &bufferA{values: make([]rune, capacity), capacity: capacity, received: 0}
}

func (buffer *bufferA) observe(character rune) {
	for i := 0; i < buffer.capacity-1; i++ {
		buffer.values[i] = buffer.values[i+1]
	}

	buffer.values[buffer.capacity-1] = character
	buffer.received += 1
}

func (buffer *bufferA) allUnique() bool {
	if buffer.received < buffer.capacity {
		return false
	}

	for i := 0; i < buffer.capacity-1; i++ {
		for j := i + 1; j < buffer.capacity; j++ {
			if buffer.values[i] == buffer.values[j] {
				return false
			}
		}
	}

	return true
}

func SolveA(lines []string) common.Solution {
	buffer := createBufferA(4)

	for index, character := range lines[0] {
		buffer.observe(character)

		if buffer.allUnique() {
			return common.ToIntSolution(index + 1)
		}
	}

	panic("No start-of-packet marker found!")
}
