package array


func merge(A []int, m int, B []int, n int)  {
	i, j, k := m-1, n-1, 1
	for i >= 0 && j >= 0 {
		if A[i] > B[j] {
			A[m+n-k] = A[i]
			i--
		} else {
			A[m+n-k] = B[j]
			j--
		}
		k++
	}

	if j >= 0 {
		for i := 0; i <= j; i++ {
			A[i] = B[i]
		}
	}
}
