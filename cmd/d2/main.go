package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Instruction struct {
	command string
	arg     int
}

const (
	up      string = "up"
	down    string = "down"
	forward string = "forward"
)

func parseInstruction(instruction string) Instruction {
	tuple := strings.Split(instruction, " ")
	command, value := tuple[0], tuple[1]

	arg, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("String convertion failed")
	}

	return Instruction{command, arg}
}

func interpretInstruction(position Point, instruction Instruction) Point {

	if instruction.command == up {
		return Point{position.x, position.y - instruction.arg}
	}

	if instruction.command == down {
		return Point{position.x, position.y + instruction.arg}
	}

	if instruction.command == forward {
		return Point{position.x + instruction.arg, position.y}
	}

	panic("Unknown instruction")
}

func main() {
	position1 := doPart1()
	fmt.Printf("Part 1: %d", position1.x*position1.y)
}
