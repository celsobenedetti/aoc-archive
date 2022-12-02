package polymer

import (
	"math"
	"strings"
)

type Polymer map[string]int

func (p Polymer) Copy() (copy Polymer) {
	copy = Polymer{}
	for k, v := range p {
		copy[k] = v
	}
	return
}

func mapOcurrences(s string, polymer Polymer, ocurrences map[string]int) {
	charArr := strings.Split(s, "")

	for i := 0; i < len(charArr)-1; i++ {
		pattern := charArr[i] + charArr[i+1]
		polymer[pattern]++
		ocurrences[charArr[i]]++
	}
	ocurrences[charArr[len(charArr)-1]]++

}

func BuildPolymer(base string, rules map[string]string, steps int) int {
	charOccurrences := map[string]int{}
	polymer := Polymer{}

	mapOcurrences(base, polymer, charOccurrences)

	for step := 0; step < steps; step++ {
		currPolymer := Polymer{}

		for k, v := range polymer {

			newChar, hasRule := rules[k]

			if !hasRule {
				currPolymer[k]++
				continue
			}

			chars := strings.Split(k, "")
			firstPattern := chars[0] + newChar
			secondPattern := newChar + chars[1]

			currPolymer[firstPattern] += v
			currPolymer[secondPattern] += v
			charOccurrences[newChar] += v
		}

		polymer = currPolymer.Copy()
	}

	max, min := 0, math.MaxInt

	for _, v := range charOccurrences {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return (max - min)
}
