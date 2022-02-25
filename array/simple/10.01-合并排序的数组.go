package simple

/*
	1. 将sub合入master;
	2. 从后往前比较，并从后往前排序（下标为m+n-k），数据移动次数最少，为O(m+n);
	3. 空间复杂度为O(1);
 */

func Merge(master []int, m int, sub []int, n int)  {
	i, j, k := m-1, n-1, 1
	for i >= 0 && j >= 0 {
		if master[i] > sub[j] {
			master[m+n-k] = master[i]
			i--
		} else {
			master[m+n-k] = sub[j]
			j--
		}
		k++
	}

	for j >= 0 {
		master[m+n-k] = sub[j]
		j--
	}
}
