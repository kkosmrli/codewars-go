package kata

import (
	"strings"
)

func duplicate_count(s1 string) int {
	var res int
	counts := make(map[string]int)
	for _, s := range s1 {
		counts[strings.ToLower(string(s))]++

	}

	for _, count := range counts {
		if count > 1 {
			res++
		}
	}
	return res
}
