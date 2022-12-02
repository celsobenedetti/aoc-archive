package octopuses

type Octopi struct {
	matrix [][]int
}

func (o Octopi) Size() int {
	return len(o.matrix)
}

func (o *Octopi) Make(input [][]int) {
	o.matrix = input
}

func (o *Octopi) Inc(i, j int) {
	o.matrix[i][j] = (o.matrix[i][j] + 1) % 10
}

func (o Octopi) Get(i, j int) int {
	return o.matrix[i][j]
}

func (o Octopi) ShouldVisitPosition(i, j int) bool {
	n := o.Size()
	return i >= 0 && i < n && j >= 0 && j < n
}
