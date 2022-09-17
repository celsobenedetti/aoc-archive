package main

import (
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `199
200
208
210
200
207
240
269
260
263
`
}

func parseInput(input string) []int {
	lines := strings.Split(getInput(), "\n")

	var measurements []int

	for _, measurement := range lines {

		depth, err := strconv.Atoi(strings.Trim(measurement, " "))

		if err != nil {
			log.Fatal(measurement, err, " String conversion failed")
		}
		measurements = append(measurements, depth)
	}
	return measurements
}

func main() {

	log.Printf("Part 1: %d", doPart1())

	log.Printf("Part 2: %d", doPart2())
}
