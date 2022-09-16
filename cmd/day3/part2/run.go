package part2

import (
	"strconv"
	"strings"

	"github.com/celso-patiri/aoc/cmd/day3/common"
)

func strToArr(s string) []string {
	return strings.Split(s, "")
}

func evaluateBitFrequency(input []string) (result []string) {
	frequency := make([]int, len(input[0]))

	for _, line := range input {

		for i, char := range strToArr(line) {
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

func findRating(input []string, bit int, mostCommon bool) (remaining []string) {
	if len(input) == 1 {
		return input
	}

	frequency := evaluateBitFrequency(input)
	filter := frequency[bit]

	for _, line := range input {
		chars := strToArr(line)

		if (chars[bit] == filter && mostCommon) || (chars[bit] != filter && !mostCommon) {
			remaining = append(remaining, line)
		}
	}

	return findRating(remaining, bit+1, mostCommon)
}

func convertRating(rate string) int {
	rateNumber, err := strconv.ParseInt(rate, 2, 64)

	if err != nil {
		panic("Invalid rate rateNumber")
	}

	return int(rateNumber)
}

func getRating(name string) int {
	var rating string

	if name == "oxigen" {
		rating = findRating(common.Input, 0, true)[0]
	} else {
		rating = findRating(common.Input, 0, false)[0]
	}

	return convertRating(rating)
}

func Run() int {

	oxigenRating := getRating("oxigen")
	co2Rating := getRating("co2")

	return oxigenRating * co2Rating
}
