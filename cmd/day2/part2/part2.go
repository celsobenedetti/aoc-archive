package part2

import (
	c "github.com/celso-patiri/aoc/cmd/day2/common"
)

type Coordinate struct {
	X, Y, Aim int
}

func interpretInstruction(position Coordinate, instruction c.Instruction) Coordinate {

	amount := instruction.Arg

	if instruction.Command == c.Up {
		return Coordinate{position.X, position.Y, position.Aim - amount}
	}

	if instruction.Command == c.Down {
		return Coordinate{position.X, position.Y, position.Aim + amount}
	}

	if instruction.Command == c.Forward {
		return Coordinate{position.X + amount, position.Y + amount*position.Aim, position.Aim}
	}

	panic("Unknown instruction")
}

func Run() Coordinate {

	position := Coordinate{}

	for _, line := range c.Input {

		instruction := c.ParseInstruction(line)

		position = interpretInstruction(position, instruction)
	}

	return position
}
