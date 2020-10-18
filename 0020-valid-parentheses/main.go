package _020_valid_parentheses

func isValid(s string) bool {

	n := len(s)

	if n == 0 {
		return true
	}

	if n == 1 {
		return false
	}

	var pairs = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []string

	stack = append(stack, string(s[0]))

	for _, char := range s[1:] {
		n := len(stack) - 1
		if n >= 0 && stack[n] == string(pairs[char]) {
			stack = stack[:n]
		} else {
			stack = append(stack, string(char))
		}
	}
	return len(stack) == 0
}
