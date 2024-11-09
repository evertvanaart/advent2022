package day13

import (
	"advent2022/solutions/common"
	"fmt"
)

/* --------------------------------- Element -------------------------------- */

type element interface {
	compare(element) int
}

/* ----------------------------- Integer element ---------------------------- */

type intElement struct {
	element
	value int
}

func (thisElement intElement) String() string {
	return fmt.Sprint(thisElement.value)
}

func (thisElement intElement) compare(otherElement element) int {
	switch otherElement := otherElement.(type) {
	case intElement:
		return thisElement.value - otherElement.value
	case listElement:
		wrappedElement := listElement{values: []element{thisElement}}
		return wrappedElement.compare(otherElement)
	}

	panic("Unsupported element type")
}

/* ------------------------------ List element ------------------------------ */

type listElement struct {
	element
	values []element
}

func (thisElement listElement) String() string {
	return fmt.Sprint(thisElement.values)
}

func (thisElement listElement) compareLists(otherElement listElement) int {
	for i := 0; i < len(otherElement.values); i++ {
		if i >= len(thisElement.values) {
			return -1
		}

		thisNestedElement := thisElement.values[i]
		otherNestedElement := otherElement.values[i]
		result := thisNestedElement.compare(otherNestedElement)

		if result != 0 {
			return result
		}
	}

	return len(thisElement.values) - len(otherElement.values)
}

func (thisElement listElement) compare(otherElement element) int {
	switch otherElement := otherElement.(type) {
	case listElement:
		return thisElement.compareLists(otherElement)
	case intElement:
		wrappedElement := listElement{values: []element{otherElement}}
		return thisElement.compare(wrappedElement)
	}

	panic("Unsupported element type")
}

/* --------------------------------- Parsing -------------------------------- */

func parseList(str string) listElement {
	if len(str) == 0 {
		return listElement{values: []element{}}
	}

	elements := []element{}
	elementStart := 0
	depth := 0

	for index, char := range str {
		if char == '[' {
			depth += 1
		} else if char == ']' {
			depth -= 1
		} else if char == ',' && depth == 0 {
			element := parseElement(str[elementStart:index])
			elements = append(elements, element)
			elementStart = index + 1
		}
	}

	lastElement := parseElement(str[elementStart:])
	elements = append(elements, lastElement)
	return listElement{values: elements}
}

func parseElement(str string) element {
	if str[0] == '[' {
		innerStr := str[1 : len(str)-1]
		return parseList(innerStr)
	} else {
		intValue := common.ToInt(str)
		return intElement{value: intValue}
	}
}
