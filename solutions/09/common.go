package day09

type position [2]int

func otherDimension(dim int) int {
	if dim == 0 {
		return 1
	} else {
		return 0
	}
}

func sign(value int) int {
	if value > 0 {
		return 1
	} else {
		return -1
	}
}
