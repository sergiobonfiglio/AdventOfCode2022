package day8

import (
	"AdventOfCode2022/utils"
	"fmt"
	"strconv"
)

func SolveDay() {

	lines := utils.ReadLines("day8/input.txt")

	forest := make([][]int, len(lines))
	for r, line := range lines {
		forest[r] = make([]int, len(line))
		for c, t := range line {
			h, err := strconv.Atoi(string(t))
			if err != nil {
				panic("error parsing height")
			}

			forest[r][c] = h
		}
	}

	visibleTrees := 2*len(forest) + 2*(len(forest[0])-2)
	for r := 1; r < len(forest)-1; r++ {
		for c := 1; c < len(forest[r])-1; c++ {

			if !isHidden(forest, r, c) {
				visibleTrees++
			}

		}
	}

	fmt.Println("Part 1:", visibleTrees)

	max := 0
	for r := 1; r < len(forest)-1; r++ {
		for c := 1; c < len(forest[r])-1; c++ {
			score := scenicScore(forest, r, c)
			if score > max {
				max = score
			}
		}
	}

	fmt.Println("Part 2:", max)
}

func isHidden(forest [][]int, r int, c int) bool {

	h := forest[r][c]

	isHidden := true
	// up
	isHiddenUp := false
	for i := 0; i < r; i++ {
		isHiddenUp = isHiddenUp || forest[i][c] >= h
	}
	isHidden = isHiddenUp

	//down
	isHiddenDown := false
	for i := r + 1; i < len(forest) && isHidden; i++ {
		isHiddenDown = isHiddenDown || forest[i][c] >= h
	}
	isHidden = isHidden && isHiddenDown

	//left
	isHiddenLeft := false
	for i := 0; i < c && isHidden; i++ {
		isHiddenLeft = isHiddenLeft || forest[r][i] >= h
	}
	isHidden = isHidden && isHiddenLeft

	//right
	isHiddenRight := false
	for i := c + 1; i < len(forest[r]) && isHidden; i++ {
		isHiddenRight = isHiddenRight || forest[r][i] >= h
	}
	isHidden = isHidden && isHiddenRight

	return isHidden
}

func scenicScore(forest [][]int, r int, c int) int {

	h := forest[r][c]

	lx := 0
	//left
	for i := r - 1; i >= 0; i-- {
		lx++
		if forest[i][c] >= h {
			break
		}
	}

	if lx == 0 {
		return 0
	}

	rx := 0
	//right
	for i := r + 1; i < len(forest); i++ {
		rx++
		if forest[i][c] >= h {
			break
		}
	}

	if rx == 0 {
		return 0
	}

	up := 0
	//up
	for i := c - 1; i >= 0; i-- {
		up++
		if forest[r][i] >= h {
			break
		}
	}

	if up == 0 {
		return 0
	}

	dw := 0
	//right
	for i := c + 1; i < len(forest[r]); i++ {
		dw++
		if forest[r][i] >= h {
			break
		}
	}

	return lx * rx * up * dw
}
