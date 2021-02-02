package simple

import (
	"strings"
)

func getVal(str string) int {
	switch str {
	case "I":
		return 1
	case "a":
		return 4
	case "V":
		return 5
	case "b":
		return 9
	case "X":
		return 10
	case "c":
		return 40
	case "L":
		return 50
	case "d":
		return 90
	case "C":
		return 100
	case "e":
		return 400
	case "D":
		return 500
	case "f":
		return 900
	case "M":
		return 1000
	default:
		return 0
	}
}

// 先替换，再遍历
func romanToInt(s string) int {
	s = strings.ReplaceAll(s, "IV", "a")
	s = strings.ReplaceAll(s, "IX", "b")
	s = strings.ReplaceAll(s, "XL", "c")
	s = strings.ReplaceAll(s, "XC", "d")
	s = strings.ReplaceAll(s, "CD", "e")
	s = strings.ReplaceAll(s, "CM", "f")

	var sum int
	for _, item := range s {
		sum += getVal(string(item))
	}

	return sum
}


func romanToInt2(s string) int {
	if s == "" {
		return 0
	}

	pre, sum := getVal(string(s[0])), 0
	for i := 1; i < len(s); i++ {
		cur := getVal(string(s[i]))
		if pre < cur {
			sum -= pre
		} else {
			sum += pre
		}
		pre = cur
	}
	sum += pre

	return sum
}
