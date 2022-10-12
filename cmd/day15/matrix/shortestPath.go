package matrix

func (m Matrix) shortestPath(source, target Point) (path []Node) {
	heap := CreateHeap(2000)

	m.At(source).Dist = 0
	heap.Insert(m.At(source))

	for heap.length > 0 {

		currNode := heap.Pop()

		for _, edge := range m.getAdjacentNodes(currNode.Position) {
			dist := currNode.Dist + edge.Weight

			if dist < edge.Dist {
				edge.Prev = currNode
				edge.Dist = dist
				heap.Insert(edge)
			}
		}
	}

	curr := m.At(target)

	for curr.Prev != nil {
		path = append(path, *curr)
		curr = curr.Prev
	}

	return
}

func (m Matrix) CalculateRisk() (total int) {
	nMatrix := len(m) - 1
	path := m.shortestPath(Point{I: 0, J: 0}, Point{I: nMatrix, J: nMatrix})

	for _, point := range path {
		total += point.Weight
	}
	return
}
