package simple

// 1、滑动窗口暴力法，时间复杂度O(（N-L）L)
func strStr1(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l1 < l2 {
		return -1
	}

	var pos = -1
	for i := 0; i < l1-l2+1; i++ {
		str := haystack[i:i+l2]
		if str == needle {
			pos = i
			break
		}
	}

	return pos
}
