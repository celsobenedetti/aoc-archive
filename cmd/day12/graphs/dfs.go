package graphs

var paths []Path
var MAX_VISITS int

func DFS(g Graph, start, target string, maxVisits int) []Path {
	MAX_VISITS = maxVisits
	paths = make([]Path, 0)

	walk(g, start, target, make(Path, 0))

	return paths
}

func walk(g Graph, curr, target string, path Path) {
	if curr == target {
		paths = append(paths, path.makeCopy())
		return
	}

	if getCaveType(curr) == SMALL {
		timesVisited := path.countOccurences(curr)
		if timesVisited == MAX_VISITS || (path.hasDuplicates() && timesVisited > 0) {
			return
		}
	}

	//pre
	path = append(path, curr)

	//recurse
	for _, edge := range g.GetAdjacentNodes(curr) {
		walk(g, edge.Name, target, path)
	}
}
