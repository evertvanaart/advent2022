package common

import (
	"fmt"
	"strconv"
)

// Converts a string to an integer value, panics on failure
func ToInt(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("Cannot convert '%s' to integer", str))
	}

	return value
}

func ToIntN(strings []string) []int {
	values := make([]int, len(strings))

	for index, str := range strings {
		values[index] = ToInt(str)
	}

	return values
}
