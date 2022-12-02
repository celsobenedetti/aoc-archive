package part1

import "github.com/celso-patiri/aoc/cmd/day6/common"

func Run(fishes []int) int {

	fishList := common.MakeFishList(fishes)

	for i := 0; i < 80; i++ {
		fishList = fishList.RunDayLifeClycle()
	}

	return fishList.GetLen()
}
