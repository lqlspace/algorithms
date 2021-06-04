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
