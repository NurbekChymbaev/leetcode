package _480_running_sum_of_1d_array

import (
	"testing"
)

type ReformatDateCase struct {
	nums     []int
	expected []int
}

var cases = []ReformatDateCase{
	{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 6, 10},
	},
	{
		[]int{1, 1, 1, 1, 1},
		[]int{1, 2, 3, 4, 5},
	},
	{
		[]int{3, 1, 2, 10, 1},
		[]int{3, 4, 6, 16, 17},
	},
}

func TestRunningSumOf1dArray(t *testing.T) {

	for caseNum, caseData := range cases {
		result := runningSum(caseData.nums)
		if !Equal(result, caseData.expected) {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
