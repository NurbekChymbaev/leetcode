package _480_running_sum_of_1d_array

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		result[i] = sum
	}
	return result
}
