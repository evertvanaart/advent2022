package day01

import "advent2022/solutions/common"

// The obvious approach would be to insert the sum of each block of lines into an array, then sort
// it and take the three highest values. The solution below instead keeps track of the three highest
// values encountered so far (in no particular order), as well as the index of the lowest of these
// three values (i.e., the third-highest value encountered so far). If a newly computed sum value
// is higher than this third-place value, it replaces the old third-place value, and we again check
// which of the top-three values is lowest. This gives us a runtime complexity of O(N) (where N is
// the number of elves), as opposed to O(NlogN) if we used sorting.

func computeSum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func findMinIndex(values []int) int {
	if len(values) == 0 {
		return -1
	}

	minValue := values[0]
	minIndex := 0

	for index, value := range values {
		if value < minValue {
			minIndex = index
			minValue = value
		}
	}

	return minIndex
}

// Stores the X highest values encountered so far; in this case, X is always equal to three.
// Keeping track of the index of the lowest of these three values makes the update() function
// slightly more efficient (i.e., it saves us from having to call findMinIndex() every time).
type topValues struct {
	values   []int
	minIndex int
}

// Initializes the object to track the highest values
func createTopValues(size int) *topValues {
	return &topValues{values: make([]int, size), minIndex: 0}
}

// Check if a newly computed sum value should be in the top 3 by comparing it against the lowest
// top 3 value. If the new value is higher, add it to the top 3 (overwriting the old lowest value),
// and again find the index of the lowest value.
func (topValues *topValues) update(sum int) {
	minValue := topValues.values[topValues.minIndex]

	if sum > minValue {
		topValues.values[topValues.minIndex] = sum
	}

	topValues.minIndex = findMinIndex(topValues.values)
}

func SolveB(lines []string) common.Solution {
	topValues := createTopValues(3)
	currentSum := 0

	for _, line := range lines {
		if len(line) == 0 {
			topValues.update(currentSum)
			currentSum = 0
			continue
		}

		lineVal := common.ToInt(line)
		currentSum += lineVal
	}

	topValues.update(currentSum)
	return common.ToIntSolution(computeSum(topValues.values))
}
