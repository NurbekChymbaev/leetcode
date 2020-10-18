package _507_reformat_date

import (
	"testing"
)

type ReformatDateCase struct {
	date     string
	expected string
}

var cases = []ReformatDateCase{
	{
		"20th Oct 2052",
		"2052-10-20",
	},
	{
		"6th Jun 1933",
		"1933-06-06",
	},
	{
		"26th May 1960",
		"1960-05-26",
	},
}

func TestReformatDate(t *testing.T) {

	for caseNum, caseData := range cases {
		result := reformatDate(caseData.date)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
