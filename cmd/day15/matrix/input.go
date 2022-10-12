package matrix

import (
	"math"
	"strconv"
	"strings"

	"github.com/celso-patiri/aoc/cmd/common"
)

var DebugInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

var inputFile = `input/input.txt`

func createRow(line string, i, itRow, itColumn, rowRepeatCount, columnRepeatCount int) (row []Node) {

	columnAcc := columnRepeatCount * len(line)
	rowAcc := rowRepeatCount * len(line)

	inc := rowRepeatCount + columnRepeatCount

	for j, position := range strings.Split(line, "") {

		v, err := strconv.Atoi(position)
		if err != nil {
			panic("Invalid Matrix int input")
		}

		value := (v + inc)
		if value > 9 {
			value = (v + inc) % 10
			value++
		}

		point := Point{I: i + rowAcc, J: j + columnAcc}
		row = append(row, Node{Weight: value, Prev: nil, Dist: math.MaxInt, Position: point})
	}
	return

}

func ParseInput(input string, dimension int) (matrix Matrix) {
	rowRepeatCount := 0
	for itRow := 0; itRow < dimension; itRow++ {
		for i, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}

			row := []Node{}

			columnRepeatCount := 0
			for itColumn := 0; itColumn < dimension; itColumn++ {
				row = append(row, createRow(line, i, itRow, itColumn, rowRepeatCount, columnRepeatCount)...)
				columnRepeatCount++
			}

			matrix = append(matrix, row)
		}
		rowRepeatCount++
	}
	return
}

func ParseInputFile() string {
	return common.ReadInputFile(inputFile)
}
