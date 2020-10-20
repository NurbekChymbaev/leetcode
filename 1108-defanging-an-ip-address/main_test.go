package _058_length_of_last_word

import "testing"

type Case struct {
	s        string
	expected string
}

var cases = []Case{
	{
		"1.1.1.1",
		"1[.]1[.]1[.]1",
	},

	{
		"255.100.50.0",
		"255[.]100[.]50[.]0",
	},
}

func TestDefangIPaddr(t *testing.T) {
	for caseNum, caseData := range cases {
		result := defangIPaddr(caseData.s)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
