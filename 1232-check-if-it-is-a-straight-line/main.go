package _232_check_if_it_is_a_straight_line

func checkStraightLine(coordinates [][]int) bool {

	if len(coordinates) == 2 {
		return true
	}

	for i := 0; i < len(coordinates)-2; i++ {

		ax := float64(coordinates[i][0])
		ay := float64(coordinates[i][1])
		bx := float64(coordinates[i+1][0])
		by := float64(coordinates[i+1][1])
		cx := float64(coordinates[i+2][0])
		cy := float64(coordinates[i+2][1])

		area := (ax*(by-cy) + bx*(cy-ay) + cx*(ay-by)) / 2.0

		if area != 0.0 {
			return false
		}
	}

	return true
}
