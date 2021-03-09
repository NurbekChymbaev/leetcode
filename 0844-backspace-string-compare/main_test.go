package _844_backspace_string_compare

import "testing"

type Case struct {
	S      string
	T      string
	answer bool
}

var cases = []Case{
	{
		"ab##",
		"c#d#",
		true,
	},
}

func TestBackspaceCompare(t *testing.T) {

	for num, data := range cases {
		result := backspaceCompare(data.S, data.T)
		if data.answer != result {
			t.Error("case:", num,
				"\n\tresult:  ", result,
				"\n\texpected:", data.answer)
		}
	}
}
