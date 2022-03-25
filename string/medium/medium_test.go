package medium

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T)  {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("bbbbb"))
	t.Log(lengthOfLongestSubstring("pwwkew"))
}

func Test_myAtoi(t *testing.T) {
	t.Log(myAtoi("-42"))
	t.Log(myAtoi("  42"))
	t.Log(myAtoi("  -42abc"))
	t.Log(myAtoi("  -4+2-"))
	t.Log(myAtoi("-91283472332"))
	t.Log(myAtoi("9223372036854775808"))
	t.Log(myAtoi("21474836460"))
}

func Test_reverseWords(t *testing.T)  {
	t.Log(reverseWords("the   sky"))
	t.Log(reverseWords("the sky is blue"))
	t.Log(reverseWords("  hello world  "))
	t.Log(reverseWords("a good   example"))
	t.Log(reverseWords(""))
	t.Log(reverseWords(" "))
	t.Log(reverseWords(" a bb   cde  "))
}

func Test_multiply(t *testing.T)  {
	t.Log(multiply("11", "2"))
	t.Log(multiply("25", "15"))
	t.Log(multiply("999", "999"))
}
