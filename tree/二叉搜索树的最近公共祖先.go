package tree

// 两次遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)

	var ancestor *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i];i++ {
		ancestor = pathP[i]
	}

	return ancestor
}

// 二叉搜索树，可以比较
func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)

	return
}
