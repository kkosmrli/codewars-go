package kata

import (
	"fmt"
)

// ValidateSolution validates if a sudoku board is solved correctly
// https://www.codewars.com/kata/sudoku-solution-validator/train/go
func ValidateSolution(m [][]int) bool {

	// Horizontal and Vertical
	for i := 0; i < len(m); i++ {

		v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for j := 0; j < len(m); j++ {

			if m[i][j] == 0 || m[j][i] == 0 {
				return false
			}
			v = remove(v, m[i][j])
			h = remove(h, m[j][i])
		}
		if len(v) != 0 || len(h) != 0 {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for y := 0; y < 9; y += 3 {
			elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			for s := 0; s < 3; s++ {
				for x := 0; x < 3; x++ {
					elements = remove(elements, m[i+s][y+x])
				}
			}
			if len(elements) != 0 {
				fmt.Println(elements)
				return false
			}
		}
	}
	fmt.Println("")
	return true
}

func remove(l []int, e int) []int {
	for i, x := range l {
		if e == x {
			l = append(l[:i], l[i+1:]...)
		}
	}
	return l
}
