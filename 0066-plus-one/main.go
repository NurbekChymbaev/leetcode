package _0066_plus_one

func plusOne(digits []int) []int {
	var incr = 1
	var sum int
	for i := len(digits) - 1; i >= 0; i-- {
		sum = incr + digits[i]
		if sum > 9 {
			sum = 0
			incr = 1
		} else {
			incr = 0
		}
		digits[i] = sum
	}

	if sum == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
