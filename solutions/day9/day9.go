package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"sort"
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
	lowPoints := []string{}
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
				lowPoints = append(lowPoints, strconv.Itoa(j) + "," + strconv.Itoa(i))
				answer += 1 + grid[i][j]
			}
		}
	}

	// We have all the low points, find basins
	largestBasins := []int{}
	for _, lowPoint := range lowPoints {
		visited := [100][100]bool{}
		visited = getNeighboringNonNines(grid, lowPoint, visited)

		// Check how many trues are in visited
		trues := 0
		for i, h := range visited {
			for j, _ := range h {
				if visited[i][j] {
					trues++
				}
			}
		}

		largestBasins = append(largestBasins, trues)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(largestBasins)))

	result := 1
	for num := 0; num < 3; num++ {
		result *= largestBasins[num]
	}

	fmt.Println(result)
}

func getNeighboringNonNines(grid [100][100]int, point string, visited [100][100]bool) [100][100]bool {
	xy := strings.Split(point, ",")
	x,_ := strconv.Atoi(xy[0])
	y,_ := strconv.Atoi(xy[1])

	visited[y][x] = true

	// Go left
	for i := x - 1; i >= 0; i-- {
		if grid[y][i] != 9 && !visited[y][i] {
			visited[y][i] = true
			visited = getNeighboringNonNines(grid, strconv.Itoa(i) +"," +strconv.Itoa(y), visited)
		} else {
			break
		}
	}

	// Go right
	for i := x + 1; i < len(grid[0]); i++ {
		if grid[y][i] != 9 && !visited[y][i] {
			visited[y][i] = true
			visited = getNeighboringNonNines(grid, strconv.Itoa(i) +"," +strconv.Itoa(y), visited)
		} else {
			break
		}
	}

	// Go up
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] != 9 && !visited[i][x] {
			visited[i][x] = true
			visited = getNeighboringNonNines(grid, strconv.Itoa(x) +"," +strconv.Itoa(i), visited)
		} else {
			break
		}
	}

	// Go down
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] != 9 && !visited[i][x] {
			visited[i][x] = true
			visited = getNeighboringNonNines(grid, strconv.Itoa(x) +"," +strconv.Itoa(i), visited)
		} else {
			break
		}
	}

	return visited
}