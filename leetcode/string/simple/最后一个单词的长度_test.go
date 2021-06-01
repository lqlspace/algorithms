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
