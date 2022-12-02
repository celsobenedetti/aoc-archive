package common

import (
	"regexp"
	"strconv"
	"strings"
)

var whiteSpaceRegex = regexp.MustCompile(`\s+`)

func ClearWhiteSpaces(s string) string {
	result := whiteSpaceRegex.ReplaceAllString(s, " ")
	return strings.Trim(result, " ")
}

func ParseInputNumbers(line string) (numbers []int) {
	for _, n := range strings.Split(line, ",") {
		value := StrToInt(n)
		numbers = append(numbers, value)
	}
	return numbers
}

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("Invalid string input")
	}
	return n
}

func MakeCheckBoard() CheckBoard {
	board := make(CheckBoard, 5)

	for i := range board {
		board[i] = make(CheckRow, 5)
	}

	return board
}
