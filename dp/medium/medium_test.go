package medium

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T)  {
	t.Log(longestPalindrome("ababcd"))
	t.Log(longestPalindrome("cbbd"))
	t.Log(longestPalindrome("a"))
	t.Log(longestPalindrome(""))
	t.Log(longestPalindrome("abacdeeffe"))
}
