package simple


func sortedArrayToBST(nums []int) *TreeNode {
	return  buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := (low+high) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBST(nums, low, mid-1)
	root.Right = buildBST(nums, mid+1, high)

	return root
}
