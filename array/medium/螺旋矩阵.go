package medium

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, columns := len(matrix), len(matrix[0])
	orders, idx := make([]int, rows * columns), 0

	left, right, top, bottom := 0, columns-1, 0, rows-1
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			orders[idx] = matrix[top][col]
			idx++
		}
		for row := top+1; row <= bottom; row++ {
			orders[idx] = matrix[row][right]
			idx++
		}

		if top < bottom {
			for col := right-1; col >= left; col-- {
				orders[idx] = matrix[bottom][col]
				idx++
			}
		}

		if left < right {
			for row := bottom-1; row > top; row-- {
				orders[idx] = matrix[row][left]
				idx++
			}
		}

		left, right, top, bottom = left+1, right-1, top+1, bottom-1
	}

	return orders
}
