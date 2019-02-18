package bubble

func BubbleSort(items []int) {
	n := len(items)

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}
