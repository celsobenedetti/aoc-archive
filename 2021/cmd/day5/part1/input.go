package part1

import (
	"github.com/celso-patiri/aoc/cmd/day5/common"
	"strings"
)

func evaluatePoints(points []string) (start, end common.Point) {
	x1, y1 := common.ParsePoint(points[0])
	x2, y2 := common.ParsePoint(points[1])

	if x1 != x2 && y1 != y2 {
		return common.Point{X: -1, Y: -1}, common.Point{X: -1, Y: -1}
	}

	if x1 < x2 || y1 < y2 {
		return common.Point{X: x1, Y: y1}, common.Point{X: x2, Y: y2}
	} else {
		return common.Point{X: x2, Y: y2}, common.Point{X: x1, Y: y1}
	}
}

func drawGrid(grid common.Grid, start, end common.Point) int {
	count := 0
	if start.X == end.X {
		for i := start.Y; i <= end.Y; i++ {
			grid[i][start.X]++

			if grid[i][start.X] == 2 {
				count++
			}
		}
	} else {
		for j := start.X; j <= end.X; j++ {
			grid[start.Y][j]++

			if grid[start.Y][j] == 2 {
				count++
			}
		}
	}
	return count
}

func Run(input string) (grid common.Grid, count int) {
	grid = common.MakeGrid(1000)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		points := strings.Split(line, "->")
		start, end := evaluatePoints(points)

		if start.X >= 0 {
			count += drawGrid(grid, start, end)
		}
	}
	return grid, count
}
