package _349_intersection_of_two_arrays

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {

	var res []int

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := range nums1 {
		if has(nums2, nums1[i]) && !has(res, nums1[i]) {
			res = append(res, nums1[i])
		}
	}

	return res
}

func has(nums []int, target int) bool {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return true
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	return false
}
