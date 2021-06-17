package simple

// 时间复杂度O(N), 空间复杂度O(N)
func reverseWords(s string) string {
	n := len(s)
	arr := []byte(s)

	for i := 0; i < n; {
		start := i
		for i < n && s[i] != ' ' {
			i++
		}

		for j, k := start, i-1; j < k; {
			arr[j], arr[k] = arr[k], arr[j]
			j++
			k--
		}

		for i < n && s[i] == ' ' {
			i++
		}
	}

	return string(arr)
}
