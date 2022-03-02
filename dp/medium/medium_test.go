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

func Test_lengthOfLIS(t *testing.T)  {
	t.Log(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	t.Log(lengthOfLIS([]int{0,1,0,3,2,3}))
	t.Log(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}
