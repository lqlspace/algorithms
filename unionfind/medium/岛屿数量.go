package medium


// 深度优先搜索
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
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



// 广度优先搜索
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var num int
	nr, nc := len(grid), len(grid[0])
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				num++
				grid[i][j] = '0'
				var queue []int
				queue = append(queue, i*nc + j)
				for len(queue) > 0 {
					elem := queue[0]
					queue = queue[1:]
					row := elem / nc
					col := elem % nc
					if row - 1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, (row-1)*nc+col)
						grid[row-1][col] = '0'
					}
					if row + 1 < nr && grid[row+1][col] == '1' {
						queue = append(queue, (row+1)*nc+col)
						grid[row+1][col] = '0'
					}
					if col - 1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, row*nc+col-1)
						grid[row][col-1] = '0'
					}
					if col + 1 < nc && grid[row][col+1] == '1' {
						queue = append(queue, row*nc+col+1)
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}

	return num
}
