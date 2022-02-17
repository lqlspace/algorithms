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

// 方法2：时间复杂度O(N)，空间复杂度O(1)
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	ancestor := root

	for ancestor != nil {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			break
		}
	}

	return ancestor
}
