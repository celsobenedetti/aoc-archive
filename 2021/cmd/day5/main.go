package main

import (
	"fmt"
	"github.com/celso-patiri/aoc/cmd/common"
	part1 "github.com/celso-patiri/aoc/cmd/day5/part1"
	part2 "github.com/celso-patiri/aoc/cmd/day5/part2"
)

var inputFile = "input.txt"
var debugInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func main() {
	input := common.ReadInputFile(inputFile)
	_, count1 := part1.Run(input)
	_, count2 := part2.Run(input)

	fmt.Println("Part 1: ", count1)
	fmt.Println("Part 2: ", count2)
}
