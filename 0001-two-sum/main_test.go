package main

import (
	"testing"
)

type TwoSumCase struct {
	nums     []int
	target   int
	expected []int
}

var cases = []TwoSumCase{

	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},

	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},

	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for caseNum, caseData := range cases {
		result := twoSum(caseData.nums, caseData.target)
		if result[0] != caseData.expected[0] || result[1] != caseData.expected[1] {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
