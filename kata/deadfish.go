package kata

// Deadfish parser: https://www.codewars.com/kata/make-the-deadfish-swim/train/go
func Parse(data string) []int {
	value := 0
	res := []int{}
	for _, c := range data {
		if c == 'i' {
			value++
		} else if c == 'd' {
			value--
		} else if c == 's' {
			value *= value
		} else if c == 'o' {
			res = append(res, value)
		}
	}
	return res
}
