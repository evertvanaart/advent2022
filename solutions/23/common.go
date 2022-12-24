package day23

import (
	"fmt"
	"math"
)

/* ---------------------------- Types & Constants --------------------------- */

type position [2]int

type rule struct {
	checkOffsets [3]position
	moveOffset   position
}

func createRule(co1 position, co2 position, co3 position, mo position) rule {
	return rule{[3]position{co1, co2, co3}, mo}
}

var offsetN = position{-1, 0}
var offsetW = position{0, -1}
var offsetS = position{1, 0}
var offsetE = position{0, 1}
var offsetNW = position{-1, -1}
var offsetSW = position{1, -1}
var offsetNE = position{-1, 1}
var offsetSE = position{1, 1}

var allOffsets = []position{
	offsetN,
	offsetW,
	offsetS,
	offsetE,
	offsetNW,
	offsetSW,
	offsetNE,
	offsetSE,
}

/* -------------------------------- Utilities ------------------------------- */

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func print(elves map[position]bool) {
	tl, br := boundingRect(elves)
	fmt.Println(tl)

	for row := tl[0]; row <= br[0]; row++ {
		for col := tl[1]; col <= br[1]; col++ {
			if elves[position{row, col}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}

func boundingRect(elves map[position]bool) (position, position) {
	minRow := math.MaxInt16
	maxRow := math.MinInt16
	minCol := math.MaxInt16
	maxCol := math.MinInt16

	for pos := range elves {
		minRow = min(minRow, pos[0])
		maxRow = max(maxRow, pos[0])
		minCol = min(minCol, pos[1])
		maxCol = max(maxCol, pos[1])
	}

	return position{minRow, minCol}, position{maxRow, maxCol}
}

func checkSurrounding(elves map[position]bool, elf position) bool {
	for _, offset := range allOffsets {
		offsetPos := position{elf[0] + offset[0], elf[1] + offset[1]}
		_, exists := elves[offsetPos]

		if exists {
			return true
		}
	}

	return false
}

/* ------------------------------- Core logic ------------------------------- */

func consider(elves map[position]bool, elf position, rules []rule) position {
	if !checkSurrounding(elves, elf) {
		return elf
	}

	for _, rule := range rules {
		checkPassed := true

		for _, offset := range rule.checkOffsets {
			offsetPos := position{elf[0] + offset[0], elf[1] + offset[1]}
			_, exists := elves[offsetPos]

			if exists {
				checkPassed = false
				break
			}
		}

		if checkPassed {
			offset := rule.moveOffset
			return position{elf[0] + offset[0], elf[1] + offset[1]}
		}
	}

	return elf
}

func process(elves map[position]bool, rules []rule, ruleOffset int) bool {
	considered := map[position]position{}
	stepRules := make([]rule, len(rules))

	for index := range rules {
		stepRules[index] = rules[(ruleOffset+index)%len(rules)]
	}

	for elf := range elves {
		consideredPos := consider(elves, elf, stepRules)
		_, exists := considered[consideredPos]

		if exists {
			delete(considered, consideredPos)
		} else {
			considered[consideredPos] = elf
		}
	}

	anyMoved := false

	for targetPos, sourcePos := range considered {
		if targetPos == sourcePos {
			continue
		}

		delete(elves, sourcePos)
		elves[targetPos] = true
		anyMoved = true
	}

	return anyMoved
}
