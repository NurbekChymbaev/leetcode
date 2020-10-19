package _058_length_of_last_word

import "testing"

type LengthOfLastWordsCase struct {
	s        string
	expected int
}

var cases = []LengthOfLastWordsCase{
	{
		"Hello World",
		5,
	},

	{
		"",
		0,
	},

	{
		"Hello",
		5,
	},
}

func TestLengthOfLastWord(t *testing.T) {
	for caseNum, caseData := range cases {
		result := lengthOfLastWord(caseData.s)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
