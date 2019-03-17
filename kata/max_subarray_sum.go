package kata

// Find the greatest sum of consecutive subarray
func MaximumSubarraySum(numbers []int) int {
	max := 0
	for i := 0; i < len(numbers); i++ {
		sum := 0
		for y := i; y < len(numbers); y++ {
			sum = sum + numbers[y]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
