package matrix

import (
	"math"
	"strconv"
	"strings"

	"github.com/celso-patiri/aoc/cmd/common"
)

var DebugInput = `742
521
581`

var inputFile = `input/input.txt`

func createRow(line string, i, itRow, itColumn int) (row []Node) {
	inc := itRow + itColumn

	for j, position := range strings.Split(line, "") {

		v, err := strconv.Atoi(position)
		if err != nil {
			panic("Invalid Matrix int input")
		}

		value := (v + inc) % 10
		if value == 0 {
			value++
		}

		point := Point{I: i + itRow, J: j + itColumn}
		row = append(row, Node{Weight: value, Seen: false, Prev: nil, Dist: math.MaxInt, Position: point})
	}
	return

}

func ParseInput(input string, dimension int) (matrix Matrix) {
	for itRow := 0; itRow < dimension; itRow++ {
		for i, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}

			row := []Node{}
			for itColumn := 0; itColumn < dimension; itColumn++ {
				row = append(row, createRow(line, i, itRow, itColumn)...)
			}

			matrix = append(matrix, row)
		}
	}
	return
}

func ParseInputFile() string {
	return common.ReadInputFile(inputFile)
}
