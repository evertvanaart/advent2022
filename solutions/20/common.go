package day20

type entry struct {
	value int
	index int
}

func wrapIndex(index int, length int) int {
	mod := length - 1

	if index <= 0 {
		return index%mod + mod
	} else if index >= mod {
		return index % mod
	} else {
		return index
	}
}

func limitIndex(index int, length int) int {
	if index < 0 {
		return index%length + length
	} else {
		return index % length
	}
}

func findIndex(entries []entry, index int) int {
	halfLength := len(entries) / 2

	for distance := 0; distance <= halfLength; distance++ {
		lowerIndex := limitIndex(index-distance, len(entries))

		if entries[lowerIndex].index == index {
			return lowerIndex
		}

		upperIndex := limitIndex(index+distance+1, len(entries))

		if entries[upperIndex].index == index {
			return upperIndex
		}
	}

	return -1
}

func shift(entries []entry, source int, target int) {
	sourceEntry := entries[source]

	if target < source {
		sourceSlice := entries[target:source]
		targetSlice := entries[target+1 : source+1]
		copy(targetSlice, sourceSlice)
	} else if target > source {
		sourceSlice := entries[source+1 : target+1]
		targetSlice := entries[source:target]
		copy(targetSlice, sourceSlice)
	}

	entries[target] = sourceEntry
}
