package _014_longest_common_prefix

// Brute force solution
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	var pattern = strs[0]
	for _, str := range strs {
		if len(str) < len(pattern) {
			pattern = str
		}
	}

	var resultPrefix = ""
	var resultPrefixLength = len(resultPrefix)
	for i := range pattern {
		var prefix = pattern[0 : i+1]
		for _, str := range strs {
			if str[0:i+1] != prefix {
				prefix = ""
				break
			}
		}
		if resultPrefixLength < len(prefix) {
			resultPrefixLength = len(resultPrefix)
			resultPrefix = prefix
		}
	}

	return resultPrefix
}
