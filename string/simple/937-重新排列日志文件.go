package simple

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	letters := make([]string, 0)
	digits := make([]string, 0)

	for i := range logs {
		arr := strings.Split(logs[i], " ")
		if unicode.IsDigit(rune(arr[1][0])) {
			digits = append(digits, logs[i])
		} else {
			letters = append(letters, logs[i])
		}
	}

	sort.Slice(letters, func(i, j int) bool {
		idxA := strings.Index(letters[i], " ")
		idxB := strings.Index(letters[j], " ")
		if letters[i][idxA+1:] == letters[j][idxB+1:] {
			return letters[i][:idxA] < letters[j][:idxB]
		}

		return letters[i][idxA+1:] < letters[j][idxB+1:]
	})

	return append(letters, digits...)
}
