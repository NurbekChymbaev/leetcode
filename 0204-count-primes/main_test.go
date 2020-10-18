package _204_count_primes

import "testing"

type CountPrimesCase struct {
	n        int
	expected int
}

var cases = []CountPrimesCase{
	{
		10, 4,
	},
	{
		0, 0,
	},
	{
		1, 0,
	},
	{
		2, 0,
	},
}

func TestCountPrimes(t *testing.T) {

	for caseNum, caseData := range cases {
		result := countPrimes(caseData.n)
		if caseData.expected != result {
			t.Error("case:", caseNum,
				"\n\tresult:  ", result,
				"\n\texpected:", caseData.expected)
		}
	}
}
