package common

import (
	"strings"
)

func Setup() (openableChars, closableChars CharSet) {
	openableChars = make(CharSet)
	closableChars = make(CharSet)

	openableChars["("] = ")"
	openableChars["["] = "]"
	openableChars["{"] = "}"
	openableChars["<"] = ">"

	closableChars[")"] = "("
	closableChars["]"] = "["
	closableChars["}"] = "{"
	closableChars[">"] = "<"

	return
}

func ParseInput(input string) (lines []string) {
	lines = make([]string, 0)

	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines
}
