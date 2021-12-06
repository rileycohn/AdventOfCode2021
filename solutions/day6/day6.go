package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 6!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day6/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	vals := strings.Split(lines[0], ",")
	state := make([]int64, len(vals))
	daysMap := make(map[int64]int64)
	for i := range state {
		intVal, _ := strconv.ParseInt(vals[i], 10, 64)
		daysMap[intVal] = daysMap[intVal] + 1
	}

	// 80 day simulation
	for day := 0; day < 256; day++ {
		newMap := make(map[int64]int64)
		for fishDays := 8; fishDays >= 0; fishDays-- {
			numFish := daysMap[int64(fishDays)]
			if fishDays == 0 {
				daysMap[int64(fishDays)] = 0
				newMap[6] = newMap[6] + numFish
				newMap[8] = daysMap[8] + numFish
			} else {
				daysMap[int64(fishDays)] = 0
				newMap[int64(fishDays-1)] = numFish
			}
		}

		daysMap = make(map[int64]int64)
		for k,v := range newMap {
			daysMap[k] = v
		}

		var total int64 = 0
		for _, value := range daysMap {
			total = total + value
		}
	}

	var total int64 = 0
	for _, value := range daysMap {
		total = total + value
	}

	fmt.Println(total)
}
