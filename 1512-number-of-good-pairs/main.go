package _512_number_of_good_pairs

func numIdenticalPairs(nums []int) int {

	pairs := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pairs++
			}
		}
	}

	return pairs
}
