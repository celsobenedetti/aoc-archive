package part1

import (
	"github.com/celso-patiri/aoc/cmd/day4/common"
)

func computeDraftedNumber(draft int) int {
	occurences := common.InputMap[draft]

	for _, ocurrence := range occurences {
		board, i, j := ocurrence.Board, ocurrence.Row, ocurrence.Column
		common.BingoCheck.CheckElement(board, i, j)

		if common.BingoCheck[board].IsVictory(i, j) {
			return board
		}
	}

	return -1
}

func calculateWin(boardIndex int) int {
	board := common.Input.Boards[boardIndex]
	checkBoard := common.BingoCheck[boardIndex]

	uncheckedSum := 0

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			if !checkBoard[i][j] {
				uncheckedSum += board[i][j]
			}
		}
	}

	return uncheckedSum
}

func Run() int {

	for _, number := range common.Input.Numbers {
		boardIndex := computeDraftedNumber(number)

		if boardIndex >= 0 {
			return calculateWin(boardIndex) * number
		}
	}

	return -1
}
