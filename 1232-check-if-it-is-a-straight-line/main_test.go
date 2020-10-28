package _232_check_if_it_is_a_straight_line

import (
	"testing"
)

func TestCheckStraightLine(t *testing.T) {

	checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}})

	checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}})

	checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 5}})
}
