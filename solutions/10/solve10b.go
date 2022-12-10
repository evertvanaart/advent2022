package day10

import (
	"advent2022/solutions/common"
	"fmt"
	"strings"
)

// In keeping with the object-oriented approach of the A part solution, we model both the CPU and
// the CRT as objects; every instruction results in one or two steps; after every step, check if
// the current CRT position lies inside the CPU sprite, and if so change the corresponding pixel
// to '#'. We use a one-dimensional byte array to track CRT pixels, and only convert this array
// to the six output lines at the very end.

/* ----------------------------------- CRT ---------------------------------- */

const crtWidth = 40
const crtHeight = 6
const crtSize = crtWidth * crtHeight

type crt struct {
	grid []byte
	pos  int
}

func createCrt() *crt {
	grid := make([]byte, crtSize)

	for index := range grid {
		grid[index] = '.'
	}

	return &crt{grid, 0}
}

func (crt *crt) step() {
	crt.pos += 1
}

func (crt *crt) light() {
	crt.grid[crt.pos] = '#'
}

func (crt *crt) toStrings() []string {
	fullString := string(crt.grid)
	lines := []string{}

	for i := 0; i < crtSize; i += crtWidth {
		lines = append(lines, fullString[i:i+crtWidth])
	}

	return lines
}

/* ----------------------------------- CPU ---------------------------------- */

type cpuB struct {
	value int
}

func createCpuB() *cpuB {
	return &cpuB{1}
}

func (cpu *cpuB) inSprite(pos int) bool {
	lowerBound := cpu.value - 1
	upperBound := cpu.value + 1
	return pos >= lowerBound && pos <= upperBound
}

func (cpu *cpuB) step(diff int) {
	cpu.value += diff
}

func (cpu *cpuB) process(line string) []int {
	if line == "noop" {
		return []int{0}
	}

	fields := strings.Split(line, " ")
	diff := common.ToInt(fields[1])
	return []int{0, diff}
}

/* ------------------------------- Main logic ------------------------------- */

func step(line string, cpu *cpuB, crt *crt) {
	diffs := cpu.process(line)

	for _, diff := range diffs {
		if cpu.inSprite(crt.pos % crtWidth) {
			crt.light()
		}

		cpu.step(diff)
		crt.step()
	}
}

func SolveB(lines []string) common.Solution {
	cpu := createCpuB()
	crt := createCrt()

	for _, line := range lines {
		step(line, cpu, crt)
	}

	outLines := crt.toStrings()
	outBlock := strings.Join(outLines, "\n")
	output := fmt.Sprintf("\n%s", outBlock)
	return common.ToStringSolution(output)
}
