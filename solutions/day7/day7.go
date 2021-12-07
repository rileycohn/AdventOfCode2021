package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Day 7!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day7/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	vals := strings.Split(lines[0], ",")

	// Convert to ints
	crabPositions := make([]int64, len(vals))
	for i, v := range vals {
		crabPos, _ := strconv.ParseInt(v, 10, 0)
		crabPositions[i] = crabPos
	}

	// Find the mean
	var sum float64 = 0
	for _, crab := range crabPositions {
		sum += float64(crab)
	}

	mean := int64(sum / float64(len(crabPositions)))

	fuelUsed := 0
	for _, crab := range crabPositions {
		diff := int(math.Abs(float64(crab - mean)))
		fuelUsed += int(math.Pow(float64(diff), 2) + float64(diff)) / 2
	}

	fmt.Println(fuelUsed)
}
