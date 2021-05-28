package simple

import (
	"strings"
)

// 方法1：使用strings标准库函数取出空格
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	length := len(s)
	if length == 0 {
		return 0
	}

	var lastLength int
	for i := length-1; i >= 0; i-- {
		if s[i] != ' ' {
			lastLength++
		} else {
			break
		}
	}

	return lastLength
}

func lengthOfLastWord2(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	var lastLength int
	for i := length-1; i >= 0; i-- {
		if s[i] != ' ' {
			lastLength++
		} else if lastLength > 0 {
			return lastLength
		}
	}

	return lastLength
}
