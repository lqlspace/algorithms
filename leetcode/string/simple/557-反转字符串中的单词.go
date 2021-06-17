package simple

// 时间复杂度O(N), 空间复杂度O(1)
func reverseWords(s string) string {
	length := len(s)
	arr := []byte(s)
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for j, k := start, i-1; j < k; j++ {
			arr[j], arr[k]  = arr[k], arr[j]
			k--
		}

		for i < length && s[i] == ' ' {
			i++
		}
	}
	return string(arr)
}
