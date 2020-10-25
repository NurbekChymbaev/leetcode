package _047_remove_all_adjacent_duplicates_in_string

func removeDuplicates(S string) string {
	var stack []rune

	for _, char := range S {
		if len(stack) > 0 && last(stack) == char {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return string(stack)
}

func last(stack []rune) rune {
	return stack[len(stack)-1]
}
