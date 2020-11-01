package _283_move_zeroes

func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[pos]
			nums[pos] = temp
			pos++
		}
	}
}
