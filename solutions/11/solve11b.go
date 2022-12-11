package day11

import "advent2022/solutions/common"

// With 10000 rounds and unbounded worry growth, we will eventually start overflowing the integers
// used to track the worry score. We could of course just throw more bits at the problem, but this
// won't scale as the number of rounds continue to go up. Based on nothing but intuition, I figured
// that we could instead limit the worry score to the range dictated by the common divisor, i.e. if
// (value % divisor) is equal to zero, then ((value % common_divisor) % divisor) must also be equal
// to zero. The common divisor is computed simply by multiplying all the test divisors used by the
// monkeys; I did note that these divisors are all prime numbers, but I couldn't figure out how to
// use this to my advantage. It's not perfect - in particular, I think a badly-timed square
// operation could still cause an overflow - but it does produce the correct answer.
//
// As a side note, this is the first problem this year which took more than a millisecond to run;
// none of the turn processing code was designed with speed in mind, and we now have to execute
// 80,000 of those turns in total. Maybe there's a smart way to model the path of an item in
// such a way that we don't have to execute every step along the way? Maybe items will always
// settle into a specific loop between monkeys after some time, and we can somehow compute the
// periods of those loops? Who knows! (not me)

func SolveB(lines []string) common.Solution {
	monkeys := []*monkey{}

	for lineNo := 0; lineNo < len(lines); lineNo += linesPerMonkey {
		monkeyLines := lines[lineNo : lineNo+linesPerMonkey]
		monkeys = append(monkeys, parseMonkey(monkeyLines))
	}

	commonDivisor := 1

	for _, monkey := range monkeys {
		commonDivisor *= monkey.divisor
	}

	adjust := func(value int) int {
		return value % commonDivisor
	}

	for round := 0; round < 10000; round++ {
		processRound(monkeys, adjust)
	}

	output := determineOutput(monkeys)
	return common.ToInt64Solution(output)
}
