package main

import (
	"AdventOfCode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	SampleDay()
}

func SampleDay() {

	lines := utils.ReadLines("day13/input.txt")

	var lists = [][]item{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		list := []item{}

		var stack []*[]item
		stack = append(stack, &list)

		var numAcc *int
		for i := 1; i < len(line)-1; i++ {
			c := line[i]
			if c == '[' {
				newItem := []item{}
				*stack[len(stack)-1] = append(*stack[len(stack)-1], item{list: &newItem})

				stack = append(stack, &newItem)

			} else if c == ']' {
				if numAcc != nil {
					*stack[len(stack)-1] = append(*stack[len(stack)-1], item{x: numAcc})
					numAcc = nil
				}
				stack = stack[:len(stack)-1]
			} else {
				//comma or number
				num, err := strconv.Atoi(string(c))
				if err != nil {
					//comma
					if numAcc != nil {
						*stack[len(stack)-1] = append(*stack[len(stack)-1], item{x: numAcc})
						numAcc = nil
					}
					continue
				}
				//num
				if numAcc == nil {
					numAcc = &num
				} else {
					*numAcc = *numAcc*10 + num
				}
			}
		}
		if numAcc != nil {
			*stack[len(stack)-1] = append(*stack[len(stack)-1], item{x: numAcc})
		}

		lists = append(lists, list)

	}

	var correctIndexes []int
	ix := 0
	for i := 0; i < len(lists); i += 2 {
		packet1 := lists[i]
		packet2 := lists[i+1]
		ix++

		//fmt.Printf("== Pair %d ==\n", ix)
		comp := compare(item{list: &packet1}, item{list: &packet2}, 0)
		//fmt.Printf("correct: %t\n", comp < 0)

		if comp < 0 {
			correctIndexes = append(correctIndexes, ix)
		}

	}

	sum := 0
	for _, index := range correctIndexes {
		sum += index
	}
	fmt.Println("Part 1:", sum)

	tmp2 := 2
	tmp6 := 6
	lists = append(lists, []item{{list: &[]item{{x: &tmp2}}}}, []item{{list: &[]item{{x: &tmp6}}}})
	sort.SliceStable(lists, func(i, j int) bool {
		return compare(item{list: &lists[i]}, item{list: &lists[j]}, 0) < 0
	})

	p1, p2 := 0, 0

	for i, list := range lists {
		str := fmt.Sprintf("%v", list)
		if str == "[[2]]" {
			p1 = i + 1
		} else if str == "[[6]]" {
			p2 = i + 1
		}
	}

	fmt.Println("Part 2:", p1*p2)

}

func compare(left, right item, nesting int) int {
	//fmt.Printf("%sCompare %v vs %v\n", strings.Repeat(" ", nesting), left, right)
	if left.x != nil && right.x != nil {
		// both integer
		diff := *left.x - *right.x
		if diff > 0 {
			return 1
		} else if diff < 0 {
			return -1
		}
		return 0

	} else if left.list != nil && right.list != nil {
		// both lists
		for i := 0; i < len(*left.list); i++ {
			if len(*right.list) <= i {
				return 1
			} else {
				comp := compare((*left.list)[i], (*right.list)[i], nesting+1)
				if comp > 0 {
					return 1
				} else if comp < 0 {
					return -1
				}

			}
		}
		if len(*right.list) > len(*left.list) {
			return -1
		}
		return 0
	} else {
		// one is integer
		if left.x != nil {
			return compare(item{list: &[]item{{x: left.x}}}, right, nesting+1)
		} else {
			return compare(left, item{list: &[]item{{x: right.x}}}, nesting+1)
		}
	}
}

type item struct {
	list *[]item
	x    *int
}

func (i item) String() string {

	if i.x != nil {
		return strconv.Itoa(*i.x)
	} else {
		var strs []string
		for _, it := range *i.list {
			strs = append(strs, fmt.Sprintf("%v", it))
		}
		items := strings.Join(strs, ",")
		return "[" + items + "]"
	}
}

//type parser struct {
//	currentlevel int
//	rootList     []interface{}
//
//	parentList  []interface{}
//	currentList []interface{}
//}
//
//type list struct {
//	items []interface{}
//}
//
//func (p *parser) newList() {
//	if p.currentlevel == 0 {
//		p.rootList = []interface{}{}
//	} else {
//
//	}
//	p.currentlevel++
//}

//
//type item interface {
//	size() int
//	less(x item) int
//}
//
//type itemList struct {
//	items []item
//}
//
//
//type emtpyItem struct {
//
//}
//func (i itemList) size() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (i itemList) less(x item) int {
//	//TODO implement me
//	panic("implement me")
//}
//
//type list struct {
//	items []int
//}
