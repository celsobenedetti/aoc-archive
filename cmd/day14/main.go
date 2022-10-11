package main

import (
	"fmt"

	"github.com/celso-patiri/aoc/cmd/day14/input"
	"github.com/celso-patiri/aoc/cmd/day14/polymer"
)

func main() {
	rules, base := input.ParseInputFile()

	part1 := polymer.BuildPolymer(base, rules, 10)
	part2 := polymer.BuildPolymer(base, rules, 40)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
