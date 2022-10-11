package main

import (
	"fmt"

	"github.com/celso-patiri/aoc/cmd/day13/input"
)

func main() {
	page, instructions := input.ParseInputFile()

	page = page.InterpretFold(instructions[0])

	fmt.Println("Part 1: ", page.CountDots())

	for i := 1; i < len(instructions); i++ {
		page = page.InterpretFold(instructions[i])
	}

	fmt.Println("Part 2:")
	page.PrintPaper()
}
