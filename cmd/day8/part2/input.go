package part2

import (
	"strings"
)

type iterationInput [][]string

func (ii *iterationInput) newLine() {
	*ii = append(*ii, make([]string, 0))
}

func (ii iterationInput) newCode(line int, code string) {
	ii[line] = append(ii[line], code)
}

func ParseInput(input string) (codes, output iterationInput) {
	codes = make(iterationInput, 0)
	output = make(iterationInput, 0)

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		segments := strings.Split(line, "|")

		if len(segments) == 2 {
			codes.newLine()
			output.newLine()

			outputSplit := strings.Split(segments[1], " ")
			codesSplit := strings.Split(segments[0], " ")

			for _, code := range outputSplit {
				output.newCode(i, code)
			}

			for _, code := range codesSplit {
				codes.newCode(i, code)
			}
		}
	}
	return
}
