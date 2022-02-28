package medium


// 深度优先搜索
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	var num int
	nr, nc := len(grid), len(grid[0])
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, nr, nc, i, j)
			}
		}
	}

	return num
}

func dfs(grid [][]byte, nr, nc, row, col int) {
	if row < 0 || col < 0 || row >= nr || col >= nc || grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'
	dfs(grid, nr, nc, row-1, col)
	dfs(grid, nr, nc, row+1, col)
	dfs(grid, nr, nc, row, col-1)
	dfs(grid, nr, nc, row, col+1)
}
