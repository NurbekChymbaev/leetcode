package _020_valid_parentheses

import "testing"

type ValidParenthesesCase struct {
	s        string
	expected bool
}

var cases = []ValidParenthesesCase{
	{
		"()",
		true,
	},
	{
		"()[]{}",
		true,
	},

	{
		"(]",
		false,
	},

	{
		"([)]",
		false,
	},

	{
		"{[]}",
		true,
	},
	{
		"()[]{}",
		true,
	},
}

func TestValidParentheses(t *testing.T) {

	for caseNum, caseData := range cases {

		var result = isValid(caseData.s)

		if result != caseData.expected {
			t.Error("case:", caseNum,
				"\n\tstring: ", caseData.s,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
