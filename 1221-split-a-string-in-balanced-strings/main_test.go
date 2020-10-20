package _221_split_a_string_in_balanced_strings

import (
	"testing"
)

type SplitStringCase struct {
	s        string
	expected int
}

var cases = []SplitStringCase{
	{
		"RLRRLLRLRL",
		4,
	},
	{
		"RLLLLRRRLR",
		3,
	},
	{
		"LLLLRRRR",
		1,
	},
	{
		"RLRRRLLRLL",
		2,
	},
}

func TestBalancedStringSplit(t *testing.T) {

	for caseNum, caseData := range cases {
		result := balancedStringSplit(caseData.s)
		if result != caseData.expected {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
