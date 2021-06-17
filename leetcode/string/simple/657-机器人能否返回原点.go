package simple

// 时间复杂度O(N), 空间复杂度O(1)只需要常数的空间来存储变量
func judgeCircle(moves string) bool {
	var horizatal, vertical int

	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			vertical++
		}
		if moves[i] == 'D' {
			vertical--
		}
		if moves[i] == 'L' {
			horizatal--
		}
		if moves[i] == 'R' {
			horizatal++
		}
	}

	return horizatal == 0 && vertical == 0
}
