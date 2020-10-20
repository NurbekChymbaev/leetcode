package _221_split_a_string_in_balanced_strings

func balancedStringSplit(s string) int {

	count := 0
	rCount := 0
	lCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			lCount++
		} else {
			rCount++
		}
		if lCount == rCount {
			count++
			rCount = 0
			lCount = 0
		}
	}

	return count
}
