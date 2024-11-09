package day06

import "advent2022/solutions/common"

// As mentioned, the naive buffer approach used in the A part doesn't scale well. To safeguard
// performance now that we've got 14 characters instead of 4, the following changes were made:
//
// 1) Avoid shifting every buffer value one-by-one on every call to observe() by using a ring
//    buffer; write new elements to a position indicated by an index, which is incremented
//    after every write; reset the index to zero when we reach the end of the buffer. Also
//    track whether the buffer is 'full', i.e. whether we've received at least 14 values.
//
// 2) Keep track of the number of duplicates in the current buffer, so that we can check for
//    the marker in constant time. We use the knowledge that the input is restricted to the
//    lower-case alphabet to efficiently track the number of occurrences of a specific letter
//    using a fixed-size array, rather than a map. Whenever any letter count becomes larger
//    than one, we increase the duplicate count, and whenever we remove a letter (when it's
//    dropped from the buffer) we decrease the duplicate count if the letter count is back
//    down to one. To find the marker, we only need to check if the duplicate count is
//    zero (and the buffer is full), which is a single O(1) comparison.

type counter struct {
	counts     []int
	duplicates int
}

func createCounter() *counter {
	return &counter{counts: make([]int, 26), duplicates: 0}
}

func (counter *counter) add(character rune) {
	index := character - 'a'
	counter.counts[index] += 1

	if counter.counts[index] == 2 {
		counter.duplicates += 1
	}
}

func (counter *counter) remove(character rune) {
	index := character - 'a'
	counter.counts[index] -= 1

	if counter.counts[index] == 1 {
		counter.duplicates -= 1
	}
}

type bufferB struct {
	values   []rune
	counter  *counter
	capacity int
	index    int
	full     bool
}

func createBufferB(capacity int) *bufferB {
	return &bufferB{
		values:   make([]rune, capacity),
		counter:  createCounter(),
		capacity: capacity,
		index:    0,
		full:     false,
	}
}

func (buffer *bufferB) observe(character rune) {
	if buffer.full {
		toRemove := buffer.values[buffer.index]
		buffer.counter.remove(toRemove)
	}

	buffer.counter.add(character)
	buffer.values[buffer.index] = character
	buffer.index += 1

	if buffer.index == buffer.capacity {
		buffer.full = true
		buffer.index = 0
	}
}

func (buffer *bufferB) allUnique() bool {
	return buffer.full && (buffer.counter.duplicates == 0)
}

func SolveB(lines []string) common.Solution {
	buffer := createBufferB(14)

	for index, character := range lines[0] {
		buffer.observe(character)

		if buffer.allUnique() {
			return common.ToIntSolution(index + 1)
		}
	}

	panic("No start-of-message marker found!")
}
