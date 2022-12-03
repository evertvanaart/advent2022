package day03

func getPriority(c rune) int {
	if c <= 'Z' {
		return int(c - 'A' + 27)
	} else {
		return int(c - 'a' + 1)
	}
}

func createCharacterArray(str string) []bool {
	array := make([]bool, 128)

	for _, c := range str {
		array[c] = true
	}

	return array
}
