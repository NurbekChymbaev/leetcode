package _011_container_with_most_water

func maxArea(height []int) int {

	p1 := 0
	p2 := len(height) - 1
	maxArea := 0

	for p1 < p2 {
		tempMax := min(height[p1], height[p2]) * (max(p1, p2) - min(p1, p2))
		maxArea = max(maxArea, tempMax)
		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
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
