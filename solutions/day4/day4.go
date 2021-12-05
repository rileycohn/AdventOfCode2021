package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Day 4!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day4/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bingoInput := strings.Split(lines[0], ",")

	bingoBoards := [] [5][5]int64{}
	rowNum := 0
	// Now we need to read the lines into a 2D array
	currentBoard := [5][5]int64{}
	for index, line := range lines {

		// Skip the bingo input and the empty lines (new input)
		if index == 0 || line == "" {

			if rowNum != 0 {
				bingoBoards = append(bingoBoards, currentBoard)
			}

			rowNum = 0

			currentBoard = [5][5]int64{}
			continue
		}

		rowInput := strings.Fields(line)
		for i, val := range rowInput {
			currentBoard[rowNum][i], _ = strconv.ParseInt(val, 10, 64)
		}

		rowNum++
	}

	// Now we have all the boards, loop through the input and find winners
	// Start at 5 numbers since that's the min number to win the game
	for i := 5; i < len(bingoInput); i++ {
		losers := [] [5][5]int64{}
		inputSlice := bingoInput[0:i]
		fmt.Println(inputSlice)
		for boardNum := 0; boardNum < len(bingoBoards); boardNum++ {
			isWinner := isBingoBoardAWinner(bingoBoards[boardNum], inputSlice)
			if isWinner {
				sumOfUnmarked(bingoBoards[boardNum], inputSlice)
				continue
			} else {
				losers = append(losers, bingoBoards[boardNum])
			}
		}

		bingoBoards = losers

		if len(bingoBoards) == 0 {
			break
		}
	}

}

func isBingoBoardAWinner(board [5][5]int64, input []string) bool {
	for i := 0; i < 5; i++ {
		// Check rows against input
		inARow := 0
		for j := 0; j < 5; j++ {
			if contains(input, board[i][j]) {
				inARow++
			}
		}

		if inARow == 5 {
			return true
		}
	}

	// If we didn't find a row winner, check columns
	for i := 0; i < 5; i++ {
		// Check rows against input
		inARow := 0
		for j := 0; j < 5; j++ {
			if contains(input, board[j][i]) {
				inARow++
			}
		}

		if inARow == 5 {
			return true
		}
	}

	return false
}

func contains(s []string, e int64) bool {
	for _, a := range s {
		inputVal, _ := strconv.ParseInt(a, 10, 64)
		if  inputVal == e {
			return true
		}
	}
	return false
}

func sumOfUnmarked(board [5][5]int64, input []string) {
	sum, _ := strconv.ParseInt("0", 10, 64)
	lastCalled, _ := strconv.ParseInt(input[len(input) - 1], 10, 64)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !contains(input, board[j][i]) {
				sum += board[j][i]
			}
		}
	}
	fmt.Println("Solution")
	fmt.Println(sum * lastCalled)
}
