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



//插入排序对于部分有效的数组十分高效，也很适合小规模数组
func InsertionSort(arr Interface, a, b int) {
	for i := a; i < b; i++ {
		for j := i; j > a && arr.Less(j, j-1); j-- {
			arr.Swap(j, j-1)
		}
	}
}

//希尔排序是对插入排序的改进，分组可缩小无序数组规模，而后合并成部分有序数组也可以发挥插入排序的长处
func ShellSort(arr Interface, a, b int) {
	n := b - a
	if n < 2 {
		return
	}

	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			for j := i; j >= key && arr.Less(j,j-key); j -= key {
				arr.Swap(j, j-key)
			}
		}
		key = key / 2
	}
}

//内层循环：遍历无序集合，找到min值对应索引下标;
//外层循环：指定内层循环求出的min所在的索引，并交换；
//时间复杂度：O(n^2)
func SelectionSort(arr Interface, a, b int) {
	for i := a; i < b; i++ {
		min := i
		for j := i + 1; j < b; j++ {
			if arr.Less(j, min) {
				min = j
			}
		}
		arr.Swap(min, i)
	}
}



