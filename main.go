package main

import (
	"advent2022/solutions/day1"
	"bufio"
	"fmt"
	"os"
	"time"
)

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := readLines("input/1/input.txt")

	if err != nil {
		panic(fmt.Sprintf("Failed to read input: %s", err))
	}

	startTime := time.Now()
	solution := day1.SolveB(lines)
	duration := time.Since(startTime)

	fmt.Printf("Elapsed time: %v\n", duration)
	fmt.Printf("Solution: %v\n", solution)
}
