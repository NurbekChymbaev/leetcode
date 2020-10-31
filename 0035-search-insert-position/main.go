package _035_search_insert_position

func searchInsert(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
