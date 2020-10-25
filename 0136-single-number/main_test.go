package _136_single_number

import (
	"testing"
)

type Case struct {
	nums     []int
	expected int
}

var cases = []Case{
	{
		[]int{2, 2, 1},
		1,
	},
	{
		[]int{4, 1, 2, 1, 2},
		4,
	},
	{
		[]int{-1, -1, -2},
		-2,
	},
}

func TestSingleNumber(t *testing.T) {
	for num, data := range cases {
		result := singleNumber(data.nums)
		if data.expected != result {
			t.Error("case:", num,
				"\n\tresult:  ", result,
				"\n\texpected:", data.expected)
		}
	}
}
