package recursion

import (
	"fmt"
	"strings"
)

// recursion for permutation
// such as Permutaion({1, 2, 3}) = {123, 132, 231, 213, 312, 321} etc.
func Permutation(input string) []string {
	ret := make([]string, 0)

	if len(input) == 1 {
		ret = append(ret, input)
		return ret
	}

	for idx, v := range input {
		var left string = ""

		if idx == 0 {
			left = input[1:]
		} else {
			pre := input[:idx]
			post := input[idx+1:]
			left = strings.Join([]string{pre, post}, "")
		}
		leftPermutaions := Permutation(left)
		for _, per := range leftPermutaions {
			oneUnit := strings.Join([]string{string(v), per}, "")
			ret = append(ret, oneUnit)
		}
	}
	return ret
}

func TestPermutaion() {
	strs := Permutation("abcdefg")
	for _, v := range strs {
		fmt.Println(v)
	}
	fmt.Println("length: ", len(strs))
}
