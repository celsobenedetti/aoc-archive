package input

import (
	"strings"

	"github.com/celso-patiri/aoc/cmd/day12/graphs"
)

func ParseInput(input string) (g graphs.Graph) {
	g = &graphs.CaveGraph{}
	g.CreateGraph()

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		caves := strings.Split(line, "-")
		g.AddEdge(caves[0], caves[1])
	}

	return g
}
