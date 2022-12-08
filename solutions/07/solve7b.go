package day07

import "advent2022/solutions/common"

// Constructing the flat map is done in exactly the same way as in the A part; once we
// have this map, extracing the required minimum directory size is straightforward.

func SolveB(lines []string) common.Solution {
	directories := initializeDirectories()
	currentPath := rootPath

	for _, line := range lines {
		if line[0] == '$' {
			currentPath = processCommand(line, directories, currentPath)
		} else {
			processEntry(line, directories, currentPath)
		}
	}

	totalDiskSize := 70000000
	usedDiskSpace := directories[rootPath].size
	freeDiskSpace := totalDiskSize - usedDiskSpace
	spaceToFree := 30000000 - freeDiskSpace
	minDirSize := totalDiskSize

	for _, dir := range directories {
		if dir.size >= spaceToFree && dir.size < minDirSize {
			minDirSize = dir.size
		}
	}

	return common.ToIntSolution(minDirSize)
}
