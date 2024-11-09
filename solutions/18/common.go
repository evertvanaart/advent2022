package day18

import (
	"advent2022/solutions/common"
	"math"
	"strings"
)

/* -------------------------------- Constants ------------------------------- */

var offsets = []position{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
	{-1, 0, 0},
	{0, -1, 0},
	{0, 0, -1},
}

/* ---------------------------- Utility Functions --------------------------- */

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/* -------------------------------- Position -------------------------------- */

type position [3]int

func addPositions(a position, b position) position {
	return position{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

/* --------------------------------- Bounds --------------------------------- */

type bounds struct {
	mins [3]int
	maxs [3]int
}

func createBounds() *bounds {
	mins := [3]int{math.MaxInt16, math.MaxInt16, math.MaxInt16}
	maxs := [3]int{math.MinInt16, math.MinInt16, math.MinInt16}
	return &bounds{mins, maxs}
}

func (b *bounds) update(pos position) {
	b.mins[0] = min(b.mins[0], pos[0])
	b.mins[1] = min(b.mins[1], pos[1])
	b.mins[2] = min(b.mins[2], pos[2])
	b.maxs[0] = max(b.maxs[0], pos[0])
	b.maxs[1] = max(b.maxs[1], pos[1])
	b.maxs[2] = max(b.maxs[2], pos[2])
}

func (b *bounds) widen() {
	b.mins[0] -= 1
	b.mins[1] -= 1
	b.mins[2] -= 1
	b.maxs[0] += 1
	b.maxs[1] += 1
	b.maxs[2] += 1
}

/* ---------------------------------- Grid ---------------------------------- */

type grid struct {
	values []int
	dims   [3]int
	offset [3]int
}

func createGrid(bounds *bounds) *grid {
	xWidth := bounds.maxs[0] - bounds.mins[0] + 1
	yWidth := bounds.maxs[1] - bounds.mins[1] + 1
	zWidth := bounds.maxs[2] - bounds.mins[2] + 1
	dims := [3]int{xWidth, yWidth, zWidth}
	gridSize := xWidth * yWidth * zWidth
	values := make([]int, gridSize)
	offset := bounds.mins

	return &grid{values, dims, offset}
}

func (g *grid) index(pos position) int {
	x := pos[0] - g.offset[0]
	y := pos[1] - g.offset[1]
	z := pos[2] - g.offset[2]

	return x + g.dims[0]*y + g.dims[0]*g.dims[1]*z
}

func (g *grid) contains(pos position) bool {
	x := pos[0] - g.offset[0]
	y := pos[1] - g.offset[1]
	z := pos[2] - g.offset[2]

	return x >= 0 && x < g.dims[0] &&
		y >= 0 && y < g.dims[1] &&
		z >= 0 && z < g.dims[2]
}

/* --------------------------------- Parsing -------------------------------- */

func parseLine(line string, bounds *bounds) position {
	fields := strings.Split(line, ",")
	values := make([]int, len(fields))

	for index, str := range fields {
		values[index] = common.ToInt(str)
	}

	pos := position{values[0], values[1], values[2]}
	bounds.update(pos)
	return pos
}
