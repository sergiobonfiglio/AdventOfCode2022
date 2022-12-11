package day11

import (
	"AdventOfCode2022/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolveDay() {

	lines := utils.ReadLines("day11/input.txt")

	var monkeys = readMonkeys(lines)
	maxWorry := calcMaxWorry(monkeys)

	for round := 0; round < 20; round++ {
		for _, currMonkey := range monkeys {
			currMonkey.play(monkeys, true, maxWorry)
		}
	}

	monkeyBusiness := calcMonkeyBusiness(monkeys)

	fmt.Println("Part 1:", monkeyBusiness)

	//reset monkeys
	monkeys = readMonkeys(lines)

	for round := 0; round < 10000; round++ {
		for _, currMonkey := range monkeys {
			currMonkey.play(monkeys, false, maxWorry)
		}
	}

	monkeyBusiness = calcMonkeyBusiness(monkeys)
	fmt.Println("Part 2:", monkeyBusiness)

}

const Debug = false

func debug(format string, a ...any) {
	if Debug {
		fmt.Printf(format, a)
	}
}

func calcMaxWorry(monkeys []*monkey) int {
	magicNumber := 1
	for _, m := range monkeys {
		magicNumber *= m.divTest
	}
	return magicNumber
}

func readMonkeys(lines []string) []*monkey {
	var monkeys []*monkey
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			monkeys = append(monkeys, &monkey{
				id:             len(monkeys),
				inspectedItems: 0,
			})
		} else if line != "" {
			currMonkey := monkeys[len(monkeys)-1]
			parseToMonkey(line, currMonkey)
		}
	}
	return monkeys
}

func calcMonkeyBusiness(monkeys []*monkey) int {
	var activity []int
	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		activity = append(activity, m.inspectedItems)
		debug("Monkey %d inspected items %d times.\n", i, m.inspectedItems)
	}

	sort.Ints(activity)
	monkeyBusiness := activity[len(activity)-1] * activity[len(activity)-2]
	return monkeyBusiness
}

func parseToMonkey(line string, currMonkey *monkey) {
	if strings.HasPrefix(line, "  Starting items: ") {
		data := line[len("  Starting items: "):]
		split := strings.Split(data, ", ")
		var items []int
		for _, x := range split {
			n, _ := strconv.Atoi(x)
			items = append(items, n)
		}
		currMonkey.items = items
	} else if strings.HasPrefix(line, "  Operation: ") {
		data := line[len("  Operation: new = "):]
		split := strings.Split(data, " ")

		currMonkey.operation = buildOperation(split[0], split[1], split[2])

	} else if strings.HasPrefix(line, "  Test") {
		data := line[len("  Test: divisible by "):]
		arg, _ := strconv.Atoi(data)
		currMonkey.test = func(x int) bool {
			return x%arg == 0
		}
		currMonkey.divTest = arg

	} else if strings.HasPrefix(line, "    If true") {
		data := line[len("    If true: throw to monkey "):]
		x, _ := strconv.Atoi(data)
		currMonkey.throwTrue = x

	} else if strings.HasPrefix(line, "    If false") {
		data := line[len("    If false: throw to monkey "):]
		x, _ := strconv.Atoi(data)
		currMonkey.throwFalse = x
	}
}

func buildOperation(x string, operation string, y string) func(old int) int {

	return func(old int) int {

		op1 := getOperand(x, old)
		op2 := getOperand(y, old)

		switch operation {
		case "+":
			return op1 + op2
		case "*":
			return op1 * op2

		default:
			panic("unknown operation " + operation)
		}
	}
}
func getOperand(x string, old int) int {
	var op1 int
	if x == "old" {
		op1 = old
	} else {
		op1, _ = strconv.Atoi(x)
	}
	return op1
}

type monkey struct {
	id         int
	items      []int
	operation  func(old int) int
	test       func(x int) bool
	throwTrue  int
	throwFalse int

	divTest        int
	inspectedItems int
}

func (m *monkey) play(monkeys []*monkey, divWorryLevel bool, maxWorry int) {

	debug("Monkey %d\n", m.id)
	for _, item := range m.items {
		m.inspectedItems++
		debug("\tMonkey inspects an item with a worry level of %d.\n", item)

		wl := m.operation(item)
		debug("\t\tWorry level is multiplied by ? to %d.\n", wl)

		if divWorryLevel {
			wl = int(math.Floor(float64(wl) / 3))
			debug("\t\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n", wl)
		}
		wl = wl % maxWorry

		testRes := m.test(wl)
		debug("\t\tCurrent worry level pass test: %t\n", testRes)

		var target int
		if testRes {
			target = m.throwTrue
		} else {
			target = m.throwFalse
		}

		m.items = m.items[1:]
		monkeys[target].items = append(monkeys[target].items, wl)
		debug("\t\tItem with worry level %d is thrown to monkey %d.\n", wl, target)
	}

}
