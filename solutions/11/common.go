package day11

import (
	"advent2022/solutions/common"
	"fmt"
	"sort"
	"strings"
)

/* -------------------------------- Constants ------------------------------- */

const linesPerMonkey = 7

const (
	lineNoStartingItems = 1
	lineNoOperation     = 2
	lineNoDivisor       = 3
	lineNoTargetTrue    = 4
	lineNoTargetFalse   = 5
)

const (
	offsetStartItems = len("  Starting items: ")
	offsetOperation  = len("  Operation: new = ")
	offsetDivisor    = len("  Test: divisible by ")
)

const (
	operationSquare   = 0
	operationMultiply = 1
	operationAdd      = 2
)

/* ---------------------------------- Item ---------------------------------- */

type item struct {
	value  int
	target int
}

/* --------------------------- Adjustment function -------------------------- */

type adjust func(int) int

/* --------------------------------- Monkey --------------------------------- */

type monkey struct {
	items       []int
	operation   int
	operand     int
	divisor     int
	targetTrue  int
	targetFalse int
	inspected   int
}

func (monkey *monkey) addItem(itemValue int) {
	monkey.items = append(monkey.items, itemValue)
}

func (monkey *monkey) processItem(itemValue int, adjust adjust) *item {
	monkey.inspected += 1
	newValue := itemValue

	switch monkey.operation {
	case operationSquare:
		newValue = itemValue * itemValue
	case operationMultiply:
		newValue = itemValue * monkey.operand
	case operationAdd:
		newValue = itemValue + monkey.operand
	}

	newValue = adjust(newValue)

	if (newValue % monkey.divisor) == 0 {
		return &item{newValue, monkey.targetTrue}
	} else {
		return &item{newValue, monkey.targetFalse}
	}
}

func (monkey *monkey) processTurn(adjust adjust) []*item {
	thrownItems := make([]*item, len(monkey.items))

	for index, item := range monkey.items {
		thrownItems[index] = monkey.processItem(item, adjust)
	}

	monkey.items = []int{}
	return thrownItems
}

/* --------------------------------- Parsing -------------------------------- */

func parseOperation(line string) (int, int) {
	operation := line[offsetOperation:]

	if operation == "old * old" {
		return operationSquare, 0
	} else if strings.HasPrefix(operation, "old * ") {
		operand := common.ToInt(operation[len("old * "):])
		return operationMultiply, operand
	} else if strings.HasPrefix(operation, "old + ") {
		operand := common.ToInt(operation[len("old + "):])
		return operationAdd, operand
	}

	panic(fmt.Sprintf("Unsupported operation: %s", line))
}

func parseDivisor(line string) int {
	return common.ToInt(line[offsetDivisor:])
}

func parseTarget(line string) int {
	lastSpaceIndex := strings.LastIndex(line, " ")
	return common.ToInt(line[lastSpaceIndex+1:])
}

func parseMonkey(lines []string) *monkey {
	itemsString := lines[lineNoStartingItems][offsetStartItems:]
	itemsFields := strings.Split(itemsString, ", ")
	itemsValues := common.ToIntN(itemsFields)

	operation, operand := parseOperation(lines[lineNoOperation])
	divisor := parseDivisor(lines[lineNoDivisor])
	targetTrue := parseTarget(lines[lineNoTargetTrue])
	targetFalse := parseTarget(lines[lineNoTargetFalse])

	return &monkey{
		items:       itemsValues,
		operation:   operation,
		operand:     operand,
		divisor:     divisor,
		targetTrue:  targetTrue,
		targetFalse: targetFalse,
		inspected:   0,
	}
}

/* ------------------------------- Main logic ------------------------------- */

func determineOutput(monkeys []*monkey) int64 {
	inspected := make([]int, len(monkeys))

	for index, monkey := range monkeys {
		inspected[index] = monkey.inspected
	}

	sort.Ints(inspected)

	value1 := int64(inspected[len(inspected)-1])
	value2 := int64(inspected[len(inspected)-2])

	return value1 * value2
}

func processRound(monkeys []*monkey, adjust adjust) {
	for _, monkey := range monkeys {
		thrownItems := monkey.processTurn(adjust)

		for _, thrownItem := range thrownItems {
			monkeys[thrownItem.target].addItem(thrownItem.value)
		}
	}
}
