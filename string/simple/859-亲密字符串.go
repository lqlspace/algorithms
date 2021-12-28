package simple

// 此处使用arr存放26个小写字母的个数
// 时间复杂度O(N), 空间复杂度O(1)
func buddyStrings(s string, goal string) bool {
	n, m := len(s), len(goal)
	if n != m {
		return false
	}

	if s == goal {
		arr := make([]int, 26)
		for i := range s {
			arr[s[i]-'a']++
			if arr[s[i]-'a'] > 1 {
				return true
			}
		}

		return false
	}

	first, second := -1, -1
	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	return second > -1 && s[first] == goal[second] && s[second] == goal[first]

}
