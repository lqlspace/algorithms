package medium

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	rows, columns := len(matrix), len(matrix[0])
	orders := make([]int, rows * columns)

	left, right, top, bottom := 0, columns-1, 0, rows-1
	for i := 0; left <= right && top <= bottom; {
		for col := left; col <= right; col++ {
			orders[i] = matrix[top][col]
			i++
		}
		for row := top+1; row <= bottom; row++ {
			orders[i] = matrix[row][right]
			i++
		}

		if top < bottom {
			for col := right-1; col >= left; col-- {
				orders[i] = matrix[bottom][col]
				i++
			}
		}

		if left < right {
			for row := bottom-1; row > top; row-- {
				orders[i] = matrix[row][left]
				i++
			}
		}

		left, right, top, bottom = left+1, right-1, top+1, bottom-1
	}

	return orders
}
