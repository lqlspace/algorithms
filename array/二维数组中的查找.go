package array

/*
二维数组中的查找：每一行都按照从左到右递增的顺序排序，每一列都按照从上到下的顺序排序
 */

func Find(matrix [][]int, rows, columns, number int) bool {
	var found bool

	if matrix != nil && rows > 0 && columns > 0 {
		row := 0
		column := columns -1
		for row < rows && column >= 0 {
			if matrix[row][column]== number {
				found = true
				break
			} else if matrix[row][column] > number {
				column--
			} else {
				row++
			}
		}
	}

	return found
}
