package day4

import (
	"aoc24/util"
	"fmt"
	"strings"
)

func Task42() {
	fmt.Println("\nDay 4, Part 2: Ceres Search")
	fmt.Println("===============================================")
	data := util.ReadFileAsGrid("/data/day4/input.txt")
	result := 0

	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			val := data[i+1][j+1] + data[i-1][j-1] + data[i+1][j-1] + data[i-1][j+1]
			if data[i][j] == "A" && strings.Contains("MSMS;MSSM;SMMS;SMSM", val) {
				result += 1
			}
		}
	}
	fmt.Println("Result=", result)
}

func Task41() {
	fmt.Println("\nDay 4, Part 1: Ceres Search")
	fmt.Println("===============================================")
	data := util.ReadFileAsGrid("/data/day4/input.txt")

	result := 0
	for y, _ := range data {
		for x, _ := range data[y] {
			// rechts
			if len(data[y]) >= x+4 && "XMAS" == data[y][x]+data[y][x+1]+data[y][x+2]+data[y][x+3] {
				result += 1
			}
			// rechts unten
			if len(data) >= y+4 && len(data[y]) >= x+4 && "XMAS" == data[y][x]+data[y+1][x+1]+data[y+2][x+2]+data[y+3][x+3] {
				result += 1
			}
			// unten
			if len(data) >= y+4 && "XMAS" == data[y][x]+data[y+1][x]+data[y+2][x]+data[y+3][x] {
				result += 1
			}
			// links unten
			if len(data) >= y+4 && x-4 >= -1 && "XMAS" == data[y][x]+data[y+1][x-1]+data[y+2][x-2]+data[y+3][x-3] {
				result += 1
			}
			// links
			if x-4 >= -1 && "XMAS" == data[y][x]+data[y][x-1]+data[y][x-2]+data[y][x-3] {
				result += 1
			}
			// links oben
			if y-4 >= -1 && x-3 >= 0 && "XMAS" == data[y][x]+data[y-1][x-1]+data[y-2][x-2]+data[y-3][x-3] {
				result += 1
			}
			// oben
			if y-4 >= -1 && "XMAS" == data[y][x]+data[y-1][x]+data[y-2][x]+data[y-3][x] {
				result += 1
			}
			// rechts oben
			if y-4 >= -1 && len(data[y]) >= x+4 && "XMAS" == data[y][x]+data[y-1][x+1]+data[y-2][x+2]+data[y-3][x+3] {
				result += 1
			}
		}
	}

	fmt.Println("Result=", result)

}
