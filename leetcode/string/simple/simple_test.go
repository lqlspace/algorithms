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

func Test_tree2str(t *testing.T)  {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	s := tree2str(tree)
	t.Log(s)

	tree = &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	s = tree2str(tree)
	t.Log(s)
}

func Test_judgeCircle(t *testing.T)  {
	s := "UUDDLURD"
	res := judgeCircle(s)
	t.Log(res)

	s = "UU"
	res = judgeCircle(s)
	t.Log(res)
}

func Test_validPalindrome(t *testing.T)  {
	s := "aba"
	res := validPalindrome(s)
	t.Log(res)

	s = "abca"
	res = validPalindrome(s)
	t.Log(res)

	s = "abcca"
	res = validPalindrome(s)
	t.Log(res)
}

func Test_countBinarySubstrings(t *testing.T)  {
	s := "00110011"
	count := countBinarySubstrings(s)
	t.Log(count)

	s = "01"
	count = countBinarySubstrings(s)
	t.Log(count)
}

func Test_toLowerCase(t *testing.T)  {
	s := "Hello World"
	s = toLowerCase(s)
	t.Log(s)
}

func Test_rotatedDigits(t *testing.T)  {
	a := 10
	ans := rotatedDigits(a)
	t.Log(ans)
}

func Test_uniqueMorseRepresentations(t *testing.T)  {
	words := []string{"gin", "zen", "gig", "msg"}
	n := uniqueMorseRepresentations3(words)
	t.Log(n)
}

func Test_mostCommonWord(t *testing.T)  {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	word := mostCommonWord(paragraph, banned)
	t.Log(word)
}

func Test_toGoatLatin(t *testing.T)  {
	para := "I speak Goat Latin"
	ans := toGoatLatin(para)
	res := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
	if ans != res {
		t.Log("not equal")
	} else {
		t.Log("equal")
	}
	t.Log(ans)
}

func Test_buddyStrings(t *testing.T)  {
	s := "ab"
	g := "ba"
	ans := buddyStrings(s, g)
	t.Log(ans)

	s = "ab"
	g = "ab"
	ans = buddyStrings(s, g)
	t.Log(ans)

	s = "aa"
	g = "aa"
	ans = buddyStrings(s, g)
	t.Log(ans)

	s = "aaaaaaabc"
	g = "aaaaaaacb"
	ans = buddyStrings(s, g)
	t.Log(ans)

	s = ""
	g = "aa"
	ans = buddyStrings(s, g)
	t.Log(ans)
}

func Test_reverseOnlyLetters(t *testing.T)  {
	s := "ab-cd"
	res := reverseOnlyLetters(s)
	t.Log(res)

	s = "a-bC-dEf-ghIj"
	res = reverseOnlyLetters(s)
	t.Log(res)
}

func Test_isLongPressedName(t *testing.T)  {
	//name := "alex"
	//typed := "aaleex"
	//res := isLongPressedName(name, typed)
	//t.Log(res)
	//
	//name = "saeed"
	//typed = "ssaaedd"
	//res = isLongPressedName(name, typed)
	//t.Log(res)
	//
	//name = "leelee"
	//typed = "lleeelee"
	//res = isLongPressedName(name, typed)
	//t.Log(res)
	//
	//name = "vtkgn"
	//typed = "vttkgnn"
	//res = isLongPressedName(name, typed)
	//t.Log(res)

	name := "abcd"
	typed := "aaabbbcccddd"
	res := isLongPressedName(name, typed)
	t.Log(res)
}

func Test_numUniqueEmails(t *testing.T)  {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	num := numUniqueEmails(emails)
	t.Log(num)

	emails = []string{"test.email+alex@leetcode.com","test.email.leet+alex@code.com"}
	num = numUniqueEmails(emails)
	t.Log(num)
}

func Test_reorderLogFiles(t *testing.T)  {
	logs := []string{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"}
	logs = reorderLogFiles(logs)
	t.Log(logs)

	logs = []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	logs = reorderLogFiles(logs)
	t.Log(logs)

	logs = []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo","a2 act car"}
	logs = reorderLogFiles(logs)
	t.Log(logs)

	logs = []string{"dig1 8 1 5 1","let1 art zero can","dig2 3 6","let2 own kit dig","let3 art zero"}
	logs = reorderLogFiles(logs)
	t.Log(logs)
}

func Test_gcdOfStrings(t *testing.T)  {
	a := "AB"
	b := "ABAB"
	res := gcdOfStrings(a, b)
	t.Log(res)

	a = "ABABAB"
	res = gcdOfStrings(a, b)
	t.Log(res)

	a = "A"
	res = gcdOfStrings(a, b)
	t.Log(res)
}

func Test_defangIPaddr(t *testing.T)  {
	ip := "1.1.1.127"
	ip = defangIPaddr(ip)
	t.Log(ip)
}
