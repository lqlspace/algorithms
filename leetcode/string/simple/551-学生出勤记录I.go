package simple

func checkRecord(s string) bool {
	n := len(s)
	var numL, numA int
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			numL = 0
			numA++
			if numA > 1 {
				return false
			}
		} else if s[i] == 'L' {
			numL++
			if numL > 2 {
				return false
			}
		} else {
			numL = 0
		}
	}

	return true
}
