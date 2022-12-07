package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Play = string

const (
	playRock    Play = "A"
	playPaper   Play = "B"
	playScissor Play = "C"
)

type Response = string

const (
	respRock    Response = "X"
	respPaper   Response = "Y"
	respScissor Response = "Z"
)

type outcome = int

const (
	win  outcome = iota
	draw outcome = iota
	loss outcome = iota
)

func responseToOutcome(r Response) outcome {
	switch r {
	case respRock:
		return loss
	case respPaper:
		return draw
	case respScissor:
		return win
	default:
		panic("cannot convert response to play")
	}
}

func responseToPlay(r Response) Play {
	switch r {
	case respRock:
		return playRock
	case respPaper:
		return playPaper
	case respScissor:
		return playScissor
	default:
		panic("cannot convert response to play")
	}
}

type strategy struct {
	opponent Play
	response string
}

func readInput() []strategy {
	readFile, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var data []strategy
	for fileScanner.Scan() {
		var line = fileScanner.Text()

		split := strings.Split(line, " ")
		if len(split) != 2 {
			panic("input parse error")
		}

		data = append(data, strategy{
			opponent: split[0],
			response: split[1],
		})
	}

	return data
}

func (s strategy) outcome() outcome {
	var respPlay = responseToPlay(s.response)
	return outcomeLUT[s.opponent][respPlay]
}

func (s strategy) score() int {
	return calcScore(responseToPlay(s.response), s.outcome())
}

var (
	outcomeLUT = map[string]map[string]outcome{
		playRock: {
			playRock:    draw,
			playPaper:   win,
			playScissor: loss,
		},
		playPaper: {
			playRock:    loss,
			playPaper:   draw,
			playScissor: win,
		},
		playScissor: {
			playRock:    win,
			playPaper:   loss,
			playScissor: draw,
		},
	}
)

func buildInverseLUT() map[string]map[outcome]string {
	var invLUT = map[string]map[outcome]string{}
	for play, options := range outcomeLUT {
		invLUT[play] = map[outcome]string{}
		for resp, out := range options {
			invLUT[play][out] = resp
		}
	}
	return invLUT
}

var inverseLUT = buildInverseLUT()

func (s strategy) score2() int {

	wantedOutcome := responseToOutcome(s.response)

	respForOutcome := inverseLUT[s.opponent][wantedOutcome]

	return calcScore(respForOutcome, wantedOutcome)
}

func calcScore(choice Play, result outcome) int {
	var score = 0
	if choice == playRock {
		score += 1
	} else if choice == playPaper {
		score += 2
	} else if choice == playScissor {
		score += 3
	}

	if result == win {
		score += 6
	} else if result == draw {
		score += 3
	}

	return score
}

func SolveDay() {

	data := readInput()

	tot := 0
	for _, strat := range data {
		tot += strat.score()
	}

	fmt.Println("Part 1:", tot)

	tot = 0
	for _, strat := range data {
		tot += strat.score2()
	}

	fmt.Println("Part 2:", tot)
}
