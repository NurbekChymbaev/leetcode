package _200_number_of_islands

func numIslands(grid [][]byte) int {

	islands := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				islands++
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, i int, j int) {

	if i >= len(grid) || i < 0 || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
