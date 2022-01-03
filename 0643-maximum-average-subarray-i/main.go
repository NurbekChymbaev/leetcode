package _0643_maximum_average_subarray_i

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	var avr = float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		var tempAvr = float64(sum) / float64(k)
		if tempAvr > avr {
			avr = tempAvr
		}
	}

	return avr
}
