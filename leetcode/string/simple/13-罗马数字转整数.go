package simple

import (
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"I": 1,
	"a": 4,
	"V": 5,
	"b": 9,
	"X": 10,
	"c": 40,
	"L": 50,
	"d": 90,
	"C": 100,
	"e": 400,
	"D": 500,
	"f": 900,
	"M": 1000,
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
		elem := strconv.Itoa(int(item))
		if val, ok := romanMap[elem]; ok {
			sum += val
		}
	}

	return sum
}
