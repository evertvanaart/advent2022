package day15

import (
	"advent2022/solutions/common"
	"fmt"
	"sort"
	"strings"
)

// Obviously we want to avoid having to allocate an entire 2D grid, especially given the large
// coordinate values in the main input. Similarly, we want to avoid the naive solution where we
// step through the target row column by column, and at every step check if the current position
// is within range of any of the sensors; due to the large number of columns, this would require
// several million steps for just one single row.
//
// Instead, we iterate through the sensors, and for each sensor check which range of columns
// is covered by the sensor's range (i.e., by the area covered in '#' characters in the sample).
// The math for this is actually very simple; if D is the Manhatten distance between S and B,
// we define W at row R as "D - |R - R_S|" (i.e., the Manhatten distance minus the distance
// in rows between R and S), and if W is not negative, row R overlaps the range of sensor S
// from columns C_S - W to C_S + W (inclusive on both ends; C_S is the column of S).
//
// Once we've obtained this list of ranges, we merge them as much as possible; we first sort
// the list by start columns, and then search for maximum non-overlapping ranges. We then
// simply count the sizes of the merged ranges, taking care to subtract one for each known
// beacon in the target row that's within the current merged range. This gives us a final
// runtime complexity of O(SlogS) (where S is the number of sensors).
//
// Note, the input was extended to include the target row (for the A part) and the search
// area bound (for the B part); this way I can run both the sample and the main input
// without having to change these values in code.

type beaconPositions map[int]bool

type columnRange struct {
	start int
	end   int
}

func (columnRange *columnRange) contains(col int) bool {
	return col >= columnRange.start && col <= columnRange.end
}

func (columnRange *columnRange) String() string {
	return fmt.Sprintf("[%d, %d]", columnRange.start, columnRange.end)
}

func getRangeForRow(line string, row int, beacons beaconPositions) *columnRange {
	fields := strings.Split(line, " ")
	sensorCol := common.ToInt(fields[2][2 : len(fields[2])-1])
	sensorRow := common.ToInt(fields[3][2 : len(fields[3])-1])
	beaconCol := common.ToInt(fields[8][2 : len(fields[8])-1])
	beaconRow := common.ToInt(fields[9][2:])

	manhattan := abs(sensorRow-beaconRow) + abs(sensorCol-beaconCol)
	halfWidth := manhattan - abs(row-sensorRow)

	if halfWidth < 0 {
		return nil
	}

	if beaconRow == row {
		beacons[beaconCol] = true
	}

	return &columnRange{sensorCol - halfWidth, sensorCol + halfWidth}
}

func mergeRanges(ranges []*columnRange) []*columnRange {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	outRanges := []*columnRange{}
	startColumn := ranges[0].start
	endColumn := ranges[0].end

	for _, currentRange := range ranges {
		if currentRange.start <= endColumn {
			endColumn = max(endColumn, currentRange.end)
		} else {
			outRange := &columnRange{startColumn, endColumn}
			outRanges = append(outRanges, outRange)
			startColumn = currentRange.start
			endColumn = currentRange.end
		}
	}

	outRange := &columnRange{startColumn, endColumn}
	outRanges = append(outRanges, outRange)
	return outRanges
}

func count(ranges []*columnRange, beacons beaconPositions) int {
	totalSum := 0

	for _, currentRange := range ranges {
		totalSum += currentRange.end - currentRange.start + 1

		for beacon := range beacons {
			if currentRange.contains(beacon) {
				totalSum -= 1
			}
		}
	}

	return totalSum
}

func SolveA(lines []string) common.Solution {
	targetRow := common.ToInt(lines[0])
	beacons := beaconPositions{}
	ranges := []*columnRange{}

	for _, line := range lines[2:] {
		columnRange := getRangeForRow(line, targetRow, beacons)

		if columnRange != nil {
			ranges = append(ranges, columnRange)
		}
	}

	mergedRanges := mergeRanges(ranges)
	coveredCount := count(mergedRanges, beacons)
	return common.ToIntSolution(coveredCount)
}
