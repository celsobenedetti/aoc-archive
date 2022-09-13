package main

func doPart1() int {
	measurements := parseInput(getInput())

	result := 0

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			result++
		}
	}

	return result
}
