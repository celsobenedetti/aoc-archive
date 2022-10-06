package input

import (
	"log"
	"strconv"
	"strings"
)

func ParseInput(input string) (octopi [][]int) {
	octopi = make([][]int, 0)

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		octopi = append(octopi, make([]int, 0))

		positions := strings.Split(line, "")
		for _, position := range positions {

			number, err := strconv.Atoi(position)
			if err != nil {
				log.Fatal("Invalid matrix int position")
			}

			octopi[i] = append(octopi[i], number)
		}

	}
	return
}
