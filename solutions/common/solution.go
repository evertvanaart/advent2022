package common

import "fmt"

/* -------------------------------- Interface ------------------------------- */

type Solution interface {
	String() string
}

/* ------------------------------- IntSolution ------------------------------ */

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

/* ------------------------------ Int64Solution ----------------------------- */

type Int64Solution struct {
	Solution
	value int64
}

func (solution Int64Solution) String() string {
	return fmt.Sprintf("%v", solution.value)
}

func ToInt64Solution(value int64) Solution {
	return Int64Solution{value: value}
}

/* ----------------------------- StringSolution ----------------------------- */

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
