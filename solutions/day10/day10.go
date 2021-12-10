package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Day 10!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day10/day10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var scores []int
	for _, line := range lines {
		total := 0
		autoComplete := processRow(line)
		// We only care about incomplete rows now
		if autoComplete != "" {
			for _, c := range autoComplete {
				total = total * 5
				total = total + getAutocompleteCharValue(string(c))
			}

			scores = append(scores, total)
		}
	}

	fmt.Println(scores)
	// Now sort
	sort.Ints(scores)
	fmt.Println(scores[len(scores) / 2])
}

func processRow(row string) string {
	// Create a stack to determine if the row is valid
	var stack []string
	for _, char := range row {
		// Append open chars to the stack
		if isOpenChar(string(char)) {
			stack = append(stack, string(char))
		} else {
			// If this is a closing char, the top of the stack needs to be its matching open character
			if len(stack) == 0 || stack[len(stack) - 1] != getMatchingOpenChar(string(char)) {
				// This is not a match, return empty
				return ""
			} else {
				// Slice the stack to "pop" off the stack
				stack = stack[:len(stack) - 1]
			}
		}
	}

	// If we get here, the row is valid, just missing some vals, figure those out
	autoComplete := ""
	for _, openChar := range stack {
		autoComplete += getMatchingClosingChar(openChar)
	}

	// Reverse the order
	return lib.Reverse(autoComplete)
}

func isOpenChar(char string) bool {
	return char == "(" || char == "{" || char == "[" || char == "<"
}

func getMatchingOpenChar(char string) string {
	switch char {
	case ")":
		return "("
	case "}":
		return "{"
	case "]":
		return "["
	case ">":
		return "<"
	default:
		return ""
	}
}

func getMatchingClosingChar(char string) string {
	switch char {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	default:
		return ""
	}
}

func getAutocompleteCharValue(char string) int {
	charVal := map[string]int{
		"": 0,
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	return charVal[char]
}
