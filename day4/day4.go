package day4

import (
	"aoc24/util"
	"bytes"
	"fmt"
)

func Task41() {
	fmt.Println("\nDay 4, Part 1: Ceres Search")
	fmt.Println("===============================================")
	data := util.ReadFile("/data/day4/sample.txt")

	//idea:
	//search for "X" "M" "A" "S" and remember indicies
	// --> search M, A, S in 8 directions --> if match --> count
	// index: left (+1) , right ( -1 ), down (+length), up (-length), down_left (+length+1),  down_right (-length-1)

	for _, line := range data {

		arr := make([]string, 1)
		arr[0] = "X"
		res := findAllOccurrences([]byte(line), arr)
		fmt.Println(res)

	}

}

func findAllOccurrences(data []byte, searches []string) map[string][]int {
	results := make(map[string][]int, 0)

	for _, search := range searches {
		index := len(data)
		tmp := data
		for true {
			match := bytes.LastIndex(tmp[0:index], []byte(search))
			if match == -1 {
				break
			} else {
				index = match
				results[search] = append(results[search], match)
			}
		}
	}

	return results
}
