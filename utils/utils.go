package utils

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
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

type Coord struct {
	X int
	Y int
}

func NewCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func (p Coord) DistManhattan(p2 Coord) int {
	return int(math.Abs(float64(p.X-p2.X)) + math.Abs(float64(p.Y-p2.Y)))
}

type IntScanner struct {
	currentNumb     *int
	currentNegative bool
	Numbers         []int
}

func NewIntScanner() *IntScanner {
	return &IntScanner{
		currentNumb:     nil,
		currentNegative: false,
	}
}

func (s *IntScanner) ReadLine(line string) {

	lineReader := strings.NewReader(line)
	c, _, err := lineReader.ReadRune()
	for err == nil {
		if c == '-' {
			s.startInt()
			s.setNegative()
		} else if c >= '0' && c <= '9' {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				panic("error parsing rune " + string(c))
			}
			if s.currentNumb == nil {
				s.startInt()
			}
			s.setDigit(digit)
		} else {
			s.startInt()
		}

		c, _, err = lineReader.ReadRune()
	}
	s.startInt()
}

func (s *IntScanner) startInt() {
	if s.currentNumb != nil {
		if s.currentNegative {
			*s.currentNumb = *s.currentNumb * -1
		}
		s.Numbers = append(s.Numbers, *s.currentNumb)
		s.currentNumb = nil
		s.currentNegative = false
	}
}

func (s *IntScanner) setDigit(digit int) {
	if s.currentNumb == nil {
		s.startInt()
		s.currentNumb = &digit
	} else {
		*s.currentNumb = *s.currentNumb*10 + digit
	}
}

func (s *IntScanner) setNegative() {
	s.currentNegative = true
}
