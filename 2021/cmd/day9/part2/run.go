package part2

import (
	"github.com/celso-patiri/aoc/cmd/day9/common"
	"sort"
)

type Point = common.Point
type Matrix = common.Matrix

var visitedRegions VisitedRegions
var regionDepths Matrix

func visitRegion(region Point, currentDepth int) (size int) {
	depth := regionDepths.Get(region)
	if depth < 9 && depth > currentDepth {
		return calculateBasinSize(region)
	}
	return 0
}

func calculateBasinSize(current Point) (size int) {
	top := Point{X: current.X, Y: current.Y - 1}
	right := Point{X: current.X + 1, Y: current.Y}
	bottom := Point{X: current.X, Y: current.Y + 1}
	left := Point{X: current.X - 1, Y: current.Y}

	currentDepth := regionDepths.Get(current)

	//1. Check all neighbors, recurse if neighbor is not visited and 9 > neighbor > current
	if regionDepths.Contains(top) && !visitedRegions.IsVisited(top) {
		size += visitRegion(top, currentDepth)
	}
	if regionDepths.Contains(right) && !visitedRegions.IsVisited(right) {
		size += visitRegion(right, currentDepth)
	}
	if regionDepths.Contains(bottom) && !visitedRegions.IsVisited(bottom) {
		size += visitRegion(bottom, currentDepth)
	}
	if regionDepths.Contains(left) && !visitedRegions.IsVisited(left) {
		size += visitRegion(left, currentDepth)
	}

	//2. if point is not selected size++
	if !visitedRegions.IsVisited(current) {
		size++
		visitedRegions.VisitRegion(current)
	}

	//3. return size
	return size
}

func Run(lowPoints []Point, regions Matrix) (total int) {
	visitedRegions = MakeVisitedRegions(len(regions), len(regions[0]))
	regionDepths = regions

	basinSizes := make([]int, 0)

	for _, point := range lowPoints {
		basinSizes = append(basinSizes, calculateBasinSize(point))
	}

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}
