func dfs(i int, j int, n int, m int, grid [][]byte) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	if i+1 < n {
		dfs(i+1, j, n, m, grid)
	}

	if i-1 >= 0 {
		dfs(i-1, j, n, m, grid)
	}

	if j+1 < m {
		dfs(i, j+1, n, m, grid)
	}

	if j-1 >= 0 {
		dfs(i, j-1, n, m, grid)
	}
}

func numIslands(grid [][]byte) int {
	count := 0
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j, n, m, grid)
			}
		}
	}
	return count
}
