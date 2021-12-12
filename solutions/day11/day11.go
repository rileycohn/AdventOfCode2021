package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

var grid = [10][10]int{}
var alreadyFlashed = [10][10]bool{}

func main() {
	fmt.Println("Day 11!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day11/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, line := range lines {
		// fmt.Println(line)
		for j, c := range line {
			val, _ := strconv.ParseInt(string(c), 10, 0)
			grid[i][j] = int(val)
		}
	}

	printGridPretty(grid)

	// 100 Steps
	totalFlashes := 0
	for i :=0; i < 10000; i++ {
		alreadyFlashed = [10][10]bool{}
		for row:= 0; row < len(grid); row++ {
			for col, _ := range grid[row] {
				grid[row][col] = grid[row][col] + 1

				if grid[row][col] == 10 && !alreadyFlashed[row][col] {
					grid[row][col] = 0
					alreadyFlashed[row][col] = true
					flash(row, col)
				}
			}
		}

		// Check how many trues already flashes has
		numFlashes := 0
		for r, flashes := range alreadyFlashed {
			for c, _ := range flashes {
				if alreadyFlashed[r][c] {
					numFlashes++
					grid[r][c] = 0
				}
			}
		}

		totalFlashes = totalFlashes + numFlashes

		if numFlashes == 100 {
			fmt.Println(i)
			break
		}
		fmt.Println(numFlashes)
		printGridPretty(grid)
	}

	fmt.Println(totalFlashes)
}

func flash(row int, col int){
	// Go through each row and col and increase each cell by 1
	nearby := []string{"-1,1", "0,1", "1,1", "1,0", "1,-1", "0,-1", "-1,-1", "-1,0"}
	for _, pair := range nearby {
		left,_ := strconv.ParseInt(strings.Split(pair, ",")[0], 10, 0)
		right,_ := strconv.ParseInt(strings.Split(pair, ",")[1], 10, 0)

		newLeft := row + int(left)
		newRight := col + int(right)

		if newLeft < 0 || newLeft > 9 || newRight < 0 || newRight > 9 {
			continue
		}

		grid[newLeft][newRight] = grid[newLeft][newRight] + 1

		if grid[newLeft][newRight] == 10 && !alreadyFlashed[newLeft][newRight] {
			grid[newLeft][newRight] = 0
			alreadyFlashed[newLeft][newRight] = true
			flash(newLeft, newRight)
		}
	}
}

func printGridPretty(grid [10][10]int) {
	for i := 0; i< 10; i++ {
		for j :=0; j < 10; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println("----------")
}