package medium


func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	queue, level := []*TreeNode{root}, -1
	for len(queue) > 0 {
		level++
		sz := len(queue)
		var arr []int
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		if level % 2 != 0 {
			reverseNums(arr)
		}

		ans = append(ans, arr)
	}

	return ans
}

func reverseNums(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
