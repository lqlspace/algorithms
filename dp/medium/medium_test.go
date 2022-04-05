package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_longestCommonSubsequence(t *testing.T)  {
	t.Log(longestCommonSubsequence("abcde", "ace"))
	t.Log(longestCommonSubsequence("abcde", ""))
	t.Log(longestCommonSubsequence("abcde", "fgh"))
	t.Log(longestCommonSubsequence("abcde", "acef"))
	t.Log(longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"))
}

func Test_minPathSum(t *testing.T) {
	v := minPathSum([][]int{{1, 3, 1}, {1, 2, 1}, {4, 2, 1}})
	assert.Equal(t, 6, v)
}

func Test_coinChange(t *testing.T) {
	assert.Equal(t, 3, coinChange([]int{1, 2, 5}, 11))
	assert.Equal(t, -1, coinChange([]int{2}, 3))
	assert.Equal(t, 0, coinChange([]int{1}, 0))
}


func Test_findLength(t *testing.T) {
	res := findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7})
	t.Log(res)
	assert.Equal(t, 3, res)
}


func Test_maximalSquare(t *testing.T) {
	square := maximalSquare([][]byte{
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	})

	assert.Equal(t, 4, square)

	square = maximalSquare([][]byte{
		{'0', '1'},
	})

	assert.Equal(t, 1, square)
}