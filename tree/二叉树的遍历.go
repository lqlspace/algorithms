package tree

import (
	"fmt"
)

func (bt *BTree) InOrderTraverse(node *Node) {
	if  node == nil {
		return
	}
	bt.InOrderTraverse(node.left)
	fmt.Printf("%s ", node.data.(string))
	bt.InOrderTraverse(node.right)
}

func (bt *BTree) PreOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%s ", node.data.(string))
	bt.PreOrderTraverse(node.left)
	bt.PreOrderTraverse(node.right)
}

func (bt *BTree) PostOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	bt.PostOrderTraverse(node.left)
	bt.PostOrderTraverse(node.right)
	fmt.Printf("%s ", node.data.(string))
}

func (bt *BTree) LevelTraverse() {
	var queue []*Node
	queue = append(queue, bt.root)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front.left != nil {
			queue = append(queue, front.left)
		}
		if front.right != nil {
			queue = append(queue, front.right)
		}
		fmt.Printf("%v ", front.data)
	}
}
