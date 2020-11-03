package _347_top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {

	var frequent []int

	var numsMap = make(map[int]int, len(nums))

	for i := range nums {
		numsMap[nums[i]]++
	}

	for p := 0; p < k; p++ {
		maxKey := -1
		maxVal := -1
		for key, val := range numsMap {
			if maxVal < val {
				maxKey = key
				maxVal = val
			}
		}
		frequent = append(frequent, maxKey)
		delete(numsMap, maxKey)
	}

	return frequent
}
