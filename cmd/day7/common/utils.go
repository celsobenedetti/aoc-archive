package common

import (
	"log"
	"strconv"
	"strings"
)

func ParseStringsToInts(s []string) (numbers []int) {
	for _, v := range s {
		v = strings.Trim(v, "\n")
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Invalid input number")
		}
		numbers = append(numbers, number)
	}

	return
}
