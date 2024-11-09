package day07

import (
	"advent2022/solutions/common"
	"fmt"
	"strings"
)

const (
	rootPath  = "/"
	cdRootCmd = "$ cd /"
	cdUpCmd   = "$ cd .."
	listCmd   = "$ ls"
	cdPrefix  = "$ cd "
	dirPrefix = "dir"
	noParent  = "none"
)

type directory struct {
	parent string
	listed int
	size   int
}

type directories map[string]*directory

func (directory *directory) String() string {
	return fmt.Sprintf("directory{parent: %s, size: %d}", directory.parent, directory.size)
}

func createDirectory(parent string) *directory {
	return &directory{parent: parent, listed: 0, size: 0}
}

func createPath(parent string, dir string) string {
	if parent == rootPath {
		return fmt.Sprintf("/%s", dir)
	} else {
		return fmt.Sprintf("%s/%s", parent, dir)
	}
}

func initializeDirectories() directories {
	return directories{rootPath: &directory{parent: noParent, listed: 0, size: 0}}
}

func getDirectoryName(line string) string {
	if !strings.HasPrefix(line, cdPrefix) {
		panic(fmt.Sprintf("Unknown command '%s'", line))
	}

	return line[len(cdPrefix):]
}

func processCommand(line string, directories directories, currentPath string) string {
	switch line {
	case cdRootCmd:
		return rootPath
	case cdUpCmd:
		return directories[currentPath].parent
	case listCmd:
		directories[currentPath].listed += 1
		return currentPath
	default: // $ cd <x>
		newPath := createPath(currentPath, getDirectoryName(line))
		directories[newPath] = createDirectory(currentPath)
		return newPath
	}
}

func processEntry(line string, directories directories, currentPath string) {
	if strings.HasPrefix(line, dirPrefix) {
		return
	}

	if directories[currentPath].listed != 1 {
		return
	}

	fields := strings.Split(line, " ")
	size := common.ToInt(fields[0])
	path := currentPath

	for path != noParent {
		directories[path].size += size
		path = directories[path].parent
	}
}
