package simple

func validPalindrome(s string) bool {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		if s[low] != s[high] {
			left, right := true, true
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					left = false
					break
				}
			}
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					right = false
					break
				}
			}
			return left || right
		}
	}

	return true
}
