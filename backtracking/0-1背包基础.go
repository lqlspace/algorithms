package backtracking


func Bag(i, cw, w int, items []int, maxWeight *int)  {
	if cw == w || i == len(items) {
		if cw > *maxWeight {
			*maxWeight = cw
		}
		return
	}

	Bag(i+1, cw, w, items, maxWeight)
	if cw + items[i] < w {
		Bag(i+1, cw+items[i], w, items, maxWeight)
	}
}
