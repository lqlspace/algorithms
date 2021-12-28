package array

/*
1. 每一行都按照从左到右递增的顺序排序，每一列都按照从上到下的顺序排序
2. 比较的次数最大为O(rows+numbers)
 */

func Find(matrix [][]int, rows, columns, number int) bool {
	if matrix == nil || rows <= 0 || columns <= 0 {
		return false
	}

	row, column := 0, columns-1
	for row < rows && column >= 0 {
		if matrix[row][column] == number {
			return true
		} else if matrix[row][column] > number {
			column--
		} else {
			row++
		}
	}

	return false
}
