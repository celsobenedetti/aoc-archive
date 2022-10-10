package graphs

import (
	"fmt"
	"strings"
)

type CaveType int

const (
	SMALL CaveType = iota
	BIG
)

func getCaveType(name string) CaveType {
	if strings.ToUpper(name) == name {
		return BIG
	}
	return SMALL
}

type CaveGraph struct {
	vertice map[string]*Adjacent
}

func (g *CaveGraph) CreateGraph() {
	g.vertice = map[string]*Adjacent{}
}

func (g *CaveGraph) AddEdge(from, to string) {
	if to == "start" || from == "end" {
		from, to = to, from
	}

	newNodeFrom := Node{Name: to, CaveType: getCaveType(to)}
	newNodeTo := Node{Name: from, CaveType: getCaveType(from)}

	adjacentFrom, ok := g.vertice[from]
	if !ok {
		g.vertice[from] = &Adjacent{newNodeFrom}
	} else {
		*adjacentFrom = append(*adjacentFrom, newNodeFrom)
	}

	if from == "start" || to == "end" {
		return
	}

	adjacentTo, ok := g.vertice[to]
	if !ok {
		g.vertice[to] = &Adjacent{newNodeTo}
	} else {
		*adjacentTo = append(*adjacentTo, newNodeTo)
	}

}

func (g *CaveGraph) GetAdjacentNodes(curr string) Adjacent {
	if _, ok := g.vertice[curr]; ok {
		return *(g.vertice[curr])
	}
	return Adjacent{}
}

func (g *CaveGraph) PrintGraph() {
	for from, adj := range g.vertice {
		for _, to := range *adj {
			fmt.Printf("[%s, %s] ", from, to.Name)
		}
		fmt.Println("")
	}
}
