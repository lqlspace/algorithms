package tree

func CreateBTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] ==  null {
		return nil
	}

	m := map[int]*TreeNode{}

	root := &TreeNode{Val: arr[0]}
	m[0] = root
	cur := 0

	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] == null {
			if i % 2 == 0 {
				for cur <len(arr) {
					cur++
					if arr[cur] != null {
						break
					}
				}
			}
			continue
		}

		if i % 2 == 1 {
			node := &TreeNode{Val: arr[i]}
			m[cur].Left = node
			m[i] = node
		} else {
			node := &TreeNode{Val: arr[i]}
			m[cur].Right = node
			m[i] = node
			for cur <len(arr) {
				cur++
				if arr[cur] != null {
					break
				}
			}
		}
	}

	return root
}
