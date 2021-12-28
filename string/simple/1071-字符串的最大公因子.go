package simple

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	for i := min(n, m); i >= 1; i-- {
		if n %  i ==  0 && m % i == 0 {
			sub := str1[:i]
			if check(sub, str1) && check(sub, str2) {
				return sub
			}
		}
	}

	return ""
}

func check(sub, s string) bool {
	n := len(s) / len(sub)
	res := new(strings.Builder)
	for i := 0; i < n; i++ {
		res.WriteString(sub)
	}

	return res.String() == s
}
