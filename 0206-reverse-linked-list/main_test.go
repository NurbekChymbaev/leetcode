package _206_reverse_linked_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {

	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}

	head = reverseList(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
