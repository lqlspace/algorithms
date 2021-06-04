package simple

func countSegments(s string) int {
	lenS := len(s)

	var count int
	var start bool
	for i := 0; i < lenS; i++ {
		if s[i] !=  ' ' {
			if !start {
				start = true
				count++
			}
		} else {
			start = false
		}
	}

	return count
}
