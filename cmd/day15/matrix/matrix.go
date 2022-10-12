package matrix

import (
	"fmt"
)

type Node struct {
	Weight   int
	Prev     *Node
	Dist     int
	Position Point
}

type Matrix [][]Node

type Point struct {
	I int
	J int
}

func (m Matrix) At(p Point) *Node {
	return &m[p.I][p.J]
}

func (m Matrix) getAdjacentNodes(p Point) (adj []*Node) {
	i, j := p.I, p.J
	if i > 0 {
		top := m.At(Point{I: i - 1, J: j})
		adj = append(adj, top)
	}
	if i < len(m)-1 {
		bottom := m.At(Point{I: i + 1, J: j})
		adj = append(adj, bottom)
	}
	if j > 0 {
		left := m.At(Point{I: i, J: j - 1})
		adj = append(adj, left)
	}
	if j < len(m)-1 {
		right := m.At(Point{I: i, J: j + 1})
		adj = append(adj, right)
	}
	return
}

func (m Matrix) PrintMatrix() {
	for _, row := range m {
		for _, node := range row {
			fmt.Printf("%d ", node.Weight)
		}
		fmt.Println("")
	}
}
