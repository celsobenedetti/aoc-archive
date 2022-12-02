package common

import (
	"fmt"
)

type Node struct {
	To     int
	Weight int
}

type Graph interface {
	CreateGraph()
	AddEdge(from, to, weight int)
	GetAdjacentNodes(from int) []Node
	PrintGraph()
}

type AdJList struct {
	List map[int][]Node
}

func (g *AdJList) CreateGraph() {
	g.List = make(map[int][]Node)
}

func (g *AdJList) AddEdge(from, to, weight int) {
	g.List[from] = append(g.List[from], Node{To: to, Weight: weight})
}

func (g *AdJList) GetAdjacentNodes(from int) []Node {
	return g.List[from]
}

func (g *AdJList) GetNumberOfVertices() int {
	return len(g.List)
}

func (g *AdJList) PrintGraph() {
	for from, vertices := range g.List {
		for _, edge := range vertices {
			fmt.Printf("[%d,%d|%d] ", from, edge.To, edge.Weight)
		}
		fmt.Println("")
	}
}
