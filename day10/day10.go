package day10

import (
	"AdventOfCode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolveDay() {

	var lines = utils.ReadLines("day10/input.txt")

	nextSampling := 20

	cycle := 0
	reg := 1
	samplesSum := 0
	screen := ""
	for _, line := range lines {

		if line == "noop" {
			draw(cycle, reg, &screen)
			cycle++
			nextSampling, samplesSum = sample(cycle, nextSampling, samplesSum, reg)
		} else {
			split := strings.Split(line, " ")
			arg, _ := strconv.Atoi(split[1])

			draw(cycle, reg, &screen)
			cycle++
			nextSampling, samplesSum = sample(cycle, nextSampling, samplesSum, reg)

			draw(cycle, reg, &screen)
			cycle++
			nextSampling, samplesSum = sample(cycle, nextSampling, samplesSum, reg)

			reg += arg
		}

	}

	fmt.Println("Part 1:", samplesSum)
	fmt.Println("Part 2:", screen)

}

func draw(cycle int, reg int, screen *string) {

	x := cycle % 40
	if x == 0 {
		*screen += "\n"
	}
	var pixel string
	if x >= reg-1 && x <= reg+1 {
		pixel = "#"
	} else {
		pixel = "."
	}

	*screen += pixel
}

func sample(cycle int, nextSampling int, samplesSum int, reg int) (int, int) {
	if cycle == nextSampling {
		samplesSum += reg * nextSampling
		nextSampling += 40
	}
	return nextSampling, samplesSum
}
