package sort

// 时间复杂度O(n),目前实现要求排序的元素必须不重复；
func BitmapSort(items []int, min, max int) {
	lens := (max-min) / 8 + 1
	s := make([]int, lens)

	for i := range items {
		s[items[i]/8] |= (1 << uint(items[i]%8))
	}

	var k int
	for i := range s {
		for j := 0; j < 8; j++ {
			if s[i] & (1 << uint(j)) > 0 {
				items[k] = i * 8 + j
				k++
			}
		}
	}
}
