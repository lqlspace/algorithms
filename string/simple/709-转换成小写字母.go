package simple

func toLowerCase(s string) string {
	arr := []byte(s)
	n := len(s)

	for i := 0; i < n; i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] + 32
		}
	}

	return string(arr)
}

