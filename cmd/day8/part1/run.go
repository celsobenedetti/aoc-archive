package part1

import (
	"strings"
)

var acceptedLengths map[int]bool

func ParseInput(input string) (result []string) {
	result = make([]string, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		segments := strings.Split(line, "|")

		if len(segments) == 2 {

			codes := strings.Split(segments[1], " ")

			for _, code := range codes {
				result = append(result, code)
			}
		}
	}
	return
}

func Run(input string) (result int) {
	output := ParseInput(input)

	acceptedLengths = make(map[int]bool)
	acceptedLengths[2] = true
	acceptedLengths[3] = true
	acceptedLengths[4] = true
	acceptedLengths[7] = true

	for _, code := range output {
		if acceptedLengths[len(code)] {
			result++
		}
	}

	return
}
