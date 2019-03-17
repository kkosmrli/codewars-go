package kata

import (
	"strings"
)

//https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/solutions/go
//Find the first chars that is not repeated in a given string
func FirstNonRepeating(str string) string {
	occ := map[string]int{}
	for _, s := range str {
		ls := strings.ToLower(string(s))
		occ[ls]++
	}
	for _, s := range str {
		ls := strings.ToLower(string(s))
		if occ[ls] == 1 {
			return string(s)
		}
	}
	return ""
}
