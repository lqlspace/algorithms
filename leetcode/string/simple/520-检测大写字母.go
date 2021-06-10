package simple

func detectCapitalUse(word string) bool {
	var cntlower, cntLarger int
	for i := range word {
		if 'A' <= word[i] && word[i] <= 'Z' {
			cntLarger++
		} else {
			cntlower++
		}
	}

	length := len(word)
	if cntlower == length || cntLarger == length {
		return true
	}

	if length > 0 && word[0] >= 'A' && word[0] <=  'Z' && cntlower == length-1 {
		return true
	}

	return false
}
