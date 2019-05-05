package sort

type Interface interface{
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[j], is[i] = is[i], is[j]
}

func (is IntSlice) Sort() {
	insertionSort(is, 0, len(is))
}

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
