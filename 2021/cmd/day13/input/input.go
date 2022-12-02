package input

import (
	"log"
	"strconv"
	"strings"

	"github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day13/paper"
)

var DebugInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7`

var inputFile = "./input/input.txt"

func ParseInput(input string) (page paper.Paper, instructions []paper.Instruction) {
	bigPaper := make(paper.Paper, 4000)

	lines := strings.Split(input, "\n")

	maxX, maxY := 0, 0

	i := 0
	for _, line := range lines {
		if line == "" {
			break
		}

		if len(bigPaper[i]) == 0 {
			bigPaper[i] = make([]bool, 4000)
		}

		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal("Invalid coordinate string")
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal("Invalid coordinate string")
		}

		if len(bigPaper[y]) == 0 {
			bigPaper[y] = make([]bool, 2000)
		}

		bigPaper[y][x] = true

		if maxX < x {
			maxX = x
		}

		if maxY < y {
			maxY = y
		}

		i++
	}

	for i := 0; i <= maxY; i++ {
		page = append(page, make([]bool, 0))
		for j := 0; j <= maxX; j++ {
			page[i] = append(page[i], bigPaper[i][j])
		}
	}

	for line := i; line < len(lines); line++ {
		if lines[line] == "" {
			continue
		}

		parts := strings.Split(lines[line], " ")
		instruction := strings.Split(parts[2], "=")

		v, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal("Invalid instruction string")
		}

		instructions = append(instructions, paper.Instruction{Axis: paper.Axis(instruction[0]), Position: v})
	}
	return
}

func ParseInputFile() (page paper.Paper, instructions []paper.Instruction) {
	inputString := common.ReadInputFile(inputFile)
	return ParseInput(inputString)
}
