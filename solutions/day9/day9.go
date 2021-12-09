package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 9!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day9/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := [100][100]int{}
	answer := 0
	for i, line := range lines {
		splitLine := strings.Split(line, "")
		for j, val := range splitLine {
			valInt,_ := strconv.ParseInt(val, 10, 0)
			grid[i][j] = int(valInt)
		}
	}

	// Now we have the grid, find the low points
	// lowPoints := [][]int{}
	largestBasins := [3]int{}
	for i, h := range grid {
		for j, cell := range h {
			left :=  j - 1
			right := j + 1
			up := i - 1
			down := i + 1

			if !(left >= 0 && cell >= grid[i][left]) &&
				!(right < len(h) && cell >= grid[i][right]) &&
				!(up >= 0 && cell >= grid[up][j]) &&
				!(down < len(grid) && cell >= grid[down][j]) {
				fmt.Println("Found low point")
				fmt.Println(i)
				fmt.Println(j)
				fmt.Println(grid[i][j])
				answer += 1 + grid[i][j]
			}
		}
		fmt.Println()
	}

	fmt.Println(answer)
}