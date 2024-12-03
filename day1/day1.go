package day1

import (
	"aoc24/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Task11() {
	fmt.Println("\nDay 1, Part 1: Historian Hysteria")
	fmt.Println("===============================================")
	first, second := prepare("/data/day1/1/input.txt")

	fmt.Println("Result=", first)
	fmt.Println("Result=", second)
	sort.Slice(first, func(i, j int) bool { return first[i] < first[j] })
	sort.Slice(second, func(i, j int) bool { return second[i] < second[j] })
	fmt.Println("Result=", first)
	fmt.Println("Result=", second)
	total := 0
	for idx, firstValue := range first {
		secondValue := second[idx]
		diff := 0
		if firstValue > secondValue {
			diff = firstValue - secondValue
		} else {
			diff = secondValue - firstValue
		}
		total = total + diff
	}
	fmt.Println("Result=", total)

}

func Task12() {
	fmt.Println("\nDay 1, Part 2: Historian Hysteria")
	fmt.Println("===============================================")
	first, second := prepare("/data/day1/2/input.txt")

	fmt.Println("First=", first)
	fmt.Println("Second=", second)

	dict := make(map[int]int)
	for _, num := range second {
		dict[num] = dict[num] + 1
	}
	fmt.Println("SecondDictionary=", dict)
	total := 0
	for _, value := range first {
		total = total + dict[value]*value
	}
	fmt.Println("Result=", total)
}

func prepare(data string) ([]int, []int) {
	lines := util.ReadFile(data)

	first := make([]int, 0)
	second := make([]int, 0)
	var i int
	for _, l := range lines {
		words := strings.Fields(l)
		i, _ = strconv.Atoi(words[0])
		first = append(first, i)

		i, _ = strconv.Atoi(words[1])
		second = append(second, i)

	}
	return first, second
}
