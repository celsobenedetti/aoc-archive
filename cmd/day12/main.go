package main

import (
	"fmt"

	input "github.com/celso-patiri/aoc/cmd/common"
	"github.com/celso-patiri/aoc/cmd/day12/graphs"
	graphInput "github.com/celso-patiri/aoc/cmd/day12/input"
)

var debugInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var debugInput2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var debugInput3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

var inputFile = "./cmd/day12/input.txt"

func main() {
	inputString := input.ReadInputFile(inputFile)

	var g graphs.Graph
	g = graphInput.ParseInput(inputString)

	pathsPart1 := graphs.DFS(g, "start", "end", 1)
	pathsPart2 := graphs.DFS(g, "start", "end", 2)

	fmt.Println("Part 1: ", len(pathsPart1))
	fmt.Println("Part 2: ", len(pathsPart2))
}
