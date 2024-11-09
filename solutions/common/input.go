package common

import (
	"bufio"
	"fmt"
	"os"
)

// Reads all lines from the indicated input file, panics in case of problems
func ReadLines(day string, name string) []string {
	filename := fmt.Sprintf("input/%s/%s.txt", day, name)

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

// Parses the input arguments
func ParseArgs() (string, string, string) {
	args := os.Args[1:]

	if len(args) == 2 {
		return args[0], args[1], ""
	} else if len(args) == 3 {
		return args[0], args[1], args[2]
	} else {
		printUsage()
		os.Exit(1)
	}

	return "", "", ""
}

// Extract the day from the task string
func ParseTask(task string) string {
	if len(task) != 3 {
		fmt.Printf("Task '%s' does not match expected length\n", task)
		printUsage()
		os.Exit(1)
	}

	return task[0:2]
}

// Prints usage
func printUsage() {
	fmt.Println("Usage: go run . <task> <input> [--profile]")
	fmt.Println(" <task>     Day number (two digits) plus part ('a' or 'b')")
	fmt.Println(" <input>    Input file base name, e.g. 'input' or 'sample'")
	fmt.Println(" --profile  Run solution multiple times and compute average duration")
	fmt.Println("Example: go run . 01a sample")
}
