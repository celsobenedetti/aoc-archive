package common

import (
	"strings"
)

var Input = parseInput(debugInput)
var InputMap BingoMap = make(BingoMap)
var BingoCheck CheckMatrix = make(CheckMatrix, 0)

func parseInput(input string) (bingo BingoInput) {
	bingo.makeNewBoard()
	BingoCheck.makeNewboard()

	lines := strings.Split(input, "\n")
	bingo.Numbers = ParseInputNumbers(lines[1])

	boardIndex := 0
	currentBoardLine := 0

	for _, line := range lines[3 : len(lines)-1] {
		if line == "" {
			bingo.makeNewBoard()
			BingoCheck.makeNewboard()

			boardIndex++
			currentBoardLine = 0
		} else {
			bingo.Boards[boardIndex].makeNewRow()

			line := ClearWhiteSpaces(line)

			for i, v := range strings.Split(line, " ") {
				bingo.Boards[boardIndex][currentBoardLine].addElement(StrToInt(v))
				InputMap.addItem(StrToInt(v), boardIndex, currentBoardLine, i)
			}

			currentBoardLine++
		}
	}
	return bingo
}

var debugInput = `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`
