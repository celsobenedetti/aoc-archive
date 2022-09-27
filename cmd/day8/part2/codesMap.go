package part2

type CodeLengthMap map[int][]string

func (codeMap CodeLengthMap) mapCode(input string) {
	strLen := len(input)
	if strLen == 2 || strLen == 3 || strLen == 4 || strLen == 7 {
		codeMap[strLen] = []string{input}
	}
}
