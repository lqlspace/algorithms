package simple

import (
	"strconv"
)

func freqAlphabets(s string) string {
	arr := []byte(s)

	res := make([]byte, 0)

	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] != '#' {
			num, _ := strconv.Atoi(string(arr[i]))
			res = append(res, 'a' + uint8(num)-1)
		} else {
			i -= 2
			if i >= 0 {
				 num, _ := strconv.Atoi(s[i:i+2])
				 res = append(res, 'a' + uint8(num)-1)
			}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
