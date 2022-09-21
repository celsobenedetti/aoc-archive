package common

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

func (crabMap CrabMap) GetConstantFuelUsageToPosition(position int) (total int) {
	for k, v := range crabMap {
		diff := math.Abs(float64(position - k))
		total += int(diff) * v
	}
	return
}

func (crabMap CrabMap) GetArithmeticFuelUsageToPosition(position int) (total int) {
	for k, v := range crabMap {
		diff := math.Abs(float64(position - k))
		usageInDiff := 0

		for i := 1; i <= int(diff); i++ {
			usageInDiff += i
		}

		total += usageInDiff * v
	}
	return
}
