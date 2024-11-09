package day24

/* --------------------------------- Globals -------------------------------- */

var source = vector2d{0, 0}
var target = vector2d{0, 0}

var offset0 = vector2d{0, 0}
var offsetD = vector2d{1, 0}
var offsetR = vector2d{0, 1}
var offsetU = vector2d{-1, 0}
var offsetL = vector2d{0, -1}

var initOffsets = []vector2d{offsetD, offset0}
var allOffsets = []vector2d{offsetR, offsetD, offset0, offsetL, offsetU}

/* -------------------------------- Utilities ------------------------------- */

func limit(value int, max int) int {
	modulo := value % max

	if modulo < 0 {
		return modulo + max
	} else {
		return modulo
	}
}

func getOffsets(pos vector2d) []vector2d {
	if pos == source {
		return initOffsets
	} else {
		return allOffsets
	}
}

/* --------------------------------- Vector --------------------------------- */

type vector2d [2]int

func (pos vector2d) outside(size vector2d) bool {
	return pos[0] < 0 || pos[0] >= size[0] || pos[1] < 0 || pos[1] >= size[1]
}

/* ---------------------------------- Grid ---------------------------------- */

type grid struct {
	values    []bool
	size      vector2d
	direction vector2d
}

func (g *grid) get(pos vector2d, time int) bool {
	row := limit(pos[0]-time*g.direction[0], g.size[0])
	col := limit(pos[1]-time*g.direction[1], g.size[1])
	index := row*g.size[1] + col
	return g.values[index]
}

func parseGrid(lines []string, char byte, direction vector2d) *grid {
	rows := len(lines)
	cols := len(lines[0]) - 2
	values := make([]bool, rows*cols)

	for row, line := range lines {
		for col := 0; col < cols; col++ {
			if line[col+1] == char {
				index := row*cols + col
				values[index] = true
			}
		}
	}

	return &grid{values, vector2d{rows, cols}, direction}
}

/* ------------------------------- Core logic ------------------------------- */

func collidesWithBlizzards(blizzards []*grid, pos vector2d, time int) bool {
	for _, blizzard := range blizzards {
		if blizzard.get(pos, time) {
			return true
		}
	}

	return false
}

func validPosition(time int, pos vector2d, blizzards []*grid, size vector2d) bool {
	if pos == source || pos == target {
		return true
	}

	return !(pos.outside(size) || collidesWithBlizzards(blizzards, pos, time))
}

func step(time int, pos vector2d, blizzards []*grid, size vector2d) []vector2d {
	if !validPosition(time, pos, blizzards, size) {
		return []vector2d{}
	}

	offsets := getOffsets(pos)
	outPositions := []vector2d{}

	for _, offset := range offsets {
		nextPos := vector2d{pos[0] + offset[0], pos[1] + offset[1]}
		outPositions = append(outPositions, nextPos)
	}

	return outPositions
}
