package _512_number_of_good_pairs

import "testing"

type NumberOfGoodPairs struct {
	nums     []int
	expected int
}

var cases = []NumberOfGoodPairs{
	{
		[]int{1, 2, 3, 1, 1, 3},
		4,
	},
	{
		[]int{1, 1, 1, 1},
		6,
	},

	{
		[]int{1, 2, 3},
		0,
	},
}

func TestNumberOfGoodPairs(t *testing.T) {
	for caseNum, caseData := range cases {
		result := numIdenticalPairs(caseData.nums)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
