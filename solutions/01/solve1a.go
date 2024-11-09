package day01

import "advent2022/solutions/common"

// Basic solution: Iterate through the input lines, convert each non-empty line to an integer,
// keep track of the sum for the current block of lines; when encountering the end of a block
// (either because of an empty line, or because we've reached the end), compare the current
// sum value to the maximum; finally return the highest block sum.

// Returns the larger of two numeric values
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func SolveA(lines []string) common.Solution {
	currentSum := 0
	maxSum := 0

	for _, line := range lines {
		if len(line) == 0 {
			maxSum = max(maxSum, currentSum)
			currentSum = 0
			continue
		}

		lineVal := common.ToInt(line)
		currentSum += lineVal
	}

	return common.ToIntSolution(max(maxSum, currentSum))
}
