package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Day 3!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day3/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// We need 2 lists
	oxRatingLines := make([]string, len(lines))
	copy(oxRatingLines, lines)

	co2RatingLines := make([]string, len(lines))
	copy(co2RatingLines, lines)

	finalox := ""
	finalco2 := ""

	stringLen := len(lines[0])

	// First, get the most common bit
	for i := 0; i < stringLen; i++ {
		// Now that we have the most common bit, we need to update the ox and co2 lists for non-matchers
		if finalox == "" {
			commonBitAtIndex := mostCommonBitForList(oxRatingLines, i)
			tmp := make([]string, 0)
			for _, oxLine := range oxRatingLines {
				if commonBitAtIndex == string(oxLine[i]) {
					tmp = append(tmp, oxLine)
				}
			}

			oxRatingLines = make([]string, len(tmp))
			copy(oxRatingLines, tmp)

			if len(oxRatingLines) == 1 {
				finalox = oxRatingLines[0]
			}
		}

		if finalco2 == "" {
			commonBitAtIndex := mostCommonBitForList(co2RatingLines, i)
			tmp2 := make([]string, 0)
			for _, co2Line := range co2RatingLines {
				if commonBitAtIndex != string(co2Line[i]) {
					tmp2 = append(tmp2, co2Line)
				}
			}

			co2RatingLines = make([]string, len(tmp2))
			copy(co2RatingLines, tmp2)

			if len(co2RatingLines) == 1 {
				finalco2 = co2RatingLines[0]
			}
		}
	}

	finaloxInt, _ := strconv.ParseInt(finalox, 2, 64)
	finalco2Int, _ := strconv.ParseInt(finalco2, 2, 64)

	fmt.Println(finaloxInt * finalco2Int)
}

func mostCommonBitForList(lines []string, index int) string {
	zeroCount := 0
	oneCount := 0

	for _, line := range lines {
		bit := line[index]
		if string(bit) == "0" {
			zeroCount++
		} else {
			oneCount++
		}
	}

	if zeroCount > oneCount {
		return "0"
	} else {
		return  "1"
	}
}
