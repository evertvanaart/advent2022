package day07

import "advent2022/solutions/common"

// To avoid having to deal with tree traversal, we keep a flat map of all directories, using
// their absolute paths as keys. Each directory object contains the path of its parent directory
// (although we could of course also determine this dynamically from the path), the combined size
// of all objects in this directory and its subdirectories, and a counter tracking how many times
// we've called "ls" on this directory (to avoid counting files more than once, although I don't
// believe the input actually does this at any point).
//
// While looping through the input, we keep track of the path of the current directory, which is
// updated via the "cd" commands. Whenever we encounter a file, we add its size to the total size
// of the current directory _and_ all of its parent directories (provided that this is the first
// time "ls" was called for the current directory). Note that "dir <name>" lines are ignored; we
// will create those directories only when we step into them, and the input doesn't appear to
// try to step into any non-existant directories.
//
// After processing all input lines, we'll have a flat map of all directories with their total
// size (including subdirectories), so all we need to do then is to sum up the sizes of the
// directories that are smaller than the provided limit.
//
// One possible optimization would be to use an array instead of a map, and identify directories
// and their parents by their array index instead of by their path. Doing so would not change
// the theoretical runtime complexity, but would still somewhat speed up accessing the values.

func SolveA(lines []string) common.Solution {
	directories := initializeDirectories()
	currentPath := rootPath

	for _, line := range lines {
		if line[0] == '$' {
			currentPath = processCommand(line, directories, currentPath)
		} else {
			processEntry(line, directories, currentPath)
		}
	}

	sum := 0

	for _, dir := range directories {
		if dir.size < 100000 {
			sum += dir.size
		}
	}

	return common.ToIntSolution(sum)
}
