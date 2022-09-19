package common

import (
	"log"
	"strconv"
	"strings"
)

func coordsToInt(a, b string) (x, y int) {
	x, errx := strconv.Atoi(strings.Trim(a, " "))
	y, erry := strconv.Atoi(strings.Trim(b, " "))
	if errx != nil || erry != nil {
		log.Fatalln("Error: invalid string")
		return
	}
	return x, y
}

func parsePoint(point string) (x, y int) {
	coords := strings.Split(point, ",")
	return coordsToInt(coords[0], coords[1])
}

func evaluatePoints(points []string) (start, end Point) {
	x1, y1 := parsePoint(points[0])
	x2, y2 := parsePoint(points[1])

	if x1 != x2 && y1 != y2 {
		return Point{-1, -1}, Point{-1, -1}
	}

	if x1 < x2 || y1 < y2 {
		return Point{x1, y1}, Point{x2, y2}
	} else {
		return Point{x2, y2}, Point{x1, y1}
	}
}

func drawGrid(grid Grid, start, end Point) int {
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

func ParseInput(input string) (grid Grid, count int) {
	grid = MakeGrid(1000)

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
