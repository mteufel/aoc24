package day3

import (
	"aoc24/util"
	"fmt"
	"regexp"
	"strconv"
)

func Task32() {
	fmt.Println("\nDay 3, Part 2: Mull It Over")
	fmt.Println("===============================================")
	input := util.ReadFileAsString("/day3/data/2/input.txt")

	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	matches := r.FindAllStringSubmatch(input, -1)
	result := 0
	on := true
	for _, v := range matches {
		if v[0] == "don't()" {
			on = false
		}
		if v[0] == "do()" {
			on = true
		}

		if on {
			x, _ := strconv.Atoi(v[1])
			y, _ := strconv.Atoi(v[2])
			result = result + x*y
		}
	}

	fmt.Println("Result=", result)
}

func Task31() {
	fmt.Println("\nDay 3, Part 1: Mull It Over")
	fmt.Println("===============================================")
	input := util.ReadFileAsString("/day3/data/1/input.txt")

	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	result := 0
	for _, v := range matches {
		x, _ := strconv.Atoi(v[1])
		y, _ := strconv.Atoi(v[2])
		result = result + x*y
	}

	fmt.Println("Result=", result)
}
