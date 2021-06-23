package simple

import (
	"math"
)

func maxNumberOfBalloons(text string) int {
	arr := make([]int, 26)

	for _, v := range text {
		arr[v-'a']++
	}

	var min = math.MaxInt64
	for i := range arr {
		if i  == 'b' - 'a' || i == 'a' - 'a' || i == 'n' - 'a' {
			if arr[i] < min {
				min = arr[i]
			}
		} else if i == 'l' - 'a' || i == 'o' - 'a' {
			if arr[i] / 2  < min {
				min = arr[i] / 2
			}
		}
	}

	return min
}
