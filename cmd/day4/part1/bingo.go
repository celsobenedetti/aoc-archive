package part1

import (
	"github.com/celso-patiri/aoc/cmd/day4/common"
)

func computeDraftedNumber(draft int) {}
func checkWin(row, col int)          {}

func Run() common.BingoInput {

	for _, number := range common.Input.Numbers {
		computeDraftedNumber(number)
	}

	return common.Input
}
