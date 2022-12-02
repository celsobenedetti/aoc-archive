package main

import (
	"fmt"

	input "github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day10/common"
	"github.com/celso-patiri/aoc/cmd/day10/part1"
	"github.com/celso-patiri/aoc/cmd/day10/part2"
)

var debugInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

var fileName = "input.txt"

func main() {
	inputString := input.ReadInputFile(fileName)
	lines := common.ParseInput(inputString)

	fmt.Println("Part 1: ", part1.Run(lines))
	fmt.Println("Part 2: ", part2.Run(lines))

}
