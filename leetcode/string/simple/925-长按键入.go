package simple

func isLongPressedName(name string, typed string) bool {
	n, m := len(name), len(typed)

	i, j := 0, 0
	for j < m {
		if i < n && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}

	return i == n
}
