package day2

import (
	"aoc24/util"
	"fmt"
	"math"
)

func Task22() {
	fmt.Println("\nDay 2, Part 2: Red-Nosed Reports")
	fmt.Println("===============================================")
	reports := util.ReadFile("/data/day2/input.txt")
	const WithTolerance = true
	fmt.Println("Result=", analyzeReports(reports, WithTolerance))
}

func Task21() {
	fmt.Println("\nDay 2, Part 1: Red-Nosed Reports")
	fmt.Println("===============================================")
	reports := util.ReadFile("/data/day2/input.txt")
	const WithoutTolerance = false
	fmt.Println("Result=", analyzeReports(reports, WithoutTolerance))

}

func analyzeReports(reports []string, withTolerance bool) int {
	result := 0
	for _, l := range reports {
		levels := util.StringToIntArray(l, " ")

		checkDirection, checkRange := checkLevels(levels)
		if checkDirection == true && checkRange == true {
			result++
		} else {
			if withTolerance {
				for idx, _ := range levels {
					newLevels := util.RemoveIndexFromSlice(levels, idx)
					checkDirection, checkRange = checkLevels(newLevels)
					if checkDirection == true && checkRange == true {
						result++
						break
					}
				}
			}
		}

	}
	return result
}

func checkLevels(level []int) (bool, bool) {
	diffs := make([]int, 0)
	for index, value := range level {
		if index < len(level)-1 {
			diffs = append(diffs, level[index+1]-value)
		}
	}

	isDirection := checkDirection(diffs)
	isRange := checkRange(diffs)

	//fmt.Println("\nlevel=", level)
	//fmt.Println("diffs=", diffs)
	//fmt.Println("checkDirection=", isDirection)
	//fmt.Println("checkRange=", isRange)

	return isDirection, isRange
}

func checkDirection(diffs []int) bool {
	negatives := 0
	positives := 0
	for _, diff := range diffs {
		if math.Signbit(float64(diff)) == true {
			negatives++
		} else {
			positives++
		}
	}
	if negatives != 0 && positives != 0 {
		return false
	}
	return true
}

func checkRange(diffs []int) bool {
	minRange := 1
	maxRange := 3
	safe := true
	for _, diff := range diffs {

		absoluteDiff := int(math.Abs(float64(diff)))

		if absoluteDiff == 0 {
			safe = false
		}

		if (absoluteDiff < minRange) || (absoluteDiff > maxRange) {
			safe = false
		}

	}

	return safe
}
