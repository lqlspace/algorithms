package medium

func CreateBTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	tree := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{tree}
	curIdx := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if curIdx < len(arr) && arr[curIdx] != null {
			node.Left = &TreeNode{Val: arr[curIdx]}
			queue = append(queue, node.Left)
		}
		if curIdx + 1 < len(arr) && arr[curIdx+1] != null {
			node.Right = &TreeNode{Val: arr[curIdx+1]}
			queue = append(queue, node.Right)
		}
		curIdx += 2
	}

	return tree
}


func LevelTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var arr []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			arr = append(arr, null)
			continue
		}

		arr = append(arr, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	num := 0
	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] != null {
			break
		}
		num++
	}

	return  arr[:len(arr)-num]
}

func dfs(head *TreeNode, val int) (*TreeNode, bool) {
	if head == nil {
		return nil, false
	}

	if head.Val == val {
		return head, true
	}

	left, exist := dfs(head.Left, val)
	if !exist {
		return dfs(head.Right, val)
	}

	return left, true
}
