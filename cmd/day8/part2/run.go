package part2

import "fmt"

func Run(input string) (result int) {
	codes, output := ParseInput(input)

	fmt.Println(codes)
	fmt.Println(output)

	return
}
