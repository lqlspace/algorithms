package tree

import (
	"testing"
)

func TestCreateBTree(t *testing.T) {
	bt := CreateBTree()
	bt.Insert(nil, "a", 0)
	bt.Insert("a", "b", 0)
	bt.Insert("a", "d", 1)
	bt.Insert("b", "c", 1)
	bt.Insert("d", "e", 0)
	bt.Insert("d", "f", 1)

	bt.InOrderTraverse(bt.root)
	println()
	bt.PreOrderTraverse(bt.root)
	println()
	bt.PostOrderTraverse(bt.root)
	println()
	bt.LevelTraverse()
	println()
}

func TestBTree_LevelSearch(t *testing.T) {
	bt := CreateBTree()
	bt.InsertByLevelSearch(nil, "a", 0)
	bt.InsertByLevelSearch("a", "b", 0)
	bt.InsertByLevelSearch("a", "d", 1)
	bt.InsertByLevelSearch("b", "c", 1)
	bt.InsertByLevelSearch("d", "e", 0)
	bt.InsertByLevelSearch("d", "f", 1)

	bt.InOrderTraverse(bt.root)
	println()
	bt.PreOrderTraverse(bt.root)
	println()
	bt.PostOrderTraverse(bt.root)
	println()
	bt.LevelTraverse()
	println()
}
