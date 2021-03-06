package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadLinesToStringList Reads lines from file into list
func ReadLinesToStringList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// ReadLinesToIntList Reads lines from file into list
func ReadLinesToIntList(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, x)
	}

	return lines, scanner.Err()
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}