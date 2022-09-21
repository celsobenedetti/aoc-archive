package main

import (
	"fmt"
	"strings"

	file "github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day7/common"
	"github.com/celso-patiri/aoc/cmd/day7/part1"
)

var inputFile = "input.txt"
var debugInput = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func main() {
	inputString := file.ReadInputFile(inputFile)
	input := common.ParseStringsToInts(strings.Split(inputString, ","))

	fmt.Println("Part 1: ", part1.Run(input))
}
