package dp

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	for i := range rows {
		rows[i] = make([]int, i+1)
		rows[i][0] = 1
		rows[i][i] = 1
		for j := 1; j < i; j++ {
			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
	}

	return rows
}
