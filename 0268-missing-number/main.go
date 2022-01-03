package _0268_missing_number

func missingNumber(nums []int) int {
	var sum int
	var max int
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if len(nums) > max {
		max = len(nums)
	}
	return (max * (max + 1) / 2) - sum
}
