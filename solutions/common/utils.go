package common

import (
	"fmt"
	"strconv"
)

/* -------------------------------------------------------------------------- */
/*                                    Types                                   */
/* -------------------------------------------------------------------------- */

// Number interface for generic functions
type Number interface {
	int
}

/* -------------------------------------------------------------------------- */
/*                                 Mathematics                                */
/* -------------------------------------------------------------------------- */

// Computes the sum of all numerical elements in a slice
func ComputeSum[V Number](values []V) V {
	sum := *new(V)

	for _, value := range values {
		sum += value
	}

	return sum
}

// Finds the index of the lowest numerical element in a slice
func FindMinIndex[V Number](values []V) int {
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

// Returns the larger of two numeric values
func Max[V Number](a V, b V) V {
	if a > b {
		return a
	} else {
		return b
	}
}

/* -------------------------------------------------------------------------- */
/*                                   Parsing                                  */
/* -------------------------------------------------------------------------- */

// Converts a string to an integer value, panics on failure
func ToInt(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("Cannot convert '%s' to integer", str))
	}

	return value
}
