package day10

import (
	"advent2022/solutions/common"
	"strings"
)

// Simple and fast; use a CPU struct to track the clock cycle, register value, current signal
// strength, and upcoming sample points; split addition instructions into two separate steps,
// one that adds zero and one that adds the actual value; after every step, check if we
// reached a sample point, and if so, update the signal strength accordingly. The only
// real challenge in the A part is the high risk of off-by-one errors.

type cpuA struct {
	clock    int
	value    int
	strength int
	samples  []int
}

func createCpuA() *cpuA {
	return &cpuA{0, 1, 0, []int{20, 60, 100, 140, 180, 220}}
}

func (cpu *cpuA) step(diff int) {
	cpu.clock += 1

	if cpu.clock == cpu.samples[0] {
		cpu.strength += cpu.clock * cpu.value
		cpu.samples = cpu.samples[1:]
	}

	cpu.value += diff
}

func (cpu *cpuA) process(line string) {
	if line == "noop" {
		cpu.step(0)
		return
	}

	fields := strings.Split(line, " ")
	diff := common.ToInt(fields[1])

	cpu.step(0)
	cpu.step(diff)
}

func (cpu *cpuA) done() bool {
	return len(cpu.samples) == 0
}

func SolveA(lines []string) common.Solution {
	cpu := createCpuA()

	for _, line := range lines {
		cpu.process(line)

		if cpu.done() {
			break
		}
	}

	return common.ToIntSolution(cpu.strength)
}
