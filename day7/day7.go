package day7

import (
	"AdventOfCode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type state struct {
	disk    *tree
	currDir *node
}

type tree struct {
	root *node
}

func findDirs(maxSize int, rootDir *node) []*node {

	var res []*node

	if rootDir.calcSize() <= maxSize {
		res = append(res, rootDir)
	}

	for _, child := range rootDir.children {
		if child.isDir {
			res = append(res, findDirs(maxSize, child)...)
		}
	}

	return res
}

type node struct {
	parent         *node
	children       []*node
	name           string
	size           int
	isDir          bool
	calculatedSize int
}

func (n *node) calcSize() int {
	if n.isDir == false {
		return n.size
	} else {
		tot := 0
		for _, child := range n.children {
			tot += child.calcSize()
		}
		n.calculatedSize = tot
		return tot
	}
}

func newState() *state {
	rootNode := &node{
		parent:   nil,
		children: []*node{},
		name:     "/",
		isDir:    true,
	}
	return &state{
		currDir: rootNode,
		disk: &tree{
			root: rootNode,
		},
	}
}

func (s *state) exec(cmd Cmd) {

	switch cmd.cmd {
	case cd:
		s.cd(cmd.args[0])
	case ls:
		// do nothing
	}

}

func (s *state) cd(dirName string) {
	if dirName == "/" {
		s.currDir = s.disk.root
	} else if dirName == ".." {
		if s.currDir.parent == nil {
			panic("no parent dir")
		}
		s.currDir = s.currDir.parent
	} else {
		// check children
		for _, child := range s.currDir.children {
			if child.name == dirName {
				s.currDir = child
				return
			}
		}
		panic("dir not found")
	}
}

func (s *state) update(item listItem) {
	//assume we only list a directory once
	s.currDir.children = append(s.currDir.children, &node{
		parent:   s.currDir,
		children: []*node{},
		name:     item.name,
		size:     item.size,
		isDir:    item.isDir,
	})
}

type listItem struct {
	size  int
	name  string
	isDir bool
}

type Cmd struct {
	cmd  string
	args []string
}

const (
	cd string = "cd"
	ls string = "ls"
)

const CmdPrefix = "$"

type consoleLine struct {
	cmd      *Cmd
	listItem *listItem
}

func parseLine(line string) consoleLine {
	if strings.HasPrefix(line, CmdPrefix) {
		cmdStr := line[2:]
		split := strings.Split(cmdStr, " ")
		var args []string
		if len(split) > 1 {
			args = split[1:]
		}
		return consoleLine{
			cmd: &Cmd{
				cmd:  split[0],
				args: args,
			},
		}
	} else {
		split := strings.Split(line, " ")
		if split[0] == "dir" {
			return consoleLine{
				listItem: &listItem{
					size:  0,
					name:  split[1],
					isDir: true,
				},
			}
		} else {
			size, err := strconv.Atoi(split[0])
			if err != nil {
				panic("error converting file size to int")
			}
			return consoleLine{
				listItem: &listItem{
					size:  size,
					name:  split[1],
					isDir: false,
				},
			}
		}
	}
}

func SolveDay() {

	var lines = utils.ReadLines("day7/input.txt")

	var cState = newState()
	for _, line := range lines {
		out := parseLine(line)
		if out.cmd != nil {
			cState.exec(*out.cmd)
		} else {
			cState.update(*out.listItem)
		}
	}

	var dirs = findDirs(100000, cState.disk.root)

	sum := 0
	for _, dir := range dirs {
		sum += dir.calculatedSize
	}
	fmt.Println("Part 1:", sum)

	const TotSpace = 70000000
	const NeededSpace = 30000000

	freeSpace := TotSpace - cState.disk.root.calculatedSize

	delCandidates := findDirs(TotSpace, cState.disk.root)

	sort.SliceStable(delCandidates, func(i, j int) bool {
		return delCandidates[i].calculatedSize > delCandidates[j].calculatedSize
	})

	lastSize := 0
	for _, candidate := range delCandidates {

		if candidate.calculatedSize < NeededSpace-freeSpace {
			break
		}
		lastSize = candidate.calculatedSize
	}

	fmt.Println("Part 2:", lastSize)
}
