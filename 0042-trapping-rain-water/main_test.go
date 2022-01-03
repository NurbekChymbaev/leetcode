package _042_trapping_rain_water

import (
	"testing"
)

type Traps struct {
	height   []int
	expected int
}

var cases = []Traps{
	{
		[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		6,
	},
}

func TestTrap(t *testing.T) {

	for caseNum, caseData := range cases {
		var result = trap(caseData.height)

		if result != caseData.expected {
			t.Error("case:", caseNum,
				"\n\tstring: ", caseData.height,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
