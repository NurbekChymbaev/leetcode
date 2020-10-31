package _003_longest_substring_without_repeating_characters

// Brute force solution

//func lengthOfLongestSubstring(s string) int {
//	max := 0
//	for i := range s {
//		count := 0
//		storage := make(map[uint8]bool)
//		for j := range s[i:] {
//			index := i + j
//			if _, ok := storage[s[index]]; ok {
//				break
//			}
//			storage[s[index]] = true
//			count++
//		}
//		if max < count {
//			max = count
//		}
//	}
//	return max
//}

// Dynamic solution
func lengthOfLongestSubstring(s string) int {
	longest := 0
	curr := 0
	// храним символ и последнюю позицию где она в строке
	processed := make(map[uint8]int, 26)

	for i := range s {
		// если символ еще не встречался в строке и длина подстроки длиннее максимума то берем в процесс
		if _, ok := processed[s[i]]; !ok || processed[s[i]] < (i-curr) {
			curr++
			processed[s[i]] = i
			if curr > longest {
				longest = curr
			}
		} else {
			curr = i - processed[s[i]]
			processed[s[i]] = i
		}
	}

	return longest
}
