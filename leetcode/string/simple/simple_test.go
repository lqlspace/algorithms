package simple

import (
	"testing"
)

func Test_LengthOfLastWord(t *testing.T) {
	l := lengthOfLastWord("a")
	t.Log(l)
}


func Test_strStr2(t *testing.T)  {
	hayStack := "abcbd"
	needle := "cb"

	idx := strStr2(hayStack, needle)
	t.Log(idx)

	needle = "bc"
	idx = strStr2(hayStack, needle)
	t.Log(idx)
}

func Test_strStr3(t *testing.T)  {
	haystack := "sdfsdef"
	needle := "fs"

	pos := strStr3(haystack, needle)
	t.Log(pos)

	needle = "sde"
	pos = strStr3(haystack, needle)
	t.Log(pos)

	pos = strStr1(haystack, needle)
	t.Log(pos)

}

func Test_addBinary(t *testing.T)  {
	a := "1010"
	b := "01"

	t.Log(addBinary(a, b))

}

func Test_reverseString(t *testing.T)  {
	str := "abcde"
	sliceStr := []byte(str)
	reverseString(sliceStr)

	t.Log(string(sliceStr))
}

func Test_reverseVowels(t *testing.T)  {
	str := "abedEbA"
	str = reverseVowels(str)
	t.Log(str)

	str = "aabdefgOh"
	str = reverseVowels(str)
	t.Log(str)

	str = ""
	str = reverseVowels(str)
	t.Log(str)

	str = "I"
	str = reverseVowels(str)
	t.Log(str)
}

func Test_addString(t *testing.T)  {
	a := "5024"
	b := "100"

	c := addString(a, b)
	t.Log(c)

	a = "8096"
	b = "2107"
	c = addString(a, b)
	t.Log(c)
}

func Test_countSegments(t *testing.T)  {
	words := "hello, world, b cd"
	count := countSegments(words)
	t.Log(count)

	words = ""
	count = countSegments(words)
	t.Log(count)

	words = "  a  b  "
	count = countSegments(words)
	t.Log(count)
}

func Test_repeatedSubstring(t *testing.T)  {
	a := "abcabcabcabc"
	t.Log(repeatedSubstringPattern(a))
	t.Log(repeatedSubstringPattern2(a))

	a = "abc"
	t.Log(repeatedSubstringPattern(a))
	t.Log(repeatedSubstringPattern2(a))

	a = ""
	t.Log(repeatedSubstringPattern(a))
	t.Log(repeatedSubstringPattern2(a))

}

func Test_detectCapitalWord(t *testing.T)  {
	a := "abcd"
	t.Log(detectCapitalUse(a))

	a = "Hello"
	t.Log(detectCapitalUse(a))

	a = "USA"
	t.Log(detectCapitalUse(a))
}

func Test_findUSlength(t *testing.T)  {
	a := "abc"
	b := "abc"
	t.Log(findUSlength(a, b))

	a = "abc"
	b = "abd"
	t.Log(findUSlength(a, b))

	a = "abc"
	b = ""
	t.Log(findUSlength(a, b))

	a = "abc"
	b = "abdee"
	t.Log(findUSlength(a, b))
}

func Test_reverseStr(t *testing.T)  {
	s := "abc"
	s = reverseStr(s, 2)
	t.Log(s)

	s = "abcdef"
	s = reverseStr(s, 3)
	t.Log(s)
}

func Test_checkRecord(t *testing.T)  {
	s := "PPL"
	award := checkRecord(s)
	t.Log(award)

	s = "LALL"
	award = checkRecord(s)
	t.Log(award)
}

func Test_reverseWords(t *testing.T)  {
	s := "hello words"
	s = reverseWords(s)
	t.Log(s)

	s = "ab cd e"
	s = reverseWords(s)
	t.Log(s)
}
