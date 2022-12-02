package part1

import (
	"github.com/celso-patiri/aoc/cmd/day3/common"
	"strconv"
	"strings"
)

func splitString(s string) []string {
	return strings.Split(s, "")
}

func getBase(inputLine string) []int {
	chars := splitString(inputLine)

	result := []int{}

	for range chars {
		result = append(result, 0)
	}

	return result
}

func evaluateLine(inputLine string, base *[]int) {
	chars := splitString(inputLine)

	for i, char := range chars {
		if char == "1" {
			(*base)[i]++
		} else {
			(*base)[i]--
		}
	}
}

func calculateRates(base []int) (string, string) {
	var gamma strings.Builder
	var epsilon strings.Builder

	for _, n := range base {
		if n > 0 {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}

	return gamma.String(), epsilon.String()
}

func convertRate(rate string) int {
	rateNumber, err := strconv.ParseInt(rate, 2, 64)

	if err != nil {
		panic("Invalid rate rateNumber")
	}

	return int(rateNumber)
}

func calculatePower(gamma string, epsilon string) int {
	gammaRate := convertRate(gamma)
	epsilonRate := convertRate(epsilon)

	return gammaRate * epsilonRate
}

func Run() int {

	base := getBase(common.Input[0])

	for _, line := range common.Input {
		evaluateLine(line, &base)
	}

	gamma, epsilon := calculateRates(base)

	return calculatePower(gamma, epsilon)
}
