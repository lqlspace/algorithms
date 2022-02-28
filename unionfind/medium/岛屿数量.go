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


type UnionFindSet struct {
	Parents []int // 每个结点的顶级节点
	SetCount int // 连通分量的个数
}


func (u *UnionFindSet) Init(grid [][]byte) {
	row := len(grid)
	col := len(grid[0])
	count := row * col
	u.Parents = make([]int, count)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			u.Parents[i*col+j] = i * col + j
			if grid[i][j] == '1' {
				u.SetCount++
			}
		}
	}
}

func (u *UnionFindSet) Find(node int) int {
	if u.Parents[node] == node {
		return node
	}

	root := u.Find(u.Parents[node])
	u.Parents[node] = root
	return root
}

func (u *UnionFindSet) Union(node1, node2 int) {
	root1 := u.Find(node1)
	root2 := u.Find(node2)

	if root1 == root2 {
		return
	} else if root1 < root2 {
		u.Parents[root1] = root2
	} else {
		u.Parents[root2] = root1
	}
	u.SetCount--
}

func numIslands3(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	u := &UnionFindSet{}
	u.Init(grid)

	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				if i - 1 >= 0 && grid[i-1][j] == '1' {
					u.Union(i*col+j, (i-1)*col+j)
				}
				if i + 1 < row && grid[i+1][j] == '1' {
					u.Union(i*col+j, (i+1)*col+j)
				}
				if j - 1 >= 0 && grid[i][j-1] == '1' {
					u.Union(i*col+j, i*col+j-1)
				}
				if j + 1 < col && grid[i][j+1] == '1' {
					u.Union(i*col+j, i*col+j+1)
				}
				grid[i][j] = '0'
			}
		}
	}

	return u.SetCount
}
