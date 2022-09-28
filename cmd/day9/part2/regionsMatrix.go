package part2

type VisitedRegions struct {
	regions [][]bool
}

func (s *VisitedRegions) VisitRegion(point Point) {
	i, j := point.Y, point.X
	s.regions[i][j] = true
}

func (s *VisitedRegions) IsVisited(point Point) bool {
	return s.regions[point.Y][point.X]
}

func MakeVisitedRegions(rows, columns int) (selected VisitedRegions) {
	selected = VisitedRegions{regions: make([][]bool, rows)}

	for j := 0; j < rows; j++ {
		selected.regions[j] = make([]bool, columns)
	}

	return
}
