package day20

import "advent2022/solutions/common"

// Somehow, the hardest thing to deal with in this challenge was the way that numbers wrap around
// to the other side, which I still don't fully understand; somehow these number always seem to
// move one space more than they should, which makes for some annoying off-by-one errors.
//
// One thing to note about the input is that some of the numbers are larger than the size of the
// array. To avoid having to shift the same array multiple times, we can use a modulo operation
// (keeping in mind the weird wrapping behavior) to quickly determine the final target index; in
// the sample of length 7, a value of -5 or -11 or -17 would end up at exactly the same index as
// a value of 1. Once we determine the target index, we can use a single copy() call to shift the
// slice between the source and target indices by one, and then add the shifting value to the new
// open spot. The copy() call is O(N) though, so this solution still involves a lot of one-by-one
// shifting operations, which makes it somewhat slow (~15ms for the A part).
//
// One final thing to note is how we find the next value to shift, since it will probably have
// moved away from where it started. To this end, we store the original index alongside the
// value. When we want to find the current index of a value that started at index i, we start
// looking at i and move our search left and right one step at a time, on the assumption that
// values that haven't moved yet probably don't move too far away from their starting position.

func SolveA(lines []string) common.Solution {
	entries := make([]entry, len(lines))
	zeroIndex := 0

	for index, line := range lines {
		entries[index] = entry{common.ToInt(line), index}
	}

	for index := range lines {
		sourceIndex := findIndex(entries, index)
		sourceValue := entries[sourceIndex].value
		targetIndex := sourceIndex + sourceValue
		targetIndex = wrapIndex(targetIndex, len(entries))
		shift(entries, sourceIndex, targetIndex)

		if sourceValue == 0 {
			zeroIndex = index
		}
	}

	finalZeroIndex := findIndex(entries, zeroIndex)
	value1 := entries[(finalZeroIndex+1000)%len(entries)].value
	value2 := entries[(finalZeroIndex+2000)%len(entries)].value
	value3 := entries[(finalZeroIndex+3000)%len(entries)].value
	return common.ToIntSolution(value1 + value2 + value3)
}
