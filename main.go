package main

import (
	"AdventOfCode2022/day1"
	"AdventOfCode2022/day2"
	"AdventOfCode2022/day3"
	"AdventOfCode2022/day4"
	"AdventOfCode2022/day5"
	"fmt"
)

func main() {

	var solvers = []func(){
		day1.SolveDay,
		day2.SolveDay,
		day3.SolveDay,
		day4.SolveDay,
		day5.SolveDay,
	}

	for i, solver := range solvers {
		fmt.Printf("\n== Day %d ==\n", i+1)
		solver()
	}
}
