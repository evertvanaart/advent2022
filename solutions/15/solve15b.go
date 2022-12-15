package day15

import (
	"advent2022/solutions/common"
	"strings"
)

// The knowledge that there will be exactly one "open" space within the specified area means that
// this open space must be located immediately next to the area of at least one sensor, i.e. if a
// sensor S has Manhatten distance D between S and its closest beacon B, the open space must be
// located at distance D + 1 for at least one sensor. Since sensor areas have a diamond shape,
// the set of D + 1 cells can be expressed as four non-overlapping diagonals, two each in both
// directions (i.e., two can be expressed as y = x + a, and two as y = -x + a).
//
// Conveniently ignoring some unlikely edge cases, we can assume that the point we are looking
// for is located at the intersection point of two _perpendicular_ diagonals. When we find such
// an intersection (inside the specified area), we can check all sensors to see whether or not
// any of their areas cover this intersection point; if not, we've found the space we are looking
// for. This technically gives us O(S^3) complexity, but since most pairs of diagonals do not
// intersect (only 134 unique intersections in the main input of 32 lines), the total execution
// time is still fairly low, well below 100 microseconds on my reference machine.
//
// Illustration: the / and \ fields represent the four diagonals surrounding a sensor area:
//
//        \	           Each diagonal can be expressed using a direction (positive or negative,
//       /#\           i.e. '/' or '\'), a minimum and maximum Y values (inclusive), and an
//      /###\          offset, i.e. the value of y at x = 0. Note that we are only interested
//     /####B\         in diagonal pairs that intersect in the middle of a cell, as opposed to
//    /###S###/        at the intersection between four cells; this is the case if the diffe-
//     \#####/         rence between the two offsets is even. We use two separate collections
// ^    \###/          to track positive and negative diagonals, since we are only interested
// Y     \#/           in intersections between perpendicular diagonals.
//  X >   \

type diagonal struct {
	offset int
	minY   int
	maxY   int
}

type sensor struct {
	x int
	y int
	d int
}

type position [2]int

func intersect(pos *diagonal, neg *diagonal) (position, bool) {
	offsetDiff := neg.offset - pos.offset

	if (offsetDiff % 2) != 0 {
		return position{0, 0}, false
	}

	x := offsetDiff / 2
	y := pos.offset + x

	if y >= pos.minY && y <= pos.maxY && y >= neg.minY && y <= neg.maxY {
		return position{x, y}, true
	} else {
		return position{0, 0}, false
	}
}

func findIntersections(posDiagonals []*diagonal, negDiagonals []*diagonal,
	sensors []*sensor) map[position]bool {
	intersections := map[position]bool{}

	for _, pos := range posDiagonals {
		for _, neg := range negDiagonals {
			intersection, intersects := intersect(pos, neg)

			if intersects {
				intersections[intersection] = true
			}
		}
	}

	return intersections
}

func parseDiagonals(line string) ([]*diagonal, []*diagonal, *sensor) {
	fields := strings.Split(line, " ")
	sensorX := common.ToInt(fields[2][2 : len(fields[2])-1])
	sensorY := common.ToInt(fields[3][2 : len(fields[3])-1])
	beaconX := common.ToInt(fields[8][2 : len(fields[8])-1])
	beaconY := common.ToInt(fields[9][2:])

	basePosOffset := sensorY - sensorX
	baseNegOffset := sensorY + sensorX
	manhattan := abs(sensorX-beaconX) + abs(sensorY-beaconY)

	pos := []*diagonal{
		{basePosOffset + manhattan + 1, sensorY, sensorY + manhattan},
		{basePosOffset - manhattan - 1, sensorY - manhattan, sensorY},
	}

	neg := []*diagonal{
		{baseNegOffset + manhattan + 1, sensorY + 1, sensorY + manhattan + 1},
		{baseNegOffset - manhattan - 1, sensorY - manhattan - 1, sensorY - 1},
	}

	sensor := &sensor{sensorX, sensorY, manhattan}
	return pos, neg, sensor
}

func check(sensors []*sensor, p position) bool {
	for _, sensor := range sensors {
		dist := abs(p[0]-sensor.x) + abs(p[1]-sensor.y)

		if dist <= sensor.d {
			return false
		}
	}

	return true
}

func SolveB(lines []string) common.Solution {
	bound := common.ToInt(lines[1])
	posDiagonals := []*diagonal{}
	negDiagonals := []*diagonal{}
	sensors := []*sensor{}

	for _, line := range lines[2:] {
		pos, neg, sensor := parseDiagonals(line)
		posDiagonals = append(posDiagonals, pos...)
		negDiagonals = append(negDiagonals, neg...)
		sensors = append(sensors, sensor)
	}

	intersections := findIntersections(posDiagonals, negDiagonals, sensors)

	for p := range intersections {
		if p[0] >= 0 && p[0] <= bound && p[1] >= 0 && p[1] <= bound {
			result := check(sensors, p)

			if result {
				solution := 4000000*int64(p[0]) + int64(p[1])
				return common.ToInt64Solution(solution)
			}
		}
	}

	return common.ToIntSolution(-1)
}
