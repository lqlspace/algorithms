package simple

func majorityElement(nums []int) int {
	var ans, count int
	for _, v := range nums {
		if count == 0 {
			ans = v
		}
		if v == ans {
			count++
		} else {
			count--
		}
	}

	return ans
}
