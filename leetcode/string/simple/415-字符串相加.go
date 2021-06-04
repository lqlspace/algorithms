package simple

func addString(num1 string, num2 string) string {
	lenNum1, lenNum2 := len(num1), len(num2)
	n := max(lenNum1, lenNum2)

	var carry int
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		if i < lenNum1 {
			carry += int(num1[lenNum1-1-i] - '0')
		}
		if i < lenNum2 {
			carry += int(num2[lenNum2-1-i] - '0')
		}

		res[n-1-i] = byte(carry % 10 + '0')
		carry /= 10
	}
	if carry > 0 {
		res = append([]byte("1"), res...)
	}

	return string(res)
}
