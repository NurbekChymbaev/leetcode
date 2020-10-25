package _136_single_number

func singleNumber(nums []int) int {

	var hash = make(map[int]int)

	for _, k := range nums {
		hash[k]++
	}

	for i, k := range hash {
		if k == 1 {
			return i
		}
	}

	return 0
}
