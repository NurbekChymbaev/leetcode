package _058_length_of_last_word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)

	if len(words) == 0 {
		return 0
	}

	return len(words[len(words)-1])
}
