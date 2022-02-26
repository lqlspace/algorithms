package medium

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T)  {
	t.Log(longestPalindrome("ababcd"))
	t.Log(longestPalindrome2("ababcd"))
	t.Log(longestPalindrome("cbbd"))
	t.Log(longestPalindrome2("cbbd"))
	t.Log(longestPalindrome("a"))
	t.Log(longestPalindrome2("a"))
	t.Log(longestPalindrome(""))
	t.Log(longestPalindrome2(""))
	t.Log(longestPalindrome("abacdeeffe"))
	t.Log(longestPalindrome2("abacdeeffe"))
}
