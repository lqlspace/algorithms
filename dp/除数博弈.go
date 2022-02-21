package dp

func divisorGames(n int) bool {
	f := make([]bool, n+3)
	f[1], f[2] = false, true
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i % j == 0 && !f[i-j] {
				f[i] = true
				break
			}
		}
	}

	return f[n]
}


//奇数必败，偶数必胜
func divisorGame2(n int) bool {
	return n % 2 == 0
}


