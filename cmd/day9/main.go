package main

import (
	"fmt"
	file "github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day9/common"
	"github.com/celso-patiri/aoc/cmd/day9/part1"
)

var debugInput = `2199943210
3987894921
9856789892
8767896789
9899965678
`

var fileName = "input.txt"

func main() {

	inputString := file.ReadInputFile(fileName)
	input := common.ParseInput(inputString)

	fmt.Println("Part 1: ", part1.Run(input))
}
