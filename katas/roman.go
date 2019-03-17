package katas

// Decode is a Roman numbers decoder
func Decode(roman string) int {
	l := strings.Split(roman, "")
	n := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}

	val := 0
	for i := 0; i < len(l); i++ {

		// Prevent lookup on last number
		if i == len(l)-1 {
			val = val + n[l[i]]
			return val
		}
		if n[l[i]] < n[l[i+1]] {
			val = val - n[l[i]]
		} else {
			val = val + n[l[i]]
		}
	}
	return val
}
