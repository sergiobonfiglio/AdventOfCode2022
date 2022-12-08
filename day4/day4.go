package day4

import (
	"AdventOfCode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

func newRange(x string) Range {
	split := strings.Split(x, "-")
	min, err := strconv.Atoi(split[0])
	if err != nil {
		panic("parse int error")
	}
	max, err := strconv.Atoi(split[1])
	if err != nil {
		panic("parse int error")
	}
	return Range{
		min: min,
		max: max,
	}
}

func (r1 Range) contains(r2 Range) bool {
	return r1.min >= r2.min && r1.max <= r2.max
}

func (r1 Range) overlaps(r2 Range) bool {
	return r1.min <= r2.max && r1.max >= r2.min
}
func SolveDay() {

	var data = utils.ReadLines("day4/input.txt")

	contained := 0
	overlapped := 0
	for _, line := range data {
		split := strings.Split(line, ",")
		elf1 := split[0]
		elf2 := split[1]

		rng1 := newRange(elf1)
		rng2 := newRange(elf2)

		if rng1.contains(rng2) || rng2.contains(rng1) {
			contained++
		}
		if rng1.overlaps(rng2) {
			overlapped++
		}
	}

	fmt.Println("Part 1:", contained)
	fmt.Println("Part 2:", overlapped)

}
