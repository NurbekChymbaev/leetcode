package _695_max_area_of_island

import "testing"

type Case struct {
	grid   [][]int
	answer int
}

var cases = []Case{
	{
		[][]int{
			{1, 1, 1, 1, 0},
			{1, 1, 0, 1, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		9,
	},
	{
		[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		},
		4,
	},

	{
		[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
		},
		4,
	},

	{
		[][]int{
			{1, 1},
		},
		2,
	},
}

func TestNumIslands(t *testing.T) {

	for num, data := range cases {
		result := maxAreaOfIsland(data.grid)
		if data.answer != result {
			t.Error("case:", num,
				"\n\tresult:  ", result,
				"\n\texpected:", data.answer)
		}
	}
}
