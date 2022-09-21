package part1

import (
	"math"
)

type CrabMap map[int]int

func MakeCrabMap(input []int) (crabMap CrabMap, maxDistance int) {
	crabMap = make(CrabMap)
	for _, v := range input {
		if v > maxDistance {
			maxDistance = v
		}
		crabMap[v]++
	}
	return
}

func (crabMap CrabMap) GetFuelUsageToPosition(position int) (total int) {
	for k, v := range crabMap {
		diff := math.Abs(float64(position - k))
		total += int(diff) * v
	}
	return
}
