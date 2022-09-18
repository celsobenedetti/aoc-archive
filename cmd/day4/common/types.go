package common

type Row []int
type Board []Row

type BingoInput struct {
	Numbers []int
	Boards  []Board
}

type BingoItem struct {
	Board  int
	Row    int
	Column int
}

type CheckRow []bool
type CheckBoard []CheckRow
type CheckMatrix []CheckBoard

type mappedItems []BingoItem
type BingoMap map[int]mappedItems

// Receiver functions

func (bingo *BingoInput) makeNewBoard() {
	bingo.Boards = append(bingo.Boards, make(Board, 0))
}

func (board *Board) makeNewRow() {
	*board = append(*board, make(Row, 0))
}

func (row *Row) addElement(n int) {
	*row = append(*row, n)
}

func (bingoMap BingoMap) addItem(item, board, row, column int) {
	mappedItems := bingoMap[item]
	mappedItems = append(mappedItems, BingoItem{board, row, column})
	bingoMap[item] = mappedItems
}

func (bingoCheck CheckMatrix) CheckElement(board, row, column int) {
	bingoCheck[board][row][column] = true
}

func (board CheckBoard) IsVictory(row, column int) bool {
	rowIsComplete, columnIsComplete := true, true

	for i := 0; i < 5; i++ {
		rowIsComplete = rowIsComplete && board[row][i]
		columnIsComplete = columnIsComplete && board[i][column]
	}

	return rowIsComplete || columnIsComplete
}

func (bingoCheck *CheckMatrix) makeNewboard() {
	*bingoCheck = append(*bingoCheck, MakeCheckBoard())
}
