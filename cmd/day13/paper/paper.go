package paper

import "fmt"

type Paper [][]bool
type Axis string

type Instruction struct {
	Axis     Axis
	Position int
}

const (
	X Axis = "x"
	Y Axis = "y"
)

func (p Paper) SplitX(column int) (left, right Paper) {
	for i := 0; i < len(p); i++ {
		left = append(left, make([]bool, 0))
		right = append(right, make([]bool, 0))

		for j := 0; j < column; j++ {
			left[i] = append(left[i], p[i][j])
		}

		for j := column + 1; j < len(p[i]); j++ {
			right[i] = append(right[i], p[i][j])
		}
	}
	return
}

func (p Paper) SplitY(row int) (top, bottom Paper) {
	for i := 0; i < row; i++ {
		top = append(top, p[i])
	}
	for i := row + 1; i < len(p); i++ {
		bottom = append(bottom, p[i])
	}
	return
}

func (p Paper) ReverseX() (rev Paper) {
	for i, row := range p {
		rev = append(rev, make([]bool, 0))

		for j := len(row) - 1; j >= 0; j-- {
			rev[i] = append(rev[i], row[j])
		}
	}
	return
}

func (p Paper) ReverseY() (rev Paper) {
	for i := len(p) - 1; i >= 0; i-- {
		rev = append(rev, p[i])
	}
	return
}

func (p Paper) foldPaper(page Paper) (fold Paper) {
	if len(p) != len(page) || len(p[0]) != len(page[0]) {
		panic("foldPaper: pages have different sizes")
	}

	for i := 0; i < len(p); i++ {
		fold = append(fold, make([]bool, 0))

		for j := 0; j < len(p[i]); j++ {
			fold[i] = append(fold[i], p[i][j] || page[i][j])
		}
	}
	return
}

func (p Paper) InterpretFold(i Instruction) Paper {
	if i.Axis == X {
		left, right := p.SplitX(i.Position)
		right = right.ReverseX()
		return left.foldPaper(right)
	}

	top, bottom := p.SplitY(i.Position)
	bottom = bottom.ReverseY()
	return top.foldPaper(bottom)
}

func (p Paper) CountDots() (count int) {
	for _, row := range p {
		for _, dot := range row {
			if dot {
				count++
			}

		}
	}
	return
}

func (p Paper) PrintPaper() {
	for _, row := range p {
		for _, dot := range row {
			if dot {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println("")
	}
}
