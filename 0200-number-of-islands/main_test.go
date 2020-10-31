package _200_number_of_islands

import "testing"

type Case struct {
	grid   [][]byte
	answer int
}

var cases = []Case{
	{
		[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		1,
	},
	{
		[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		3,
	},

	{
		[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
		},
		2,
	},

	{
		[][]byte{
			{'1', '1'},
		},
		1,
	},
}

func TestNumIslands(t *testing.T) {

	for num, data := range cases {
		result := numIslands(data.grid)
		if data.answer != result {
			t.Error("case:", num,
				"\n\tresult:  ", result,
				"\n\texpected:", data.answer)
		}
	}
}
