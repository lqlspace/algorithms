package simple

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	com := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		com = longestCommon(com, str)
	}

	return com
}


func longestCommon(str1, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	var comLen int
	if l1 < l2 {
		comLen = l1
	} else {
		comLen = l2
	}

	var res []byte
	for i := 0; i < comLen; i++ {
		if str1[i] != str2[i] {
			break
		}
		res = append(res, str1[i])
	}

	return string(res)
}
