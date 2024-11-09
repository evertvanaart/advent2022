package day20

import "advent2022/solutions/common"

// Very similar; since we already used modulo operation to determine the target index, the jump in
// scale in and of itself changes very little. Having to repeat the main iteration ten times does
// of course increase the running time (~130ms for the B part), and the assumption we made for the
// findIndex() function - that values can probably be found close to where they started - does not
// hold at all after the first cycle. The complexity of findIndex() is exactly the same as that of
// a simple start-to-end linear search though, so it does not hurt performance either.

func SolveB(lines []string) common.Solution {
	entries := make([]entry, len(lines))
	zeroIndex := 0

	for index, line := range lines {
		entries[index] = entry{811589153 * common.ToInt(line), index}

		if entries[index].value == 0 {
			zeroIndex = index
		}
	}

	for counter := 0; counter < 10*len(entries); counter++ {
		index := counter % len(entries)
		sourceIndex := findIndex(entries, index)
		sourceValue := entries[sourceIndex].value
		targetIndex := sourceIndex + sourceValue
		targetIndex = wrapIndex(targetIndex, len(entries))
		shift(entries, sourceIndex, targetIndex)
	}

	finalZeroIndex := findIndex(entries, zeroIndex)
	value1 := entries[(finalZeroIndex+1000)%len(entries)].value
	value2 := entries[(finalZeroIndex+2000)%len(entries)].value
	value3 := entries[(finalZeroIndex+3000)%len(entries)].value
	return common.ToIntSolution(value1 + value2 + value3)
}
