package main

import (
	"fmt"

	"github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day11/input"
	"github.com/celso-patiri/aoc/cmd/day11/octopuses"
)

var debugInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var inputFile = "input.txt"

func newFlashMatrix(n int) (m [][]bool) {
	m = make([][]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
	}
	return
}

func visitPosition(i, j int) (total int) {
	n := octopi.Size()

	if i < 0 || i == n || j < 0 || j == n {
		return 0
	}

	if !hasFlashed[i][j] {
		octopi.Inc(i, j)
		if octopi.Get(i, j) == 0 {
			hasFlashed[i][j] = true
			total++

			top := i - 1
			bottom := i + 1
			right := j + 1
			left := j - 1

			total += visitPosition(top, left)
			total += visitPosition(top, j)
			total += visitPosition(top, right)
			total += visitPosition(i, right)
			total += visitPosition(bottom, right)
			total += visitPosition(bottom, j)
			total += visitPosition(bottom, left)
			total += visitPosition(i, left)
		}
	}
	return total
}

type Octopi = octopuses.Octopi

var octopi octopuses.Octopi
var hasFlashed [][]bool

func Run(input [][]int) (total, firstSync int) {
	octopi = Octopi{}
	octopi.Make(input)

	n := octopi.Size()
	firstSync = -1

	for step := 0; firstSync < 0; step++ {
		hasFlashed = newFlashMatrix(n)
		sync := true

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if step < 100 {
					total += visitPosition(i, j)
				} else {
					visitPosition(i, j)
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if octopi.Get(i, j) > 0 {
					sync = false
				}
			}
		}

		if sync && firstSync < 0 {
			firstSync = step + 1
		}

	}
	return
}

func main() {
	inputString := common.ReadInputFile(inputFile)
	octopi := input.ParseInput(inputString)

	total, firstSync := Run(octopi)
	fmt.Println("Part 1: ", total)
	fmt.Println("Part 2: ", firstSync)
}
