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
}
