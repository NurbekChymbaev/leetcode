package main

func maxDepth(s string) int {

	cur := 0
	total := 0

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			cur++
		}

		if s[i] == ')' {
			if cur > total {
				total = cur
			}
			cur--
		}
	}

	return total
}
