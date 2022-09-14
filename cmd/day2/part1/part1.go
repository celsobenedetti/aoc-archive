package part1

import (
	c "github.com/celso-patiri/aoc/cmd/day2/common"
)

type Point struct {
	X, Y int
}

func interpretInstruction(position Point, instruction c.Instruction) Point {

	if instruction.Command == c.Up {
		return Point{X: position.X, Y: position.Y - instruction.Arg}
	}

	if instruction.Command == c.Down {
		return Point{X: position.X, Y: position.Y + instruction.Arg}
	}

	if instruction.Command == c.Forward {
		return Point{X: position.X + instruction.Arg, Y: position.Y}
	}

	panic("Unknown instruction")
}

func Run() Point {
	position := Point{X: 0, Y: 0}

	for _, line := range c.Input {
		instruction := c.ParseInstruction(line)

		position = interpretInstruction(position, instruction)
	}

	return position
}
