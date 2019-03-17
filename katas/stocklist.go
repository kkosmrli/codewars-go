package katas

import (
	"fmt"
	"strconv"
	"strings"
)

// StockList kata: https://www.codewars.com/kata/help-the-bookseller/train/go
func StockList(listArt []string, listCat []string) string {
	if len(listCat) == 0 || len(listArt) == 0 {
		return ""
	}
	res := map[string]int{}
	out := ""
	for _, cat := range listCat {
		res[cat] = 0
		for _, book := range listArt {
			a := strings.Fields(book)
			if string(book[0]) == cat {
				v, _ := strconv.Atoi(a[1])
				res[cat] += v
			}
		}
		out = out + fmt.Sprintf("(%s : %d) - ", cat, res[cat])
	}
	return strings.Trim(out, " - ")
}
