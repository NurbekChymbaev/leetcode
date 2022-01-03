
package _042_trapping_rain_water

func trap(height []int) int {

	area := 0
	for i := 0; i < len(height); i++ {

		left := 0
		right := 0
		l := i
		r := i
		for l >= 0 {
			left = max(height[l], left)
			l--
		}

		for r < len(height) {
			right = max(height[r], right)
			r++
		}

		area += min(left, right) - height[i]
	}

	return area
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}