package part1

import (
	"strings"

	"github.com/celso-patiri/aoc/cmd/day10/common"
)

type CharSet = common.CharSet

var openableChars CharSet
var closableChars CharSet
var charValues map[string]int

func evaluateLine(line string) (errorValue int) {
	stack := common.MakeNewStack()

	for _, char := range strings.Split(line, "") {
		if openableChars.Has(char) {
			stack.Push(char)
		} else {
			if stack.Top() == closableChars[char] {
				stack.Pop()
			} else {
				return charValues[char]
			}
		}
	}
	return 0
}

func Run(lines []string) (total int) {
	charValues = make(map[string]int)
	charValues[")"] = 3
	charValues["]"] = 57
	charValues["}"] = 1197
	charValues[">"] = 25137

	openableChars, closableChars = common.Setup()

	for _, line := range lines {
		total += evaluateLine(line)
	}

	return total
}
