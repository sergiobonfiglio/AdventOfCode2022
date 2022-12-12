package main

import (
	"AdventOfCode2022/utils"
	"fmt"
)

func main() {

	SolveDay()
}

func SolveDay() {
	lines := utils.ReadLines("day12/input.txt")

	grid := make([][]rune, len(lines))
	hCounts := map[rune]int{}

	var start position
	var end position
	for r, line := range lines {
		grid[r] = make([]rune, len(line))
		for c, el := range line {
			height := el
			if el == 'S' {
				start = position{
					r: r,
					c: c,
				}
				height = 'a'
			} else if el == 'E' {
				end = position{
					r: r,
					c: c,
				}
				height = 'z'
			}
			hCounts[height]++
			grid[r][c] = height
		}

	}

	graph := newDG()
	var starts []position
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {

			p := position{
				r: r,
				c: c,
			}
			pHeight := grid[p.r][p.c]

			if pHeight == 'a' {
				starts = append(starts, p)
			}

			for _, n := range []position{p.up(), p.down(), p.left(), p.right()} {
				if n.isValid(len(grid), len(grid[0])) {
					nHeight := grid[n.r][n.c]
					if nHeight-pHeight <= 1 {
						graph.addEdge(p, n)
					}
				}
			}

		}
	}

	steps := graph.bfs(start, end, map[position]bool{})
	fmt.Println("Part 1:", steps)

	minStart := steps
	for _, pos := range starts {
		altSteps := graph.bfs(pos, end, map[position]bool{})
		if altSteps != -1 && altSteps < minStart {
			minStart = altSteps
		}
	}
	fmt.Println("Part 2:", minStart)

}

type directedGraph struct {
	weights map[position]map[position]int
}

type queueItem struct {
	pos   position
	depth int
}

func (g *directedGraph) bfs(start position, end position, visited map[position]bool) int {

	queue := []queueItem{{pos: start, depth: 0}}
	for len(queue) > 0 {

		current := queue[0]

		visited[current.pos] = true

		queue = queue[1:]

		if current.pos == end {
			return current.depth
		} else {

			for neighbour := range g.weights[current.pos] {
				if !visited[neighbour] {
					visited[neighbour] = true
					queue = append(queue,
						queueItem{
							pos:   neighbour,
							depth: current.depth + 1,
						})
				}
			}
		}

	}

	return -1
}

func newDG() *directedGraph {
	return &directedGraph{
		weights: map[position]map[position]int{},
	}
}

func (g *directedGraph) addEdge(p1, p2 position) {
	if g.weights[p1] == nil {
		g.weights[p1] = map[position]int{}
	}
	g.weights[p1][p2] = 1
}

type position struct {
	r int
	c int
}

func (p position) left() position {
	return position{
		r: p.r,
		c: p.c - 1,
	}
}
func (p position) right() position {
	return position{
		r: p.r,
		c: p.c + 1,
	}
}
func (p position) up() position {
	return position{
		r: p.r - 1,
		c: p.c,
	}
}
func (p position) down() position {
	return position{
		r: p.r + 1,
		c: p.c,
	}
}

func (p position) isValid(w int, h int) bool {
	return p.r >= 0 && p.r < w &&
		p.c >= 0 && p.c < h
}
