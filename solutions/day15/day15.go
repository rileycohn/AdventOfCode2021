package main

import (
	"adventOfCode2021/lib"
	"adventOfCode2021/priorityQueue"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 16!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day16/day16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(grid))
	}

	for i, line := range lines {
		for j, c := range line {
			cInt, _ := strconv.ParseInt(string(c), 10, 0)
			grid[i][j] = int(cInt)
		}
	}

	part1(grid)

	// Part 2 requires us to expand the grid, then solve the path the same way
	part2(makeBigGrid(grid))

}

func part1(grid [][]int) {
	pathFinder(grid)
}

func part2(grid [][]int) {
	pathFinder(grid)
}

func pathFinder(grid [][]int) {
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid))
	}

	// BFS
	pq := make(priorityQueue.PriorityQueue, 1)

	firstItem := priorityQueue.Item{Value: "0,0"}

	pq[0] = &firstItem

	heap.Init(&pq)
	for pq.Len() > 0 {
		path := heap.Pop(&pq).(*priorityQueue.Item)
		val := strings.Split(path.Value, ",")
		x, _ := strconv.Atoi(val[0])
		y, _ := strconv.Atoi(val[1])

		if visited[x][y] {
			continue
		}
		if x == len(grid) - 1 && y == len(grid[0]) - 1 {
			fmt.Println(path.Priority)
			break
		}
		visited[x][y] = true

		movements := []xy{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}

		for _, movement := range movements {
			if movement.x < len(grid) && movement.x >= 0 && movement.y >= 0 && movement.y < len(grid[0]) {
				if !visited[movement.x][movement.y] {
					newVal := strconv.Itoa(movement.x) + "," + strconv.Itoa(movement.y)
					newPriority := path.Priority + grid[movement.x][movement.y]
					item := &priorityQueue.Item{
						Value:    newVal,
						Priority: newPriority,
						Index:    pq.Len(),
					}
					heap.Push(&pq, item)
					pq.Update(item, item.Value, newPriority)
				}
			}
		}
	}
}

type xy struct {
	x int
	y int
}

func makeBigGrid(grid [][]int) [][]int {
	bigGrid := make([][]int, len(grid) * 5)
	for i := range bigGrid {
		bigGrid[i] = make([]int, len(bigGrid))
	}

	xMod := len(grid)
	yMod := len(grid[0])

	for i := 0; i < len(bigGrid); i++ {
		for j := 0; j < len(bigGrid[0]); j++ {
			bigGrid[i][j] = wrap(grid[i % xMod][j % yMod] + (i / len(grid)) + (j / len(grid[0])))
		}
	}

	return bigGrid
}

func wrap(x int) int {
	return (x - 1) % 9 + 1
}