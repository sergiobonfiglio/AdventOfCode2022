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
	head position
	tail position
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

func (b *board) updateTail() position {
	if areTouching(b.head, b.tail) {
		return b.tail
	} else {

		//not touching, need to move tail

		var candidates []position
		if b.head.x != b.tail.x && b.head.y != b.tail.y {
			//needs to move diagonally first

			candidates = []position{
				b.tail.up().left(),
				b.tail.up().right(),
				b.tail.down().left(),
				b.tail.down().right(),
			}

		} else {
			candidates = []position{
				b.tail.left(),
				b.tail.right(),
				b.tail.up(),
				b.tail.down(),
			}
		}
		for _, candidate := range candidates {
			if areTouching(candidate, b.head) {
				b.tail = candidate
				return candidate
			}
		}
		panic("unable to find head")
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

func newBoard() board {
	start := position{0, 0}
	return board{
		head: start,
		tail: start,
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

	board := newBoard()
	tailVisited := map[position]bool{}
	tailVisited[board.tail] = true
	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			board.moveHead(move.dir)
			pos := board.updateTail()
			tailVisited[pos] = true
		}
	}

	fmt.Println("Part 1:", len(tailVisited))
}
