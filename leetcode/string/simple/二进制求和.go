package simple

import (
	"strconv"
)

/*
1. 在一个数组中，与下标i元素对称的是下标为len(a)-1-i的元素；
2. 数值型字符串转整型，直接减去'0'
 */
func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	var carry int
	var res string
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-1-i] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-1-i] - '0')
		}

		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}

	if carry > 0 {
		res = "1" + res
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
