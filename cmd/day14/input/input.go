package input

import (
	"github.com/celso-patiri/aoc/cmd/common"
	"strings"
)

var inputFile = "input/input.txt"

var DebugInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func ParseInput(input string) (rules map[string]string, polymer string) {
	rules = map[string]string{}

	lines := strings.Split(input, "\n")

	polymer = lines[0]

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		rule := strings.Split(lines[i], "->")

		pattern, newChar := strings.TrimSpace(rule[0]), strings.TrimSpace(rule[1])

		rules[pattern] = newChar
	}

	return
}

func ParseInputFile() (map[string]string, string) {
	input := common.ReadInputFile(inputFile)
	return ParseInput(input)
}
