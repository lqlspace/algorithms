package simple

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	n := len(paragraph)
	mp := make(map[string]int)
	var maxCnt int
	var  res string

	bannedMap := make(map[string]struct{})
	for _, val := range banned {
		bannedMap[val] = struct{}{}
	}

	for i := 0; i < n; {
		for i < n && !unicode.IsLetter(rune(paragraph[i])) {
			i++
		}
		wordBytes := make([]byte, 0)
		for i < n && unicode.IsLetter(rune(paragraph[i])) {
			wordBytes = append(wordBytes, paragraph[i])
			i++
		}
		word := strings.ToLower(string(wordBytes))
		mp[word]++

		if _, ok := bannedMap[word]; !ok && mp[word] > maxCnt {
			res = word
			maxCnt = mp[word]
		}
	}

	return res
}
