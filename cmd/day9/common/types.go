package common

type Point struct {
	X int
	Y int
}

func (p Point) IsContained(matrix [][]int) bool {
	return p.X >= 0 && p.X < len(matrix[0]) && p.Y >= 0 && p.Y < len(matrix)
}

type Matrix [][]int

func (m Matrix) Get(point Point) int {
	return m[point.Y][point.X]
}

func (m Matrix) Contains(point Point) bool {
	return point.IsContained(m)
}
