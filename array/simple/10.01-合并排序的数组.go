package simple

/*
	1. 将sub合入master;
	2. 从后往前比较，并从后往前排序（下标为m+n-k），数据移动次数最少，为O(m+n);
	3. 空间复杂度为O(1);
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	if j >= 0 {

		copy(nums1[:k+1], nums2[:j])
	}
}
