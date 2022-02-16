package tree

import (
	"math"
	"reflect"
)


// 1. 广度优先搜索，注意空节点不能省略
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	left := bfsTraverse(p)
	right := bfsTraverse(q)

	return reflect.DeepEqual(left, right)
}

// 空间复杂度O(min(m,n))，时间复杂度O(min(m,n))
func bfsTraverse(root *TreeNode) []int {
	var queue []*TreeNode
	var vals []int
	queue = append(queue, root)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front == nil {
			vals = append(vals, math.MinInt64)
			continue
		}

		queue = append(queue, front.Left)
		queue = append(queue, front.Right)
		vals = append(vals, front.Val)
	}

	return  vals
}

// 2 深度优先搜索
// 时间复杂度O(min(m,n)), 空间复杂度O(min(m,n))
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return  false
	}

	return isSameTree2(p.Left, q.Left) && isSameTree2(p.Right, q.Right)
}
