package tree

type BTree struct {
	root *TreeNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


