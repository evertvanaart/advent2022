package day16

import (
	"advent2022/solutions/common"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/* ---------------------------------- Valve --------------------------------- */

type valve struct {
	index     int
	flow      int
	neighbors connections
}

func (v *valve) String() string {
	return fmt.Sprintf("{%d, %d, %v}", v.index, v.flow, v.neighbors)
}

func (v *valve) findNextIndex(valveIndex int) int {
	for index, connection := range v.neighbors {
		if connection.index == valveIndex {
			return index
		}
	}

	return -1
}

func (v *valve) simplify(toRemove int, toAdd int, distance int) {
	oldIndex := v.findNextIndex(toRemove)
	oldDistance := v.neighbors[oldIndex].distance
	newDistance := oldDistance + distance
	newConnection := &connection{toAdd, newDistance}
	v.neighbors[oldIndex] = newConnection
}

/* -------------------------------- Valve Map ------------------------------- */

type valves []*valve

func (valves valves) simplify(startIndex int) {
	for index, valve := range valves {
		if valve.index == startIndex ||
			valve.flow > 0 ||
			len(valve.neighbors) != 2 {
			continue
		}

		neighborA := valve.neighbors[0].index
		neighborB := valve.neighbors[1].index
		distanceA := valve.neighbors[0].distance
		distanceB := valve.neighbors[1].distance

		valves[neighborA].simplify(index, neighborB, distanceB)
		valves[neighborB].simplify(index, neighborA, distanceA)
	}
}

func (valves valves) totalFlow() int {
	sum := 0

	for _, valve := range valves {
		sum += valve.flow
	}

	return sum
}

/* ---------------------------- Score estimation ---------------------------- */

func (valves valves) potentialScore(timeLeft int, timeStep int, opened []bool) int {
	flows := []int{}

	for _, valve := range valves {
		if valve.flow > 0 && !opened[valve.index] {
			flows = append(flows, valve.flow)
		}
	}

	sort.Slice(flows, func(i, j int) bool { return flows[i] > flows[j] })

	time := timeLeft - timeStep
	potentialScore := 0

	for index := 0; index < len(flows) && time > 0; index++ {
		potentialScore += time * flows[index]
		time -= timeStep
	}

	return potentialScore
}

/* ------------------------------- Connections ------------------------------ */

type connection struct {
	index    int
	distance int
}

func (c *connection) String() string {
	return fmt.Sprintf("{%d, %d}", c.index, c.distance)
}

type connections []*connection

/* -------------------------------- Candidate ------------------------------- */

type candidate struct {
	index    int
	distance int
	score    int
}

func (c *candidate) String() string {
	return fmt.Sprintf("{%d, %d, %d}", c.index, c.distance, c.score)
}

/* --------------------------------- Parsing -------------------------------- */

var linePattern = regexp.MustCompile(
	`Valve ([A-Z]+) has flow rate=(\d+); tunnel[s]? lead[s]? to valve[s]? (.+)`)

func parse(line string, label2index map[string]int) *valve {
	matches := linePattern.FindStringSubmatch(line)

	label := matches[1]
	index := label2index[label]
	flow := common.ToInt(matches[2])
	neighborLabels := strings.Split(matches[3], ", ")
	connections := make([]*connection, len(neighborLabels))

	for index, neighborLabel := range neighborLabels {
		connections[index] = &connection{label2index[neighborLabel], 1}
	}

	return &valve{index, flow, connections}
}

func createValves(lines []string) (valves, int) {
	label2index := map[string]int{}

	for _, line := range lines {
		label2index[line[6:8]] = len(label2index)
	}

	startIndex := label2index["AA"]
	valves := make(valves, len(lines))

	for index, line := range lines {
		valve := parse(line, label2index)
		valves[index] = valve
	}

	valves.simplify(startIndex)

	return valves, startIndex
}
