package _392_is_subsequence

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	if isSubsequence("abc", "ahbgdc") != true {
		t.Error("Incorrect response")
	}

	if isSubsequence("axc", "ahbgdc") != false {
		t.Error("Incorrect response")
	}
}
