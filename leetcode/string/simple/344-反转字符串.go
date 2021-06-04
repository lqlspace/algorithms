package simple

// 时间复杂度O（n） ，空间复杂度O（1）
// 字符串转换成字符数组，即字节数组；字节数组转换成字符串用string()
func reverseString(s []byte)  {
	lenS := len(s)
	for i := 0; i < lenS / 2; i++ {
		s[i], s[lenS-1-i] = s[lenS-1-i], s[i]
	}
}
