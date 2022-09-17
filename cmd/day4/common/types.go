package common

type Row []int
type Board []Row

type BingoInput struct {
	Numbers []int
	Boards  []Board
}

type bingoItem struct {
	Board  int
	Row    int
	Column int
}

type CheckRow []bool
type CheckBoard []CheckRow

type mappedItems []bingoItem
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
	mappedItems = append(mappedItems, bingoItem{board, row, column})
	bingoMap[item] = mappedItems
}
