package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Day 12!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day12/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	graph := buildGraph(lines)
	var visited []string
	numPaths := findPaths(graph, "start", visited, true)

	fmt.Println(numPaths)
}

func buildGraph(lines []string) map[string][]string {
	graph := map[string][]string{}
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		start := nodes[0]
		end := nodes[1]

		if start != "end" && end != "start" {
			if graph[start] != nil {
				graph[start] = append(graph[start], end)
			} else {
				graph[start] = []string{end}
			}
		}

		if start != "start" && end != "end" {
			if graph[end] != nil {
				graph[end] = append(graph[end], start)
			} else {
				graph[end] = []string{start}
			}
		}
	}

	return graph
}

func findPaths(graph map[string][]string, edge string, visited []string, allowTwoVisitsToSmallCave bool) int {
	if edge == "end" {
		return 1
	}

	// Don't allow lowercase caves to be visited more than once
	newAllowTwoVisitsToSmallCave := allowTwoVisitsToSmallCave
	if unicode.IsLower(rune(edge[0])) && lib.ContainsString(visited, edge) {
		// Only allow the double visit one time, so continue for now
		if newAllowTwoVisitsToSmallCave {
			newAllowTwoVisitsToSmallCave = false
		} else {
			return 0
		}
	}

	visited = append(visited, edge)

	numPaths := 0

	// Get all the edges from this edge node in the graph
	edges := graph[edge]

	for _, ed := range edges {
		numPaths += findPaths(graph, ed, visited, newAllowTwoVisitsToSmallCave)
	}

	return numPaths
}