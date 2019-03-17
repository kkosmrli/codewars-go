package kata

//InAscOder: https://www.codewars.com/kata/56b7f2f3f18876033f000307
func InAscOrder(numbers []int) bool {
	res := true
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			res = false
		}
	}
	return res
}
