package part1

var minFuel = 1<<31 - 1 //large number
var maxDistance int     //smol number
var crabMap CrabMap

func Run(input []int) int {

	crabMap, maxDistance = MakeCrabMap(input)

	for i := 0; i <= maxDistance; i++ {
		fuelUsage := crabMap.GetFuelUsageToPosition(i)

		if fuelUsage < minFuel {
			minFuel = fuelUsage
		}
	}

	return minFuel
}
