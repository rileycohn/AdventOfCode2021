package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 13!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day13/day13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// find max x and y
	maxX := 0
	maxY := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		if !strings.Contains(line, "fold") {
			coord := strings.Split(line, ",")
			x,_ := strconv.ParseInt(coord[0], 10, 0)
			y,_ := strconv.ParseInt(coord[1], 10, 0)
			if int(x) > maxX {
				maxX = int(x)
			}

			if int(y) > maxY {
				maxY = int(y)
			}
		}
	}

	coords := make([][]bool, maxX+1)
	for i := range coords {
		coords[i] = make([]bool, maxY+1)
	}

	folds := []string{}

	for _, line := range lines {
		fmt.Println(line)
		if line == "" {
			continue
		}
		if !strings.Contains(line, "fold") {
			coord := strings.Split(line, ",")
			x,_ := strconv.ParseInt(coord[0], 10, 0)
			y,_ := strconv.ParseInt(coord[1], 10, 0)
			coords[x][y] = true
		} else {
			folds = append(folds, line)
		}
	}

	// Handle fold
	for _, fold := range folds {
		axis := strings.Split(fold, "=")
		axisInt,_ := strconv.ParseInt(axis[1], 10, 0)
		if strings.Contains(fold, "x") {
			coords = foldX(coords, int(axisInt))
		} else {
			coords = foldY(coords, int(axisInt))
		}
	}

	// Count the trues
	count := 0
	maxXTrue := 0
	maxYTrue := 0
	for i := 0; i < len(coords); i++ {
		for j :=0; j < len(coords[i]); j++ {
			if coords[i][j] {
				if i > maxXTrue {
					maxXTrue = i
				}

				if j > maxYTrue {
					maxYTrue = j
				}
				count++
			}

		}
	}

	for i := 0; i <= maxYTrue; i++ {
		for j := 0; j <= maxXTrue; j++ {
			if coords[j][i] {
				fmt.Print("##")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
}

func foldY(grid [][]bool, yAxis int) [][]bool {
	newGrid := make([][]bool, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j :=0; j < len(grid[i]); j++ {

			// Top half is unchanged
			if j < yAxis {
				newGrid[i][j] = grid[i][j]
			} else {
				// Bottom half needs to be folded up
				newY := yAxis - (j - yAxis)
				if grid[i][j] {
					newGrid[i][newY] = grid[i][j]
				}
			}
		}
	}
	return newGrid
}

func foldX(grid [][]bool, xAxis int) [][]bool {
	newGrid := make([][]bool, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j :=0; j < len(grid[i]); j++ {

			// Top half is unchanged
			if i < xAxis {
				newGrid[i][j] = grid[i][j]
			} else {
				// Bottom half needs to be folded up
				newX := xAxis - (i - xAxis)
				if grid[i][j] {
					newGrid[newX][j] = grid[i][j]
				}
			}
		}
	}
	return newGrid
}