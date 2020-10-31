package _392_is_subsequence

func isSubsequence(s string, t string) bool {
	j := 0
	occur := 0
	for i := range s {
		for k := range t[j:] {
			if s[i] == t[k+j] {
				j = k + j + 1
				occur++
				break
			}
		}
	}
	return len(s) == occur
}
