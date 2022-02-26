package simple

import (
	"strconv"
)

func addString(num1, num2 string) string {
	l1, l2 := len(num1), len(num2)
	n := max(l1, l2) + 1

	var carry byte
	strs :=  make([]byte, n)
	for i := 0; i < n; i++ {
		if i < l1 {
			carry += num1[l1-1-i] - '0'
		}
		if i < l2 {
			carry += num2[l2-1-i] - '0'
		}

		strs[n-1-i] = carry%10 + '0'
		carry /= 10
	}

	if strs[0] != '0' {
		return string(strs)
	}

	return string(strs[1:])
}



func addString2(num1, num2 string) string {
	var carry int
	var ans string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		sum := x + y + carry
		ans = strconv.Itoa(sum%10) + ans
		carry = sum / 10
	}

	return ans
}
