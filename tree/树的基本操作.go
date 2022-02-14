package tree


func CreateBTree() *BTree {
	return &BTree{}
}

func (bt *BTree) Insert(parent interface{},  data , val int) {
	if parent == nil && bt.root == nil {
		bt.root = &TreeNode{Val: data}
		return
	}

	if treeNode := bt.DeepSearch(bt.root, parent); treeNode !=  nil {
		if val == 0 {
			treeNode.Left = &TreeNode{Val: data}
		} else if val == 1 {
			treeNode.Right = &TreeNode{Val: data}
		}
	}
}


func (bt *BTree) DeepSearch(TreeNode *TreeNode, data interface{}) *TreeNode {
	if TreeNode == nil {
		return nil
	}

	if TreeNode.Val  == data {
		return TreeNode
	}

	if n := bt.DeepSearch(TreeNode.Left, data); n != nil {
		return n
	}

	return bt.DeepSearch(TreeNode.Right, data)
}

//func (bt *BTree) InsertByLevelSearch(parent, data interface{}, val int) {
//	if bt.root == nil {
//		bt.root = &TreeNode{data: data}
//		return
//	}
//
//	if parentTreeNode := bt.LevelSearch(parent); parentTreeNode != nil {
//		if val == 0 {
//			parentTreeNode.left = &TreeNode{data:data}
//		} else if val == 1 {
//			parentTreeNode.right = &TreeNode{data: data}
//		}
//	}
//
//}
//
//func (bt *BTree) LevelSearch(data interface{}) *TreeNode {
//	var queue []*TreeNode
//	queue =  append(queue, bt.root)
//
//	for len(queue) > 0 {
//		front := queue[0]
//		queue = queue[1:]
//
//		if front.data == data {
//			return front
//		}
//		if front.left != nil {
//			queue = append(queue, front.left)
//		}
//		if front.right != nil {
//			queue = append(queue, front.right)
//		}
//	}
//
//	return nil
//}
