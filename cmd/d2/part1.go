package main

func doPart1() Point {
	position := Point{0, 0}

	for _, line := range input {
		instruction := parseInstruction(line)

		position = interpretInstruction(position, instruction)
	}

	return position
}
