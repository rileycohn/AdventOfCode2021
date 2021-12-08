package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 8!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day8/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	answer := 0
	for _, line := range lines {
		// Line consists of instructions? | 4 digit output value?
		split := strings.Split(line, "|")
		signalPattern := split[0]
		outputValue := split[1]

		inputChars := strings.Split(signalPattern, " ")
		outputChars := strings.Split(outputValue, " ")

		valueToCode := make(map[int]string)
		codeToValue := make(map[string]string)

		for _, input := range inputChars {
			sortedString := SortString(input)
			switch len(input) {
			case 2:
				// 1
				valueToCode[1] = input
				codeToValue[sortedString] = "1"
			case 3:
				// 7
				valueToCode[7] = input
				codeToValue[sortedString] = "7"
			case 4:
				// 4
				valueToCode[4] = input
				codeToValue[sortedString] = "4"
			case 7:
				// 8
				valueToCode[8] = input
				codeToValue[sortedString] = "8"
			}
		}

		// Now go through the input again
		for _, input := range inputChars {
			sortedString := SortString(input)

			// Already found this one
			if codeToValue[sortedString] != "" {
				continue
			}

			// 1, 4, 7 chars used to determine the rest
			oneDiff := DiffCharCount(valueToCode[1], input)
			fourDiff := DiffCharCount(valueToCode[4], input)
			sevenDiff := DiffCharCount(valueToCode[7], input)

			switch len(input) {
			case 5:
				// 2,3,5
				if sevenDiff == 0 {
					codeToValue[sortedString] = "3"
				} else if fourDiff == 2 {
					codeToValue[sortedString] = "2"
				} else {
					codeToValue[sortedString] = "5"
				}
			case 6:
				// 0,6,9
				if oneDiff == 1 {
					codeToValue[sortedString] = "6"
				} else if fourDiff == 1 {
					codeToValue[sortedString] = "0"
				} else {
					codeToValue[sortedString] = "9"
				}
			}
		}

		// Now we have to go through the ouput
		s := ""
		for _, output := range outputChars {
			output = SortString(strings.TrimSpace(output))
			val := codeToValue[output]
			s += val
		}

		valInt, _ := strconv.ParseInt(s, 10, 0)
		answer += int(valInt)
	}

	fmt.Println(answer)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func DiffCharCount(orig string, newString string) int {
	count := 0
	for _, c := range orig {
		if !strings.Contains(newString, string(c)) {
			count++
		}
	}
	return count
}