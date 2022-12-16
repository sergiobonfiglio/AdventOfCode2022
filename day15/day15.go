package main

import (
	"AdventOfCode2022/utils"
	"fmt"
)

func main() {

	SolveDay()
}

func SolveDay() {

	lines := utils.ReadLines("day15/input.txt")

	var readings []reading
	var minX, maxX *int
	for _, line := range lines {
		numScanner := utils.NewIntScanner()
		numScanner.ReadLine(line)

		data := newReading(numScanner.Numbers...)
		readings = append(readings, data)

		left := data.sensor.X - data.radius
		if minX == nil || *minX > left {
			tmp := left
			minX = &tmp
		}
		right := data.sensor.X + data.radius
		if maxX == nil || *maxX < right {
			tmp := right
			maxX = &tmp
		}
	}

	count := 0
	for x := *minX; x < *maxX; x++ {
		point := utils.Coord{
			X: x,
			Y: 2000000,
		}

		if inRange := isInAnyRange(readings, point, false); inRange {
			count++
		}

	}

	fmt.Println("Part 1:", count)

	maxSearchSpace := 4000000
	tuningFreq := -1

	searchRect := rectangle{
		tl: utils.NewCoord(0, 0),
		tr: utils.NewCoord(maxSearchSpace, 0),
		bl: utils.NewCoord(0, maxSearchSpace),
		br: utils.NewCoord(maxSearchSpace, maxSearchSpace),
	}
	outOfRange := searchOutOfRange(readings, searchRect)

	tuningFreq = outOfRange.X*4000000 + outOfRange.Y

	fmt.Println("Part 2:", tuningFreq)
}

func isInAnyRange(readings []reading, point utils.Coord, includeExisting bool) bool {
	for _, r := range readings {
		if isPointInRange(point, r, includeExisting) {
			return true
		}
	}
	return false
}

func isPointInRange(point utils.Coord, r reading, includeExisting bool) bool {
	if point == r.sensor || point == r.nearestBeacon {
		return includeExisting
	}
	dist := point.DistManhattan(r.sensor)
	if dist <= r.radius {
		return true
	}
	return false
}

type rectangle struct {
	tl, tr, bl, br utils.Coord
}

func (r rectangle) isValid() bool {
	return r.tl.X <= r.br.X && r.tl.Y <= r.br.Y
}

func (r rectangle) isPoint() bool {
	return r.tl == r.tr &&
		r.tl == r.bl &&
		r.tl == r.br
}

func (r rectangle) subDivide() []rectangle {
	midpoint := utils.Coord{
		X: r.tl.X + (r.tr.X-r.tl.X)/2,
		Y: r.tl.Y + (r.bl.Y-r.tl.Y)/2,
	}

	tlRect := rectangle{ //tl
		tl: r.tl,
		tr: utils.NewCoord(midpoint.X, r.tl.Y),
		bl: utils.NewCoord(r.tl.X, midpoint.Y),
		br: midpoint,
	}
	trRect := rectangle{ //tr
		tl: utils.NewCoord(midpoint.X+1, r.tl.Y),
		tr: r.tr,
		bl: utils.NewCoord(midpoint.X+1, midpoint.Y),
		br: utils.NewCoord(r.tr.X, midpoint.Y),
	}
	blRect := rectangle{ //bl
		tl: utils.NewCoord(r.tl.X, midpoint.Y+1),
		tr: utils.NewCoord(midpoint.X, midpoint.Y+1),
		bl: r.bl,
		br: utils.NewCoord(midpoint.X, r.bl.Y),
	}
	brRect := rectangle{ //br
		tl: utils.NewCoord(midpoint.X+1, midpoint.Y+1),
		tr: utils.NewCoord(r.tr.X, midpoint.Y+1),
		bl: utils.NewCoord(midpoint.X+1, r.bl.Y),
		br: r.br,
	}

	var resp []rectangle
	for _, r := range []rectangle{tlRect, trRect, blRect, brRect} {
		if r.isValid() {
			resp = append(resp, r)
		}
	}
	return resp
}

func isSquareInRange(readings reading, square rectangle) bool {
	return isPointInRange(square.bl, readings, true) &&
		isPointInRange(square.br, readings, true) &&
		isPointInRange(square.tl, readings, true) &&
		isPointInRange(square.tr, readings, true)
}

func isRectInAnyRange(readings []reading, square rectangle) bool {
	for _, r := range readings {
		if isSquareInRange(r, square) {
			return true
		}
	}
	return false
}

func searchOutOfRange(readings []reading, rect rectangle) *utils.Coord {
	if rect.isPoint() {
		if !isInAnyRange(readings, rect.br, true) {
			return &rect.br
		}
		return nil
	}

	if isRectInAnyRange(readings, rect) {
		return nil
	}

	subRects := rect.subDivide()

	for _, subRect := range subRects {
		outOfRangePoint := searchOutOfRange(readings, subRect)
		if outOfRangePoint != nil {
			return outOfRangePoint
		}
	}
	return nil
}

type reading struct {
	sensor        utils.Coord
	nearestBeacon utils.Coord
	radius        int
}

func newReading(coords ...int) reading {
	tmp := reading{
		sensor: utils.Coord{
			X: coords[0],
			Y: coords[1],
		},
		nearestBeacon: utils.Coord{
			X: coords[2],
			Y: coords[3],
		},
	}
	tmp.radius = tmp.sensor.DistManhattan(tmp.nearestBeacon)
	return tmp
}
