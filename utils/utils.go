package utils

import (
	"bufio"
	"os"
)

func ReadLines(inputPath string) []string {
	readFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

type Position struct {
	R int
	C int
}

func (p Position) Left() Position {
	return Position{
		R: p.R,
		C: p.C - 1,
	}
}
func (p Position) Right() Position {
	return Position{
		R: p.R,
		C: p.C + 1,
	}
}
func (p Position) Up() Position {
	return Position{
		R: p.R - 1,
		C: p.C,
	}
}
func (p Position) Down() Position {
	return Position{
		R: p.R + 1,
		C: p.C,
	}
}

func (p Position) IsValid(w int, h int) bool {
	return p.R >= 0 && p.R < w &&
		p.C >= 0 && p.C < h
}

func (p Position) StraightTo(dest Position) []Position {
	var cells []Position
	cells = append(cells, p)
	//down
	for ri := p.R + 1; ri <= dest.R; ri++ {
		cells = append(cells, Position{
			R: ri,
			C: p.C,
		})
	}

	//up
	for ri := p.R - 1; ri >= dest.R; ri-- {
		cells = append(cells, Position{
			R: ri,
			C: p.C,
		})
	}

	//right
	for ci := p.C + 1; ci <= dest.C; ci++ {
		cells = append(cells, Position{
			R: p.R,
			C: ci,
		})
	}

	//left
	for ci := p.C - 1; ci >= dest.C; ci-- {
		cells = append(cells, Position{
			R: p.R,
			C: ci,
		})
	}

	return cells
}
