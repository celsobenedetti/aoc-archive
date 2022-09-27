package part1

import (
	"github.com/celso-patiri/aoc/cmd/day9/common"
)

type Point = common.Point

func Run(matrix [][]int) (sum int) {
	lowPoints := make([]common.Point, 0)
	lines := len(matrix) - 1
	rows := len(matrix[0]) - 1

	for i := 1; i < lines; i++ {
		for j := 1; j < rows; j++ {
			current := matrix[i][j]
			if current < matrix[i-1][j] && current < matrix[i+1][j] && current < matrix[i][j-1] && current < matrix[i][j+1] {
				lowPoints = append(lowPoints, Point{X: j, Y: i})
			}
		}
	}

	for i := 1; i < lines; i++ {
		current := matrix[i][0]
		if current < matrix[i-1][0] && current < matrix[i+1][0] && current < matrix[i][1] {
			lowPoints = append(lowPoints, Point{X: 0, Y: i})
		}
		current = matrix[i][rows]
		if current < matrix[i-1][rows] && current < matrix[i+1][rows] && current < matrix[i][rows-1] {
			lowPoints = append(lowPoints, Point{X: rows, Y: i})
		}
	}

	for j := 1; j < rows; j++ {
		current := matrix[0][j]
		if current < matrix[0][j-1] && current < matrix[0][j+1] && current < matrix[1][j] {
			lowPoints = append(lowPoints, Point{X: j, Y: 0})
		}
		current = matrix[lines][j]
		if current < matrix[lines][j-1] && current < matrix[lines][j+1] && current < matrix[lines-1][j] {
			lowPoints = append(lowPoints, Point{X: j, Y: lines})
		}
	}

	if matrix[0][0] < matrix[0][1] && matrix[0][0] < matrix[1][0] {
		lowPoints = append(lowPoints, Point{X: 0, Y: 0})
	}

	if matrix[0][rows] < matrix[0][rows-1] && matrix[0][rows] < matrix[1][rows] {
		lowPoints = append(lowPoints, Point{X: rows, Y: 0})
	}

	if matrix[lines][0] < matrix[lines][1] && matrix[lines][0] < matrix[lines-1][0] {
		lowPoints = append(lowPoints, Point{X: 0, Y: lines})
	}

	if matrix[lines][rows] < matrix[lines][rows-1] && matrix[lines][rows] < matrix[lines-1][rows] {
		lowPoints = append(lowPoints, Point{X: rows, Y: lines})
	}

	for i := 0; i < len(lowPoints); i++ {
		sum += matrix[lowPoints[i].Y][lowPoints[i].X] + 1
	}

	return
}
