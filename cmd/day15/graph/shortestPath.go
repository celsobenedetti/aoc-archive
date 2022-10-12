package graph

import (
	"github.com/celso-patiri/aoc/cmd/common"
	"math"
)

type Seen []bool

func (s Seen) hasUnvisitedNode(dists []int) bool {
	for i, seen := range s {
		if !seen && dists[i] < math.MaxInt {
			return true
		}
	}
	return false
}

func (s Seen) nearestUnvisited(dists []int) int {
	idx := -1
	min := math.MaxInt
	for i, seen := range s {
		if !seen && dists[i] < min {
			min = dists[i]
			idx = i
		}
	}
	return idx
}

func ShortestPath(g common.AdJList, start, target int) (path []int) {
	nVertices := g.GetNumberOfVertices()

	if target >= nVertices {
		return []int{}
	}

	seen := make(Seen, nVertices)
	prev := make([]int, nVertices)
	dists := make([]int, nVertices)

	for i := 0; i < nVertices; i++ {
		prev[i] = -1
		dists[i] = math.MaxInt
		seen[i] = false
	}

	dists[start] = 0

	for seen.hasUnvisitedNode(dists) {
		curr := seen.nearestUnvisited(dists)
		seen[curr] = true

		for _, edge := range g.GetAdjacentNodes(curr) {
			dist := dists[curr] + edge.Weight

			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	curr := prev[target]

	for prev[curr] != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	path = append(path, start)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func CalculateRiskGraph(g common.AdJList, riskArray []int) (total int) {
	path := ShortestPath(g, 0, len(riskArray)-1)

	for _, position := range path {
		total += riskArray[position]
	}
	return total

}
