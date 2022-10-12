package matrix

func (m Matrix) shortestPath(source, target Point) (path []Node) {

	m.At(source).Dist = 0

	for m.hasUnvisitedNode() {
		currNode, currPoint := m.getNearestUnseenNode()
		currNode.Seen = true

		for _, edge := range m.getAdjacentNodes(currPoint) {
			dist := currNode.Dist + edge.Weight

			if dist < edge.Dist {
				edge.Prev = currNode
				edge.Dist = dist
			}
		}
	}

	curr := m.At(target)

	for curr.Prev != nil {
		path = append(path, *curr)
		curr = curr.Prev
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
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
