package day3

import (
	"AdventOfCode2022/utils"
	"fmt"
)

func getCommon(comp1 string, comp2 string) []byte {
	var commons []byte
	var inComp2 = map[byte]bool{}
	for _, l := range comp1 {
		c := byte(l)

		if ok := inComp2[c]; !ok {
			inComp2[c] = false
			for j := range comp2 {
				if c == comp2[j] {
					commons = append(commons, c)
					inComp2[c] = true
					break
				}
			}
		}
	}
	return commons
}

func getPriority(letter byte) byte {
	if letter <= 'z' && letter >= 'a' {
		return letter - 'a' + 1
	} else {
		return letter - 'A' + 27
	}
}

func SolveDay() {

	var rucksacks = utils.ReadLines("day3/input.txt")

	sum := 0
	for _, rucksack := range rucksacks {

		half := len(rucksack) / 2
		comp1 := rucksack[:half]
		comp2 := rucksack[half:]

		commons := getCommon(comp1, comp2)
		if len(commons) != 1 {
			panic("unexpected number of commons")
		}
		priority := getPriority(commons[0])
		sum += int(priority)
	}

	fmt.Println("Part 1:", sum)

	sum = 0
	for i := 0; i < len(rucksacks); i += 3 {

		rs1, rs2, rs3 := rucksacks[i], rucksacks[i+1], rucksacks[i+2]

		comm12 := getCommon(rs1, rs2)
		commAll := getCommon(string(comm12), rs3)
		if len(commAll) != 1 {
			panic("unexpected number of commons")
		}
		priority := getPriority(commAll[0])
		sum += int(priority)
	}
	fmt.Println("Part 2:", sum)

}
