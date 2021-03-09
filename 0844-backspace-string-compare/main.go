package _844_backspace_string_compare

import (
	"strings"
)

func backspaceCompare(S string, T string) bool {
	return backspaceString(S) == backspaceString(T)
}

func backspaceString(S string) string {

	if len(S) == 0 {
		return S
	}

	var stack []string

	if S[0] != '#' {
		stack = append(stack, string(S[0]))
	}

	for _, char := range S[1:] {
		n := len(stack) - 1
		if char == '#' {
			if n >= 0 {
				stack = stack[:n]
			}
		} else {
			stack = append(stack, string(char))
		}
	}
	return strings.Join(stack, "")
}
