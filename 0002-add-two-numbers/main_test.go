package _002_add_two_numbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var l1 = &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	var l2 = &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}

	var l3 = addTwoNumbers(l1, l2)

	if l3.Val != 7 {
		t.Error("Invalid answer")
	}
}
