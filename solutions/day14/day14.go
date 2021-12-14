package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Day 14!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day14/day14.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	polymerTemplate := lines[0]
	pairInsertions := map[string]string{}
	for i := 2; i < len(lines); i++ {
		split := strings.Split(lines[i], " -> ")
		pairInsertions[split[0]] = split[1]
	}

	fmt.Println("Part 1: ")
	part1(polymerTemplate, pairInsertions)

	fmt.Println("Part 2: ")
	part2(polymerTemplate, pairInsertions)
}

func part1(template string, insertions map[string]string) {
	updatedTemplate := template
	// 10 steps
	for i := 0; i < 10; i++ {
		// first and second pointers
		for f := 0; f < len(updatedTemplate) - 1; f++ {
			s := f + 1
			pair := string(updatedTemplate[f]) + string(updatedTemplate[s])
			if val, ok := insertions[pair]; ok{
				// Insert val between f and s
				updatedTemplate = updatedTemplate[:s] + val + updatedTemplate[s:]

				// Increase step by 1
				f++
				s++
			}
		}
		//fmt.Println(i)
		//fmt.Println(updatedTemplate)
	}

	maxMinusMin(updatedTemplate)
}

func part2(template string, insertions map[string]string) {
	// Maintaining the actual string in memory is not feasible, maintain a count of pairs instead
	pairCounts := map[string]int64{}
	// Initialize with the template string
	for f := 0; f < len(template) - 1; f++ {
		s := f + 1
		pair := string(template[f]) + string(template[s])
		pairCounts[pair]++
	}

	for i := 0; i < 40; i++ {
		newPairCounts := map[string]int64{}
		// Loop through the insertions (AB)
		for key, value := range insertions {
			// How many pairs exist with this key
			existing := pairCounts[key]

			// Increase counter for AX and XB
			newPairCounts[string(key[0]) + value] += existing
			newPairCounts[value + string(key[1])] += existing
		}

		pairCounts = newPairCounts
	}

	// Need to do the count now
	totalCounts := map[string]int64{}
	totalCounts[string(template[0])] = 1
	for k, v := range pairCounts {
		totalCounts[string(k[1])] += v
	}

	fmt.Println(totalCounts)
	maxMinusMinMap(totalCounts)
}

func charCountInString(s string) map[string]int64 {
	counts := map[string]int64{}
	for _, c := range s {
		counts[string(c)] = counts[string(c)] + 1
	}

	return counts
}

func maxMinusMin(updatedTemplate string) {
	counts := charCountInString(updatedTemplate)
	var max int64 = 0
	var min int64 = math.MaxInt64
	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	fmt.Println(max - min)
}

func maxMinusMinMap(counts map[string]int64) {
	var max int64 = 0
	var min int64 = math.MaxInt64
	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	fmt.Println(max - min)
}