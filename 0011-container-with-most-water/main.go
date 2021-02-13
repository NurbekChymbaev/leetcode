package _011_container_with_most_water

func maxArea(height []int) int {
	res := 0
	for i := range height {
		tempMax := 0
		for j := range height {
			if i == j {
				continue
			}
			tempMax = min(height[i], height[j]) * (max(i, j) - min(i, j))
			if tempMax > res {
				res = tempMax
			}
		}
	}

	return res
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
