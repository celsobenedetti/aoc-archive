package part2

import "github.com/celso-patiri/aoc/cmd/day6/common"

func Run(fishes []int) int {

	fishBuffer := common.MakeFishBuffer(fishes)

	for i := 0; i < 256; i++ {
		fishBuffer = fishBuffer.RunRoundLifeCycle()
	}

	return fishBuffer.Count
}
