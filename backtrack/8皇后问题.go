package backtrack

var res = make([]int, 8)
var cnt = 0

func Cal8Queue(row int)  {
	if row == 8 {
		printQueues(res)
		return
	}

	for col := 0; col < 8; col++ {
		if isOk(row, col) {
			res[row] =  col
			Cal8Queue(row+1)
		}
	}
}

func isOk(row,  col int) bool {
	leftup := col - 1
	rightup := col + 1

	for i := row-1; i >= 0; i-- {
		if res[i] == col {
			return  false
		}
		if leftup >= 0 {
			if res[i]  == leftup {
				return false
			}
		}

		if rightup < 8 {
			if res[i] == rightup {
				return false
			}
		}

		leftup--
		rightup++
	}
	return  true
}


func printQueues(result []int)  {
	for row :=  0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if result[row] == col {
				print("Q  ")
			} else {
				print("* ")
			}
		}
		println()
	}
	cnt++
	println(cnt)
}

