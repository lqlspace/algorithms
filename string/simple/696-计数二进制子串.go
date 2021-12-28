package simple

// 时间复杂度O(N)， 空间复杂度O(n)
func countBinarySubstrings1(s string) int {
	n := len(s)
	arr := make([]int, 0)
	var count int
	for i := 0; i < n; {
		val := s[i]
		for i < n && s[i] == val {
			count++
			i++
		}

		arr = append(arr, count)
		count = 0
	}

	var ans int
	for i := 1; i < len(arr); i++ {
		ans += min(arr[i], arr[i-1])
	}

	return ans
}

func countBinarySubstrings(s string) int {
	n := len(s)

	var last, count, ans int
	for i := 0; i < n; {
		val := s[i]
		for i < n && s[i] == val {
			count++
			i++
		}

		ans += min(last, count)
		last =  count
		count = 0
	}

	return ans
}
