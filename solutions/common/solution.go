package common

import "fmt"

type Solution interface {
	String() string
}

type IntSolution struct {
	Solution
	value int
}

func (solution IntSolution) String() string {
	return fmt.Sprintf("%v", solution.value)
}

func ToIntSolution(value int) Solution {
	return IntSolution{value: value}
}

type StringSolution struct {
	Solution
	value string
}

func (solution StringSolution) String() string {
	return solution.value
}

func ToStringSolution(value string) Solution {
	return StringSolution{value: value}
}
