package main

import (
	"fmt"
	"github.com/celso-patiri/aoc/cmd/day2/part1"
	"github.com/celso-patiri/aoc/cmd/day2/part2"
)

func main() {

	position1 := part1.Run()
	position2 := part2.Run()

	fmt.Printf("Part 1: %d\n", position1.X*position1.Y)
	fmt.Printf("Part 2: %d\n", position2.X*position2.Y)
}
