package _083_remove_duplicates_from_sorted_list

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {

	head := &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				2,
				nil,
			},
		},
	}

	deleteDuplicates(head)
}
