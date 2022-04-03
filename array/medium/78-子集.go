package medium


func subsets(nums []int) ([][]int) {
	n := len(nums)

	var ans [][]int 
	for i := 0; i < 1 << n; i++ {
		var sub []int 
		for j, v := range nums {
			if i >> j & 1 == 1 {
				sub = append(sub, v)
			} 
		}
		ans = append(ans, sub)
	}

	return ans 
}
