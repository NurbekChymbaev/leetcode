package main

import "testing"

type Case struct {
	s        string
	expected int
}

var cases = []Case{
	{
		"(1+(2*3)+((8)/4))+1",
		3,
	},
	{
		"(1)+((2))+(((3)))",
		3,
	},

	{
		"1+(2*3)/(2-1)",
		1,
	},
	{
		"1",
		0,
	},
}

func TestNumberOfGoodPairs(t *testing.T) {
	for caseNum, caseData := range cases {
		result := maxDepth(caseData.s)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
