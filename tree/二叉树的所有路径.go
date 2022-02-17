package tree

import (
	"strconv"
)

// 方法1：深度优先搜索
func binaryTreePaths(root *TreeNode) (paths []string) {
	paths = []string{}

	var nodePath func(node *TreeNode, path string)
	nodePath = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			return
		}

		path += "->"
		nodePath(node.Left, path)
		nodePath(node.Right, path)
	}

	nodePath(root, "")

	return
}

func binaryTreePaths2(root *TreeNode) (paths []string) {
	paths = []string{}
	if root == nil {
		return paths
	}

	nodeQueue := []*TreeNode{root}
	pathQueue := []string{strconv.Itoa(root.Val)}

	for i := 0; i < len(nodeQueue); i++ {
		if nodeQueue[i].Left == nil && nodeQueue[i].Right == nil {
			paths = append(paths, pathQueue[i])
			continue
		}

		if nodeQueue[i].Left != nil {
			nodeQueue = append(nodeQueue, nodeQueue[i].Left)
			pathQueue = append(pathQueue, pathQueue[i] + "->" + strconv.Itoa(nodeQueue[i].Left.Val))
		}

		if nodeQueue[i].Right != nil {
			nodeQueue = append(nodeQueue, nodeQueue[i].Right)
			pathQueue = append(pathQueue, pathQueue[i] + "->" + strconv.Itoa(nodeQueue[i].Right.Val))
		}
	}

	return
}

