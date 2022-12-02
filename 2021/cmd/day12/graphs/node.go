package graphs

type Node struct {
	Name     string
	CaveType CaveType
}

type Adjacent []Node
type Path []string

func (a *Adjacent) findEdge(to string) (idx int, node *Node) {
	for i, v := range *a {
		if v.Name == to {
			return i, &(*a)[i]
		}
	}
	return -1, nil
}

func (p Path) hasDuplicates() bool {
	for i, v := range p {
		if getCaveType(v) == SMALL {
			for j := i + 1; j < len(p); j++ {
				if p[j] == v {
					return true
				}
			}
		}
	}
	return false
}

func (p Path) countOccurences(name string) (total int) {
	for _, v := range p {
		if v == name {
			total++
		}
	}
	return total
}

func (p Path) makeCopy() Path {
	return append(make(Path, 0), p...)
}
