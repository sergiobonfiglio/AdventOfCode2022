package main

import (
	"AdventOfCode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	SolveDay()
}

func SolveDay() {

	lines := utils.ReadLines("day14/input.txt")

	var bounds *rectBound
	var cells []utils.Position
	for _, line := range lines {
		rockPositions := strings.Split(line, "->")
		var verts []utils.Position
		for _, rockPosStr := range rockPositions {
			rockPosStr = strings.Trim(rockPosStr, " ")
			coord := strings.Split(rockPosStr, ",")
			col, _ := strconv.Atoi(coord[0])
			row, _ := strconv.Atoi(coord[1])
			rockPos := utils.Position{
				R: row,
				C: col,
			}
			verts = append(verts, rockPos)

			//update bounds
			if bounds == nil {
				bounds = newBound(rockPos)
			} else {
				bounds.update(rockPos)
			}
		}

		for i := 0; i < len(verts)-1; i++ {
			gen := verts[i].StraightTo(verts[i+1])
			cells = append(cells, gen...)
		}
	}

	// assume col sand origin is within bounds

	SolvePart1(bounds, cells)
	SolvePart2(bounds, cells)

}

func SolvePart2(bounds *rectBound, cells []utils.Position) {
	depth := bounds.rows.max + 3
	totCols := (500 + depth) - (500 - depth) + 2
	var grid = newGrid(
		bounds.rows.max+3, // this includes 0 for sand
		totCols,           //add space for all the sand left and right
		-(500 - depth - 1),
	)

	for _, c := range cells {
		grid.set(c, '#')
	}

	lastRow := len(grid.grid) - 1
	for c := 0; c < len(grid.grid[0]); c++ {
		grid.grid[lastRow][c] = '#'
	}

	//grid.print()

	units := 0
	allSteady := false
	for !allSteady {
		//create sand
		sand := utils.Position{
			R: 0,
			C: 500,
		}

		//move sand
		atRest := false

		for !atRest {

			if grid.get(sand.Down()) == '.' {
				sand = sand.Down()
			} else if grid.get(sand.Down().Left()) == '.' {
				sand = sand.Down().Left()
			} else if grid.get(sand.Down().Right()) == '.' {
				sand = sand.Down().Right()
			} else {
				grid.set(sand, 'o')
				atRest = true
				units++
				if sand.R == 0 && sand.C == 500 {
					allSteady = true
					//grid.print()
				}
			}
		}

	}

	fmt.Println("Part 2:", units)

}

func SolvePart1(bounds *rectBound, cells []utils.Position) {
	var grid = newGrid(
		bounds.rows.max+1+1, // this includes 0 for sand
		bounds.cols.size()+2,
		-(bounds.cols.min - 1),
	)

	for _, c := range cells {
		grid.set(c, '#')
	}

	//grid.print()

	units := 0
	allSteady := false
	for !allSteady {
		//create sand
		sand := utils.Position{
			R: 0,
			C: 500,
		}

		//move sand
		atRest := false

		for !atRest {

			if grid.get(sand.Down()) == '.' {
				sand = sand.Down()
			} else if grid.get(sand.Down().Left()) == '.' {
				sand = sand.Down().Left()
			} else if grid.get(sand.Down().Right()) == '.' {
				sand = sand.Down().Right()
			} else {
				grid.set(sand, 'o')
				atRest = true
				units++
				//grid.print()
			}

			if grid.isBottom(sand) {
				atRest = true
				allSteady = true
				//grid.print()
			}

		}

	}

	fmt.Println("Part 1:", units)
}

type translatedGrid struct {
	grid           [][]rune
	colTranslation int
}

func (g *translatedGrid) get(pos utils.Position) rune {
	realCol := pos.C + g.colTranslation

	val := g.grid[pos.R][realCol]
	if val == 0 {
		return '.'
	}
	return val
}

func (g *translatedGrid) set(pos utils.Position, val rune) {
	realCol := pos.C + g.colTranslation
	g.grid[pos.R][realCol] = val
}

func (g *translatedGrid) print() {
	for r := 0; r < len(g.grid); r++ {
		for c := 0; c < len(g.grid[r]); c++ {
			val := g.grid[r][c]
			if val == 0 { //air
				fmt.Print(".")
			} else {
				fmt.Print(string(g.grid[r][c]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *translatedGrid) isBottom(pos utils.Position) bool {
	return pos.R == len(g.grid)-1
}

func newGrid(rows int, cols int, colTranslation int) *translatedGrid {
	grid := make([][]rune, rows)

	for r := 0; r < rows; r++ {
		grid[r] = make([]rune, cols)
	}

	return &translatedGrid{
		grid:           grid,
		colTranslation: colTranslation,
	}
}

type rectBound struct {
	rows axisBound
	cols axisBound
}

type axisBound struct {
	min int
	max int
}

func (a axisBound) size() int {
	return a.max - a.min
}

func newBound(pos utils.Position) *rectBound {
	return &rectBound{
		rows: axisBound{
			min: pos.R,
			max: pos.R,
		},
		cols: axisBound{
			min: pos.C,
			max: pos.C,
		},
	}
}
func (b *rectBound) update(pos utils.Position) {
	if pos.R < b.rows.min {
		b.rows.min = pos.R
	}
	if pos.R > b.rows.max {
		b.rows.max = pos.R
	}
	if pos.C < b.cols.min {
		b.cols.min = pos.C
	}
	if pos.C > b.cols.max {
		b.cols.max = pos.C
	}
}
