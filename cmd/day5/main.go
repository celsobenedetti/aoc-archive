package main

import (
	"fmt"
	"github.com/celso-patiri/aoc/cmd/common"
	day5 "github.com/celso-patiri/aoc/cmd/day5/common"
)

var input = "input.txt"
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
	_, count := day5.ParseInput(common.ReadInputFile(input))

	fmt.Println("Part 1: ", count)
}
