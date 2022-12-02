package common

import (
	"log"
	"strconv"
	"strings"
)

func coordsToInt(a, b string) (x, y int) {
	x, errx := strconv.Atoi(strings.Trim(a, " "))
	y, erry := strconv.Atoi(strings.Trim(b, " "))
	if errx != nil || erry != nil {
		log.Fatalln("Error: invalid string")
		return
	}
	return x, y
}

func ParsePoint(point string) (x, y int) {
	coords := strings.Split(point, ",")
	return coordsToInt(coords[0], coords[1])
}
