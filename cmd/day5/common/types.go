package common

type Grid [][]int

type Point struct {
	X int
	Y int
}

func MakeGrid(n int) (grid Grid) {
	grid = make(Grid, n)

	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	return grid
}
