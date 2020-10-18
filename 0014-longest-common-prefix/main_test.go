package _014_longest_common_prefix

import (
	"testing"
)

type LCPCase struct {
	strs     []string
	expected string
}

var cases = []LCPCase{
	{
		[]string{"flower", "flow", "flight"},
		"fl",
	},
	{
		[]string{"dog", "racecar", "car"},
		"",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for caseNum, caseData := range cases {
		result := longestCommonPrefix(caseData.strs)
		if result != caseData.expected {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
