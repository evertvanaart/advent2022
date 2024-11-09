package day17

/* -------------------------------- Constants ------------------------------- */

const (
	maxRockNoA    = 2022
	maxRockNoB    = 1000000000000
	avgRockHeight = 3
	gridWidth     = 7
)

/* ---------------------------------- Rocks --------------------------------- */

type position [2]int

type rock struct {
	offsets []position
	width   int
	height  int
}

var rocks = []rock{
	{[]position{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, 4, 1},
	{[]position{{0, 1}, {1, 1}, {2, 1}, {1, 0}, {1, 2}}, 3, 3},
	{[]position{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, 3, 3},
	{[]position{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, 1, 4},
	{[]position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, 2, 2},
}

/* ---------------------------------- Gusts --------------------------------- */

type gusts struct {
	values []int
	index  int
}

func parseGusts(line string) *gusts {
	values := make([]int, len(line))

	for index, char := range line {
		if char == '<' {
			values[index] = -1
		} else {
			values[index] = 1
		}
	}

	return &gusts{values, 0}
}

func (g *gusts) step() int {
	value := g.values[g.index]
	g.index += 1

	if g.index == len(g.values) {
		g.index = g.index % len(g.values)
	}

	return value
}

/* ---------------------------------- Grid ---------------------------------- */

type grid struct {
	values    []bool
	maxHeight int
}

func (g *grid) getMaxHeight() int {
	return g.maxHeight
}

func (g *grid) isClosed(x int, y int) bool {
	if x < 0 || y < 0 || x >= gridWidth {
		return true
	}

	index := y*gridWidth + x
	return g.values[index]
}

func (g *grid) addRock(rock *rock, origin *position) {
	maxRockHeight := origin[1] + rock.height - 1

	if maxRockHeight > g.maxHeight {
		g.maxHeight = maxRockHeight
	}

	for _, offset := range rock.offsets {
		x := origin[0] + offset[0]
		y := origin[1] + offset[1]
		index := y*gridWidth + x
		g.values[index] = true
	}
}

/* ------------------------------- Core Logic ------------------------------- */

func resolveGust(rock *rock, origin *position, grid *grid, gusts *gusts, simple bool) {
	move := gusts.step()

	if simple && (origin[0]+move < 0 || origin[0]+rock.width+move > gridWidth) {
		return
	}

	for _, offset := range rock.offsets {
		x := origin[0] + offset[0] + move
		y := origin[1] + offset[1]

		if grid.isClosed(x, y) {
			return
		}
	}

	origin[0] += move
}

func resolveDrop(rock *rock, origin *position, grid *grid) bool {
	for _, offset := range rock.offsets {
		x := origin[0] + offset[0]
		y := origin[1] + offset[1] - 1

		if grid.isClosed(x, y) {
			return false
		}
	}

	origin[1] -= 1
	return true
}

func resolveRock(rock *rock, grid *grid, gusts *gusts) *position {
	origin := position{2, grid.getMaxHeight() + 1}

	for i := 0; i < 4; i++ {
		resolveGust(rock, &origin, grid, gusts, true)
	}

	for {
		moved := resolveDrop(rock, &origin, grid)

		if !moved {
			break
		}

		resolveGust(rock, &origin, grid, gusts, false)
	}

	grid.addRock(rock, &origin)
	return &origin
}
