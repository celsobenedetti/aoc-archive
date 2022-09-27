package common

import (
	"strconv"
	"strings"
)

func ParseInput(input string) (matrix [][]int) {
	matrix = make([][]int, 0)

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}
		matrix = append(matrix, make([]int, 0))

		chars := strings.Split(line, "")

		for _, char := range chars {
			n, err := strconv.Atoi(char)
			if err != nil {
				return
			}
			matrix[i] = append(matrix[i], n)
		}
	}
	return
}
