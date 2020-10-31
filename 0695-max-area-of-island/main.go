package _695_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {

	maxArea := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				area := 0
				dfs(grid, i, j, &area)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func dfs(grid [][]int, i int, j int, area *int) {

	if i >= len(grid) || i < 0 || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return
	}

	*area++
	grid[i][j] = 0

	dfs(grid, i+1, j, area)
	dfs(grid, i-1, j, area)
	dfs(grid, i, j+1, area)
	dfs(grid, i, j-1, area)
}
