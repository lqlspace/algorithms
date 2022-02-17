package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBTree(t *testing.T) {
	arr := []int{1, 2, 3}
	bt := CreateBTree(arr)
	assert.Equal(t, arr, LevelTraversal(bt))

	arr = []int{1, null, 2}
	bt = CreateBTree(arr)
	assert.Equal(t, arr, LevelTraversal(bt))

	arr = []int{1, null, 2, null, 3}
	bt = CreateBTree(arr)
	assert.Equal(t, arr, LevelTraversal(bt))

	arr = []int{1, 2, 3, null, null, 4, 5, null, 6}
	bt = CreateBTree(arr)
	assert.Equal(t, arr, LevelTraversal(bt))
}

func Test_inorderTraversal(t *testing.T) {
	node := CreateBTree([]int{1, 2, 3, null, null, 4})

	arr := inorderTraversal(node)

	t.Log(arr)
}

func Test_isSameTree(t *testing.T) {
	node1 := CreateBTree([]int{1, 2})
	node2 := CreateBTree([]int{1, null, 2})

	t.Log(isSameTree1(node1, node2))
	t.Log(isSameTree2(node1, node2))
}

func Test_isSymmetric(t *testing.T) {
	bt := CreateBTree([]int{1, 2, 2, 2, null, 2})
	t.Log(isSymmetric(bt))

	bt = CreateBTree([]int{1, 2, 2, 3, 4, 4, 3})
	t.Log(isSymmetric(bt))

	bt = CreateBTree([]int{5, 4, 1, null, 1, null, 4, 2, null, 2, null})
	t.Log(isSymmetric(bt))
}

func Test_maxDepth(t *testing.T) {
	bt := CreateBTree([]int{3, 9, 20, null, null, 15, 7})
	t.Log(maxDepth(bt))
	t.Log(maxDepth2(bt))
}

func Test_sortedArrayToBST(t *testing.T)  {
	arr := []int{-10, -3, 0, 5, 9}
	t.Log(sortedArrayToBST(arr))
}

func Test_isBalanced(t *testing.T) {
	arr := []int{1, 2,  3, 4, 5}
	bt := CreateBTree(arr)
	t.Log(isBalanced(bt))
	t.Log(isBalanced2(bt))

	arr = []int{1, 2,  3, 4, 5, null, null, 6}
	bt = CreateBTree(arr)
	t.Log(isBalanced(bt))
	t.Log(isBalanced2(bt))
}

func Test_minDepth(t *testing.T)  {
	arr := []int{1, 2, 3, null, null, 4, 5, 6}
	bt := CreateBTree(arr)

	t.Log(minDepth(bt))
	t.Log(minDepth2(bt))

	arr = []int{1, null, 3, 4, 5, 6}
	bt = CreateBTree(arr)

	t.Log(minDepth(bt))
	t.Log(minDepth2(bt))
}

func Test_hasPathSum(t *testing.T) {
	arr := []int{5,4,8,11,null,13,4,7,2,null,null,null,1}
	bt := CreateBTree(arr)
	t.Log(hasPathSum(bt, 22))
	t.Log(hasPathSum2(bt, 22))

	arr = []int{1}
	bt = CreateBTree(arr)
	t.Log(hasPathSum(bt, 1))
	t.Log(hasPathSum2(bt, 1))

	arr = []int{1, 2, 3, 4, null, 1}
	bt = CreateBTree(arr)
	t.Log(hasPathSum(bt, 5))
	t.Log(hasPathSum2(bt, 5))

	arr = []int{}
	bt = CreateBTree(arr)
	t.Log(hasPathSum(bt, 0))
	t.Log(hasPathSum2(bt, 0))

	arr = []int{1, 2}
	bt = CreateBTree(arr)
	t.Log(hasPathSum(bt, 2))
	t.Log(hasPathSum2(bt, 2))
}

func Test_preorderTraversal(t *testing.T)  {
	arr := []int{1, null, 2, 3}
	bt := CreateBTree(arr)
	t.Log(preorderTraversal(bt))

	arr = []int{}
	bt = CreateBTree(arr)
	t.Log(preorderTraversal(bt))

	arr = []int{1}
	bt = CreateBTree(arr)
	t.Log(preorderTraversal(bt))
}

func Test_postorderTraversal(t *testing.T)  {
	arr := []int{1, 2, 3}
	bt := CreateBTree(arr)
	t.Log(postorderTraversal(bt))
}

func Test_invertTree(t *testing.T) {
	arr := []int{4, 2, 7, 1, 3, 6, 9}
	bt := CreateBTree(arr)
	t.Log(LevelTraversal(invertTree(bt)))

	arr = []int{1, 2, 3}
	bt = CreateBTree(arr)
	t.Log(LevelTraversal(invertTree(bt)))
}

func TestLeverSearch(t *testing.T) {
	arr := []int{1, null, 2}
	bt := CreateBTree(arr)
	t.Log(LevelTraversal(bt))

	arr = []int{1, 2, 3, null,  5}
	bt = CreateBTree(arr)
	t.Log(LevelTraversal(bt))

	arr = []int{1, 2, 3, null, null, 4, 5, null, null, null, 6}
	bt = CreateBTree(arr)
	t.Log(LevelTraversal(bt))

}

func Test_binaryTreePaths(t *testing.T) {
	arr := []int{1, 2, 3, null, null, 4}
	bt :=  CreateBTree(arr)
	t.Log(binaryTreePaths(bt))
	t.Log(binaryTreePaths2(bt))

	arr = []int{1, 2, 3, 7, null, 4, 5, null, null, null, null, 6}
	bt =  CreateBTree(arr)
	t.Log(binaryTreePaths(bt))
	t.Log(binaryTreePaths2(bt))
}

func Test_sumOfLeftLeaves(t *testing.T) {
	arr := []int{3,9,20,null,null,15,7}
	bt := CreateBTree(arr)
	t.Log(sumOfLeftLeaves(bt))
	t.Log(sumOfLeftLeaves2(bt))

	arr = []int{1, 2, 3, null, null, 4}
	bt = CreateBTree(arr)
	t.Log(sumOfLeftLeaves(bt))
	t.Log(sumOfLeftLeaves2(bt))
}

func Test_findMode(t *testing.T)  {
	arr := []int{1,null,2,2}
	bt := CreateBTree(arr)
	t.Log(findMode(bt))

	arr = []int{0}
	bt = CreateBTree(arr)
	t.Log(findMode(bt))
}

func Test_getMinimumDifference(t *testing.T) {
	arr := []int{4,2,6,1,3}
	bt := CreateBTree(arr)
	t.Log(getMinimumDifference(bt))

	arr = []int{4,3,6,1,3}
	bt = CreateBTree(arr)
	t.Log(getMinimumDifference(bt))
}
