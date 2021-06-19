package simple

import (
	"unicode"
)

func reverseOnlyLetters(s string) string {
	n := len(s)
	arr := []byte(s)

	for low, high := 0, n-1; low < high; {
		for low < high && !unicode.IsLetter(rune(arr[low])) {
			low++
		}

		for low < high && !unicode.IsLetter(rune(arr[high])) {
			high--
		}

		if low < high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}

	return string(arr)
}
