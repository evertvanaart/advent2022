package day04

import (
	"advent2022/solutions/common"
	"strings"
)

type sections struct {
	start int
	end   int
}

func toSections(field string) *sections {
	fields := strings.Split(field, "-")
	start := common.ToInt(fields[0])
	end := common.ToInt(fields[1])
	return &sections{start, end}
}
