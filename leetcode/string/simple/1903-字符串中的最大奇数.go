package simple

//时间复杂度O(N), 空间复杂度O(1)输出字符串不计入空间复杂度
func largestOddNumber(num string) string {
	for i := len(num)-1; i >= 0; i-- {
		if (num[i]-'0') % 2 == 1 {
			return num[:i+1]
		}
	}

	return ""
}
