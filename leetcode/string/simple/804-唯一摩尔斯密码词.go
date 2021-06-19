package simple

import (
	"strings"
)

// 使用数组存放mores，时间复杂度O(N), 空间复杂度O(N)
func uniqueMorseRepresentations1(words []string) int {
	mores := []string{".-","-...","-.-.","-..",".","..-.","--.",
		"....","..",".---","-.-",".-..","--","-.",
		"---",".--.","--.-",".-.","...","-","..-",
		"...-",".--","-..-","-.--","--.."}

	m := make(map[string]struct{})
	for _, word := range words {
		buider := new(strings.Builder)
		n := len(word)
		for i := 0; i < n; i++ {
			buider.WriteString(mores[word[i]-'a'])
		}
		m[buider.String()] = struct{}{}
	}

	return len(m)
}

// 使用map存储会占用更多空间
func uniqueMorseRepresentations2(words []string) int {
	moreMap := map[string]string{"a": ".-", "b": "-...", "c": "-.-.", "d": "-..","e": ".","f": "..-.","g": "--.",
		"h": "....","i": "..","j": ".---","k": "-.-","l": ".-..","m": "--","n": "-.",
		"o": "---","p": ".--.","q": "--.-","r": ".-.","s": "...","t": "-","u": "..-",
		"v": "...-","w": ".--","x": "-..-","y": "-.--","z": "--.."}

	m := make(map[string]struct{})
	for _, word := range words {
		buider := new(strings.Builder)
		n := len(word)
		for i := 0; i < n; i++ {
			buider.WriteString(moreMap[string(word[i])])
		}
		m[buider.String()] = struct{}{}
	}

	return len(m)
}

func uniqueMorseRepresentations3(words []string) int {
	mores := []string{".-","-...","-.-.","-..",".","..-.","--.",
		"....","..",".---","-.-",".-..","--","-.",
		"---",".--.","--.-",".-.","...","-","..-",
		"...-",".--","-..-","-.--","--.."}

	m := make(map[string]struct{})
	for _, word := range words {
		b := make([]byte, 0)
		n := len(word)
		for i := 0; i < n; i++ {
			b = append(b, []byte(mores[word[i]-'a'])...)
		}
		m[string(b)] = struct{}{}
	}

	return len(m)
}
