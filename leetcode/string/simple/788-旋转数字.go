package simple


func rotatedDigits(n int) int {
	var ans int
	var flag bool
	for i := 1; i <= n; i++ {
		if good(i, flag) {
			ans++
		}
	}

	return ans
}

func good(n int, flag bool) bool {
	if n == 0 {
		return flag
	}

	d := n % 10
	if d == 3 || d == 4 || d == 7 {
		return false
	}

	if d == 0 || d == 1 || d == 8 {
		return good(n/10, flag)
	}

	return good(n/10, true)
}
