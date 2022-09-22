package main

import (
	"fmt"
	"github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day8/part1"
)

var debugInput = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
var inputFile = "input.txt"

func main() {
	input := common.ReadInputFile(inputFile)

	fmt.Println("Part 1: ", part1.Run(input))
}
