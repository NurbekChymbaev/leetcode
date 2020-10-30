package _003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	max := 0

	for i := range s {
		count := 0
		storage := make(map[uint8]bool)

		for j := range s[i:] {
			index := i + j
			if _, ok := storage[s[index]]; ok {
				break
			}
			storage[s[index]] = true
			count++
		}

		if max < count {
			max = count
		}
	}

	return max
}
