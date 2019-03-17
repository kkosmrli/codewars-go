package kata

import "errors"

// ClosingBracket returns the index of the closing bracket for a given opening bracket
func ClosingBracket(str string, idx uint) (uint, error) {
	if str[idx] != '(' {
		return 0, errors.New("Not an opening bracket")
	}
	stack := 1
	for i := idx + 1; i < uint(len(str)); i++ {
		if str[i] == '(' {
			stack++
		} else if str[i] == ')' {
			stack--
		}
		if stack == 0 {
			return i, nil
		}

	}
	return 0, nil
}
