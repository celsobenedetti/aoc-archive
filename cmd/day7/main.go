package main

import (
	"fmt"
	"strings"

	file "github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day7/common"
)

var inputFile = "input.txt"
var debugInput = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

var minFuelConstant = 1<<31 - 1   //large number
var minFuelArithmetic = 1<<31 - 1 //large number

var maxDistance int
var crabMap common.CrabMap

func main() {
	inputString := file.ReadInputFile(inputFile)
	input := common.ParseStringsToInts(strings.Split(inputString, ","))

	crabMap, maxDistance = common.MakeCrabMap(input)

	for i := 0; i <= maxDistance; i++ {
		fuelUsageConstant := crabMap.GetConstantFuelUsageToPosition(i)
		fuelUsageArithmetic := crabMap.GetArithmeticFuelUsageToPosition(i)

		if fuelUsageConstant < minFuelConstant {
			minFuelConstant = fuelUsageConstant
		}

		if fuelUsageArithmetic < minFuelArithmetic {
			minFuelArithmetic = fuelUsageArithmetic
		}
	}

	fmt.Println("Part 1: ", minFuelConstant)
	fmt.Println("Part 2: ", minFuelArithmetic)
}
