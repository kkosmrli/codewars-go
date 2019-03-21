package kata

//checks a given string for valid parantheses https://www.codewars.com/kata/52774a314c2333f0a7000688/solutions/go
func ValidParentheses(parens string) bool {
	stack := 0
	for _, s := range parens {
		if s == '(' {
			stack++
		} else if s == ')' {
			stack--
		}

		if stack < 0 {
			return false
		}
	}
	return stack == 0
}
