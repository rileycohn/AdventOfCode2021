package main

import (
	"adventOfCode2021/lib"
	"fmt"
)

func main() {
	fmt.Println("Day 9!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day10/day10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, line := range lines {
		illegalChar := processRow(line)
		fmt.Println(illegalChar)
		total = total + getIllegalCharValue(illegalChar)
	}

	fmt.Println(total)
}

func processRow(row string) string {
	// Create a stack to determine if the row is valid
	stack := []string{}
	for _, char := range row {
		// Append open chars to the stack
		if isOpenChar(string(char)) {
			stack = append(stack, string(char))
		} else {
			// If this is a closing char, the top of the stack needs to be its matching open character
			if len(stack) == 0 || stack[len(stack) - 1] != getMatchingOpenChar(string(char)) {
				// This is not a match, return it
				return string(char)
			} else {
				// Slice the stack to "pop" off the stack
				stack = stack[:len(stack) - 1]
			}
		}
	}

	// If we get here, the row is valid
	return ""
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

func getIllegalCharValue(char string) int {
	illegalChar := map[string]int{
		"": 0,
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	return illegalChar[char]
}
