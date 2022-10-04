package part2

import (
	"sort"
	"strings"

	"github.com/celso-patiri/aoc/cmd/day10/common"
)

type CharSet = common.CharSet

var openableChars CharSet
var closableChars CharSet
var charValues map[string]int

func getIncompleteLines(lines []string) (result []string) {
	result = make([]string, 0)

	for _, line := range lines {
		if isLineValid(line) {
			result = append(result, line)
		}
	}

	return
}

func isLineValid(line string) bool {
	stack := common.MakeNewStack()

	for _, char := range strings.Split(line, "") {
		if openableChars.Has(char) {
			stack.Push(char)
		} else {
			if stack.Top() == closableChars[char] {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return true
}

func getLineCompletion(line string) string {
	var completion strings.Builder

	stack := common.MakeNewStack()

	for _, char := range strings.Split(line, "") {
		if openableChars.Has(char) {
			stack.Push(char)
		} else {
			stack.Pop()
		}
	}

	for stack.Size() > 0 {
		closeChar := openableChars[stack.Pop()]
		completion.WriteString(closeChar)
	}

	return completion.String()
}

func getCompletionScore(line string) (score int) {
	for _, char := range strings.Split(line, "") {
		score = score * 5
		score += charValues[char]
	}
	return
}

func Run(lines []string) int {
	openableChars, closableChars = common.Setup()

	charValues = make(map[string]int)
	charValues[")"] = 1
	charValues["]"] = 2
	charValues["}"] = 3
	charValues[">"] = 4

	incompleteLines := getIncompleteLines(lines)

	scores := make([]int, 0)
	for _, line := range incompleteLines {
		scores = append(scores, getCompletionScore(getLineCompletion(line)))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	middle := len(scores) / 2
	return scores[middle]
}
