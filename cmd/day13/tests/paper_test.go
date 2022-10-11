package test

import (
	"reflect"
	"testing"

	"github.com/celso-patiri/aoc/cmd/day13/input"
	"github.com/celso-patiri/aoc/cmd/day13/paper"
)

var page paper.Paper
var instructions []paper.Instruction

var debugInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7`

func init() {
	page, instructions = input.ParseInput(debugInput)
}

func getHalf(n int) int {
	return (n - 1) / 2
}

func TestSplitY(t *testing.T) {
	nRows := len(page)

	left, right := page.SplitY(7)

	leftRows := len(left)

	if !reflect.DeepEqual(leftRows, getHalf(nRows)) {
		t.Fatalf("SplitX: expected %d rows, got %d", getHalf(nRows), leftRows)
	}

	rightRows := len(right)

	if !reflect.DeepEqual(rightRows, getHalf(nRows)) {
		t.Fatalf("SplitX: expected %d rows, got %d", getHalf(nRows), rightRows)
	}

}

func TestSplitX(t *testing.T) {
	nCols := len(page[0])

	left, right := page.SplitX(5)

	rightCols := len(right[0])
	leftCols := len(left[0])

	if !reflect.DeepEqual(leftCols, getHalf(nCols)) {
		t.Fatalf("SplitX: expected %d cols, got %d", getHalf(nCols), leftCols)
	}

	if !reflect.DeepEqual(rightCols, getHalf(nCols)) {
		t.Fatalf("SplitX: expected %d cols, got %d", getHalf(nCols), rightCols)
	}
}
