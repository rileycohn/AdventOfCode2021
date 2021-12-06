package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day5/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	linesPerPoint := [1000][1000]int{}

	// Iterate the input and start adding counts
	for _, line := range lines {
		// Split on ->
		coords := strings.Split(line, "->")
		leftCoords := strings.Split(strings.TrimSpace(coords[0]), ",")
		rightCoords := strings.Split(strings.TrimSpace(coords[1]), ",")

		leftX, _ := strconv.ParseInt(leftCoords[0], 10, 0)
		leftY, _ := strconv.ParseInt(leftCoords[1], 10, 0)
		rightX, _ := strconv.ParseInt(rightCoords[0], 10, 0)
		rightY, _ := strconv.ParseInt(rightCoords[1], 10, 0)

		// If the X's are the same, go from lower Y to higher Y and add counts to map
		if leftX == rightX {
			lowerY := leftY
			higherY := rightY
			if leftY > rightY {
				lowerY = rightY
				higherY = leftY
			}

			for y := lowerY; y <= higherY; y++ {
				linesPerPoint[leftX][y]++
			}
		} else if leftY == rightY {
			lowerX := leftX
			higherX := rightX
			if leftX > rightX {
				lowerX = rightX
				higherX = leftX
			}

			for x := lowerX; x <= higherX; x++ {
				linesPerPoint[x][leftY]++
			}
		} else {
			//Handle diagonals

			// Starting from first point, go until second point
			x := leftX
			y := leftY
			for true {
				linesPerPoint[x][y]++

				if x < rightX {
					x++
				} else if x > rightX {
					x--
				} else {
					break
				}

				if y < rightY {
					y++
				} else if y > rightY {
					y--
				}
			}
		}
	}

	//fmt.Println(linesPerPoint)
	find2Overlap(linesPerPoint)
}

func find2Overlap(grid [1000][1000]int) {
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				total++
			}
		}
	}

	fmt.Println("Total")
	fmt.Println(total)
}
