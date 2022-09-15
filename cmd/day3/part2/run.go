package part2

import (
	"strconv"
	"strings"

	"github.com/celso-patiri/aoc/cmd/day3/common"
)

func stringToArr(s string) []string {
	return strings.Split(s, "")
}

func evaluateBitFrequency(input []string) (result []string) {
	frequency := make([]int, len(input[0]))

	for _, line := range input {

		for i, char := range stringToArr(line) {
			if char == "1" {
				frequency[i]++
			} else {
				frequency[i]--
			}

		}
	}

	for _, value := range frequency {
		if value < 0 {
			result = append(result, "0")
		} else {
			result = append(result, "1")
		}
	}
	return result
}

func findRating(input []string, bit int, frequency []string, mostCommon bool) (remaining []string) {
	filter := frequency[bit]

	if len(input) == 1 {
		return input
	}

	for _, line := range input {
		chars := stringToArr(line)

		if (chars[bit] == filter && mostCommon) || (chars[bit] != filter && !mostCommon) {
			remaining = append(remaining, line)
		}
	}

	return findRating(remaining, bit+1, frequency, mostCommon)
}

func convertRate(rate string) int {
	rateNumber, err := strconv.ParseInt(rate, 2, 64)

	if err != nil {
		panic("Invalid rate rateNumber")
	}

	return int(rateNumber)
}

func Run() int {
	input := common.DebugInput

	frequency := evaluateBitFrequency(input)

	oxigenRating := findRating(input, 0, frequency, true)[0]
	co2Rating := findRating(input, 0, frequency, false)[0]

	a := convertRate(oxigenRating)
	b := convertRate(co2Rating)

	return a * b
}
