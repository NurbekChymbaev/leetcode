package _047_remove_all_adjacent_duplicates_in_string

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	if removeDuplicates("abbaca") != "ca" {
		t.Error("Not match result")
	}
}
