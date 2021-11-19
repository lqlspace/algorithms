package tree


func CreateBTree() *BTree {
	return &BTree{}
}

func (bt *BTree) Insert(parent,  data interface{}, val int) {
	if parent == nil && bt.root == nil {
		bt.root = &Node{data: data}
		return
	}

	if node := bt.DeepSearch(bt.root, parent); node !=  nil {
		if val == 0 {
			node.left  = &Node{data:data}
		} else if val == 1 {
			node.right = &Node{data: data}
		}
	}
}


func (bt *BTree) DeepSearch(node *Node, data interface{}) *Node {
	if node == nil {
		return nil
	}

	if node.data  == data {
		return node
	}

	if n := bt.DeepSearch(node.left, data); n != nil {
		return n
	}

	return bt.DeepSearch(node.right, data)
}

func (bt *BTree) InsertByLevelSearch(parent, data interface{}, val int) {
	if bt.root == nil {
		bt.root = &Node{data: data}
		return
	}

	if parentNode := bt.LevelSearch(parent); parentNode != nil {
		if val == 0 {
			parentNode.left = &Node{data:data}
		} else if val == 1 {
			parentNode.right = &Node{data: data}
		}
	}

}

func (bt *BTree) LevelSearch(data interface{}) *Node {
	var queue []*Node
	queue =  append(queue, bt.root)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front.data == data {
			return front
		}
		if front.left != nil {
			queue = append(queue, front.left)
		}
		if front.right != nil {
			queue = append(queue, front.right)
		}
	}

	return nil
}
