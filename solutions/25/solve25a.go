package day25

import (
	"advent2022/solutions/common"
)

// Always nice to end on something simple.

func adjust(a int) int {
	if a == 3 {
		return -2
	} else if a == 4 {
		return -1
	} else {
		return a
	}
}

func toDecimal(snafu string) int {
	digit := 1
	sum := 0

	for index := len(snafu) - 1; index >= 0; index-- {
		char := snafu[index]

		switch char {
		case '2':
			sum += 2 * digit
		case '1':
			sum += digit
		case '-':
			sum -= digit
		case '=':
			sum -= 2 * digit
		}

		digit *= 5
	}

	return sum
}

func toSnafu(value int) string {
	characters := []byte{}
	current := value
	digit := 1

	for current > 0 {
		modulo := current % (digit * 5)
		times := adjust(modulo / digit)

		switch times {
		case -2:
			characters = append(characters, '=')
		case -1:
			characters = append(characters, '-')
		case 0:
			characters = append(characters, '0')
		case 1:
			characters = append(characters, '1')
		case 2:
			characters = append(characters, '2')
		}

		current -= times * digit
		digit *= 5
	}

	reverse := make([]byte, len(characters))

	for index, char := range characters {
		reverse[len(characters)-index-1] = char
	}

	return string(reverse)
}

func SolveA(lines []string) common.Solution {
	sum := 0

	for _, line := range lines {
		value := toDecimal(line)
		sum += value
	}

	return common.ToStringSolution(toSnafu(sum))
}
