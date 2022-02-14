package tree

import (
	"math"
)

var null = math.MinInt64

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


