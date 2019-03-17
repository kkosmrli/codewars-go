package kata

import (
	"fmt"
)

// Josephus Permutation
func Josephus(items []interface{}, k int) []interface{} {
	res := []interface{}{}
	index := 0
	for len(items) > 0 {
		fmt.Printf("\nLen : %d", len(items))
		elim := k - 1 + index

		for elim >= len(items) {
			elim = elim - len(items)
		}
		fmt.Printf("\nEliminating: %d\n", elim)
		res = append(res, items[elim])
		items = append(items[:elim], items[elim+1:]...)
		index = elim
	}

	return res
}
