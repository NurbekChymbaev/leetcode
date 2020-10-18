package _242_valid_anagram

import "testing"

type AnagramCase struct {
	s        string
	t        string
	expected bool
}

var cases = []AnagramCase{
	{
		"anagram",
		"nagaram",
		true,
	},

	{
		"rat",
		"car",
		false,
	},

	{
		"foobar",
		"bar",
		false,
	},
}

func TestValidAnagram(t *testing.T) {

	for caseNum, caseData := range cases {
		result := isAnagram(caseData.s, caseData.t)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
