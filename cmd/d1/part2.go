package main

func doPart2() int {

	result := 0
	lines := parseInput(getInput())

	prev := lines[0] + lines[1] + lines[2]

	for i := 2; i < len(lines)-1; i++ {

		sum := 0

		for j := i - 1; j < i+2; j++ {
			sum += lines[j]
		}

		if sum > prev {
			result++
		}

		prev = sum
	}

	return result
}
