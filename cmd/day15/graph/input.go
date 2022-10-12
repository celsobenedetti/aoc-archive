package graph

import (
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

func getNeighbors(curr, n int) (top, right, bottom, left int) {
	top = curr - n
	right = curr + 1
	left = curr - 1
	bottom = curr + n
	return
}

func ParseInput(input string) (graph common.AdJList, mArr []int) {
	matrix := make([][]int, 0)
	graph.CreateGraph()

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}

		matrix = append(matrix, make([]int, 0))
		for _, position := range strings.Split(line, "") {

			value, err := strconv.Atoi(position)
			if err != nil {
				panic("Invalid matrix int value")
			}

			matrix[i] = append(matrix[i], value)
		}
	}

	nMatrix := len(matrix)

	//top left corner
	curr := 0
	_, right, bottom, _ := getNeighbors(curr, nMatrix)
	graph.AddEdge(curr, right, matrix[0][1])
	graph.AddEdge(curr, bottom, matrix[1][0])

	//top right corner
	curr = nMatrix - 1
	_, _, bottom, left := getNeighbors(curr, nMatrix)
	graph.AddEdge(curr, left, matrix[0][nMatrix-2])
	graph.AddEdge(curr, bottom, matrix[1][nMatrix-1])

	//bottom left corner
	curr = (nMatrix) * (nMatrix - 1)
	top, right, _, _ := getNeighbors(curr, nMatrix)
	graph.AddEdge(curr, top, matrix[nMatrix-2][0])
	graph.AddEdge(curr, right, matrix[nMatrix-1][1])

	//bottom right corner
	curr = (nMatrix)*(nMatrix-1) + nMatrix - 1
	top, _, _, left = getNeighbors(curr, nMatrix)
	graph.AddEdge(curr, top, matrix[nMatrix-2][nMatrix-1])
	graph.AddEdge(curr, left, matrix[nMatrix-1][nMatrix-2])

	for i := 1; i < nMatrix-1; i++ {

		//first row
		curr = i
		_, right, bottom, left := getNeighbors(curr, nMatrix)
		// bottom := (nMatrix) + i
		// left := i - 1
		// right := i + 1

		graph.AddEdge(i, bottom, matrix[1][i])
		graph.AddEdge(i, left, matrix[0][i-1])
		graph.AddEdge(i, right, matrix[0][i+1])

		//last row
		curr = (nMatrix-1)*(nMatrix) + i
		top, right, _, left = getNeighbors(curr, nMatrix)
		// top = curr - nMatrix
		// left = curr - 1
		// right = curr + 1

		graph.AddEdge(curr, top, matrix[nMatrix-2][i])
		graph.AddEdge(curr, left, matrix[nMatrix-1][i-1])
		graph.AddEdge(curr, right, matrix[nMatrix-1][i+1])

		//first column
		curr = nMatrix * i
		top, right, bottom, _ = getNeighbors(curr, nMatrix)
		// top = curr - nMatrix
		// right = curr + 1
		// bottom = curr + nMatrix

		graph.AddEdge(curr, top, matrix[i-1][0])
		graph.AddEdge(curr, right, matrix[i][1])
		graph.AddEdge(curr, bottom, matrix[i+1][0])

		//last column
		curr = (nMatrix * i) + nMatrix - 1
		top, _, bottom, left = getNeighbors(curr, nMatrix)
		// top = curr - nMatrix
		// left = curr - 1
		// bottom = curr + nMatrix

		graph.AddEdge(curr, top, matrix[i-1][nMatrix-1])
		graph.AddEdge(curr, left, matrix[i][nMatrix-2])
		graph.AddEdge(curr, bottom, matrix[i+1][nMatrix-1])

		//rest of matrix
		for j := 1; j < nMatrix-1; j++ {
			curr = (i * nMatrix) + j

			// top = (i-1)*(nMatrix) + j
			// bottom = (i+1)*(nMatrix) + j
			// left = (i * nMatrix) + j - 1
			// right = (i * nMatrix) + j + 1
			top, right, bottom, left = getNeighbors(curr, nMatrix)

			graph.AddEdge(curr, top, matrix[i-1][j])
			graph.AddEdge(curr, bottom, matrix[i+1][j])
			graph.AddEdge(curr, left, matrix[i][j-1])
			graph.AddEdge(curr, right, matrix[i][j+1])
		}
	}

	for i := 0; i < nMatrix; i++ {
		for j := 0; j < nMatrix; j++ {
			mArr = append(mArr, matrix[i][j])
		}
	}

	return
}

func ParseInputFile() (common.AdJList, []int) {
	input := common.ReadInputFile(inputFile)
	return ParseInput(input)
}
