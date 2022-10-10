package graphs

type Graph interface {
	CreateGraph()
	AddEdge(from, to string)
	GetAdjacentNodes(from string) Adjacent
	PrintGraph()
}
