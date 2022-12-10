package day9

import (
	"AdventOfCode2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type move struct {
	dir    string
	amount int
}

type position struct {
	x int
	y int
}

type board struct {
	head  position
	knots []position
}

func (b *board) tail() position {
	return b.knots[len(b.knots)-1]
}

func (b *board) moveHead(dir string) {
	switch dir {
	case "R":
		b.head.x++
	case "L":
		b.head.x--
	case "U":
		b.head.y++
	case "D":
		b.head.y--
	}
}

func (b *board) updateKnots() {
	for i, knot := range b.knots {
		var parent position
		if i == 0 {
			parent = b.head
		} else {
			parent = b.knots[i-1]
		}

		if !areTouching(parent, knot) {
			//not touching, need to move tail
			var candidates []position
			found := false
			if parent.x != knot.x && parent.y != knot.y {
				//needs to move diagonally first

				candidates = []position{
					knot.up().left(),
					knot.up().right(),
					knot.down().left(),
					knot.down().right(),
				}

			} else {
				candidates = []position{
					knot.left(),
					knot.right(),
					knot.up(),
					knot.down(),
				}
			}
			for _, candidate := range candidates {
				if areTouching(candidate, parent) {
					b.knots[i] = candidate
					found = true
					break
				}
			}
			if !found {
				panic("unable to find head")
			}
		}
	}

}

func (p position) left() position {
	return position{
		x: p.x - 1,
		y: p.y,
	}
}
func (p position) up() position {
	return position{
		x: p.x,
		y: p.y + 1,
	}
}
func (p position) right() position {
	return position{
		x: p.x + 1,
		y: p.y,
	}
}
func (p position) down() position {
	return position{
		x: p.x,
		y: p.y - 1,
	}
}

func areTouching(pos1, pos2 position) bool {
	return math.Abs(float64(pos1.x-pos2.x)) <= 1 &&
		math.Abs(float64(pos1.y-pos2.y)) <= 1
}

func newBoard(numKnots int) board {
	start := position{0, 0}
	var knots []position
	for i := 0; i < numKnots-1; i++ {
		knots = append(knots, start)
	}
	return board{
		head:  start,
		knots: knots,
	}
}
func SolveDay() {

	lines := utils.ReadLines("day9/input.txt")

	var moves []move
	for _, line := range lines {
		split := strings.Split(line, " ")

		amount, _ := strconv.Atoi(split[1])

		moves = append(moves, move{
			dir:    split[0],
			amount: amount,
		})
	}

	board := newBoard(2)
	tailVisited := map[position]bool{}
	tailVisited[board.tail()] = true
	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			board.moveHead(move.dir)
			board.updateKnots()
			tailVisited[board.tail()] = true
		}
	}

	fmt.Println("Part 1:", len(tailVisited))

	board = newBoard(10)
	tailVisited = map[position]bool{}
	tailVisited[board.tail()] = true
	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			board.moveHead(move.dir)
			board.updateKnots()
			tailVisited[board.tail()] = true
		}
	}

	fmt.Println("Part 2:", len(tailVisited))

}
