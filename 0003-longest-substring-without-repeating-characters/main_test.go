package _003_longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbbl", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("", lengthOfLongestSubstring(""))
	fmt.Println("tmmzuxt", lengthOfLongestSubstring("tmmzuxt"))
}
