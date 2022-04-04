package medium

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true 
	}

	var ans int 
	for num := range set {
		if !set[num-1] {
			start := num 
			l := 1 
			for set[start+1] {
				start++
				l++
			} 
			if ans < l {
				ans = l 
			}
		}
	}

	return ans 
}
