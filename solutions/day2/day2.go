package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day2/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splitLine[1])
		switch splitLine[0] {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	fmt.Println(horizontal * depth)
}
