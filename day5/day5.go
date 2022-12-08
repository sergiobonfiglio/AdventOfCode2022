package day5

import (
	"AdventOfCode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type move struct {
	qnt  int
	from int
	to   int
}

func SolveDay() {

	var lines = utils.ReadLines("day5/input.txt")

	stacksReading := true
	stacks := map[int]string{}
	var moves []move
	for _, line := range lines {
		if len(line) == 0 || strings.HasPrefix(line, " 1 ") {
			stacksReading = false
			continue
		}

		if stacksReading {
			for i, c := range line {
				if c >= 'A' && c <= 'Z' {
					stackIx := (i / 4) + 1
					stacks[stackIx] = stacks[stackIx] + string(c)
				}
			}

		} else {
			// moves
			split := strings.Split(line, " ")
			qnt, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])
			moves = append(moves, move{
				qnt,
				from,
				to,
			})
		}
	}

	origStacks := map[int]string{}
	for i, s := range stacks {
		origStacks[i] = s
	}

	//execute moves v1
	for _, m := range moves {
		for i := 0; i < m.qnt; i++ {
			x := stacks[m.from][0]
			stacks[m.from] = stacks[m.from][1:]
			stacks[m.to] = string(x) + stacks[m.to]
		}
	}

	//get top crates
	top := getTopCrates(stacks)
	fmt.Println("Part 1:", top)

	//execute moves v2
	for _, m := range moves {
		x := origStacks[m.from][0:m.qnt]
		origStacks[m.from] = origStacks[m.from][m.qnt:]
		origStacks[m.to] = x + origStacks[m.to]
	}
	//get top crates
	top = getTopCrates(origStacks)
	fmt.Println("Part 2:", top)
}

func getTopCrates(stacks map[int]string) string {
	top := ""
	for i := 0; i < len(stacks); i++ {
		if s, found := stacks[i+1]; found {
			top += string(s[0])
		}
	}
	return top
}
