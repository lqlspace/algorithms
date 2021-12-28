package simple

import (
	"strings"
)

func toGoatLatin1(sentence string) string {
	n := len(sentence)

	ans := new(strings.Builder)
	var idx int

	for i := 0; i < n; {
		for i < n && sentence[i] == ' ' {
			ans.WriteString(" ")
			i++
		}

		isVowel := isVowel(sentence[i])
		word := make([]byte, 0)
		for i < n && sentence[i] != ' ' {
			word = append(word, sentence[i])
			i++
		}
		idx++
		ans.WriteString(convert(word, isVowel, idx))
	}

	return ans.String()
}

func convert(word []byte, isVowel bool, idx int) string {
	if isVowel {
		word = append(word, 'm', 'a')
	} else {
		word = append(word[1:], word[0], 'm', 'a')

	}

	for i := 0; i < idx; i++ {
		word = append(word, 'a')
	}

	return string(word)
}

func isVowel(a uint8) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' ||
		a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
}


// 直接使用strings.Split函数的效率更高
func toGoatLatin(sentence string) string {
	ans := new(strings.Builder)
	for idx, elem := range strings.Split(sentence, " ") {
		if isVowel(elem[0]) {
			ans.WriteString(elem)
		} else {
			ans.WriteString(elem[1:])
			ans.WriteByte(elem[0])
		}
		ans.WriteString("ma")
		for i := 0; i < idx+1; i++ {
			ans.WriteByte('a')
		}
		ans.WriteString(" ")
	}

	return ans.String()[:len(ans.String())-1]
}
