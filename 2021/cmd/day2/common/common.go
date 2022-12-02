package common

import (
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	Command string
	Arg     int
}

const (
	Up      string = "up"
	Down    string = "down"
	Forward string = "forward"
)

func ParseInstruction(instruction string) Instruction {
	tuple := strings.Split(instruction, " ")
	command, value := tuple[0], tuple[1]

	arg, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("String convertion failed")
	}

	return Instruction{command, arg}
}
