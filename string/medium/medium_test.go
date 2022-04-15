package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
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

func Test_reverseWords(t *testing.T) {
	t.Log(reverseWords("the   sky"))
	t.Log(reverseWords("the sky is blue"))
	t.Log(reverseWords("  hello world  "))
	t.Log(reverseWords("a good   example"))
	t.Log(reverseWords(""))
	t.Log(reverseWords(" "))
	t.Log(reverseWords(" a bb   cde  "))
}

func Test_multiply(t *testing.T) {
	t.Log(multiply("11", "2"))
	t.Log(multiply("25", "15"))
	t.Log(multiply("999", "999"))
}

func Test_compareVersion(t *testing.T) {
	res := compareVersion("1.0.1", "1.1")
	assert.Equal(t, -1, res)
}


func Test_validIPAddress(t *testing.T) {
	res := validIPAddress("172.16.254.1")
	assert.Equal(t, "IPv4", res)

	res = validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334")
	assert.Equal(t, "IPv6", res)

	res = validIPAddress("256.256.256.256")
	assert.Equal(t, "Neither", res)
}
