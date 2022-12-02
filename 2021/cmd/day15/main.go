package main

import (
	"fmt"

	"github.com/celso-patiri/aoc/cmd/day15/graph"
	"github.com/celso-patiri/aoc/cmd/day15/matrix"
)

func main() {
	input := matrix.ParseInputFile()

	g1, arr1 := graph.ParseInput(input)
	part1GraphImpl := graph.CalculateRiskGraph(g1, arr1)

	fmt.Println("[Graph Impl]Part 1:", part1GraphImpl)

	m1 := matrix.ParseInput(input, 1)
	m2 := matrix.ParseInput(input, 5)

	fmt.Println("[Matrix Impl]Part 1: ", m1.CalculateRisk())
	fmt.Println("[Matrix Impl]Part 2: ", m2.CalculateRisk())

}
