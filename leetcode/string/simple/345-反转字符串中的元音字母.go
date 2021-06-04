package simple

func reverseVowels(s string) string {
	arrS, lenS := []byte(s), len(s)
	i, j := 0, lenS-1

	for i < j {
		if !vowels(arrS[i]) {
			i++
			continue
		}
		if !vowels(arrS[j]) {
			j--
			continue
		}
		arrS[i], arrS[j] = arrS[j], arrS[i]
		i++
		j--
	}
	return string(arrS)
}

func vowels(e byte) bool {
	return e == 'a' || e == 'e' || e == 'o' || e == 'i' || e == 'u' ||
		e == 'A' || e == 'E' || e == 'O' || e == 'I' || e == 'U'
}
