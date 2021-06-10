package simple

func findUSlength(a, b string) int {
	if a == b {
		return -1
	}

	if len(a) == len(b) {
		return len(a)
	}

	return max(len(a), len(b))
}
