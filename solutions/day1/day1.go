package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"math"
)

func main() {
	lines, err := lib.ReadLinesToIntList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day1/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	prevSum := math.MaxInt
	for i := 0; i < len(lines) - 2; i++ {
		sum := lines[i] + lines[i + 1] + lines[i + 2]

		if sum > prevSum {
			count++
		}

		prevSum = sum
	}

	fmt.Println(count)
}
