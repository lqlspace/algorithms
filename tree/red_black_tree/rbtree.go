package rbtree

import ("log")

type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

func(rbtree *RBTree) GetRoot() *RBNode {
	return rbtree.root
}

func (rbtree *RBTree) Insert(data int64) {
	if rbtree.root == nil {
		rootnode := NewRBNode(data)
		rootnode.color = BLACK
		rbtree.root = rootnode
	} else {
		rbtree.insertNode(rbtree.root, data)
	}
}


func (rbtree *RBTree) insertNode(pnode *RBNode, data int64) {
	if data <= pnode.value {
		if pnode.left != nil {
			rbtree.insertNode(pnode.left, data)
		} else {
			tmpnode := NewRBNode(data)
			tmpnode.parent = pnode
			pnode.left = tmpnode
			rbtree.insertCheck(tmpnode)
		}
	} else {
		if pnode.right != nil {
			rbtree.insertNode(pnode.right, data)
		} else {
			tmpnode := NewRBNode(data)
			tmpnode.parent = pnode
			pnode.right = tmpnode
			rbtree.insertCheck(tmpnode)
		}
	}
}

func (rbtree *RBTree) rotateLeft(node *RBNode) {
	tmproot, err := node.rotate(LEFTROTATE)
	if err != nil {
		log.Fatalf("rotate left error: %v", err)
	}
	rbtree.root = tmproot
}


func (rbtree *RBTree) rotateRight(node *RBNode) {
	tmproot, err := node.rotate(RIGHTROTATE)
	if err != nil {
		log.Fatalf("rotate right error: %v", err)
	}
	rbtree.root = tmproot
}

func (rbtree *RBTree) insertCheck(node *RBNode) {
	if node.parent == nil {
		rbtree.root = node
		rbtree.root.color = BLACK
		return
	}

	if node.parent.color == RED {
		if node.getUncle() != nil && node.getUncle().color == RED {
			//父节点和叔叔节点都是红色，则转为黑色
			node.parent.color = BLACK
			node.getUncle().color = BLACK
			node.getGrandParent().color = RED

			rbtree.insertCheck(node.getGrandParent())
		} else {
			//父节点为红色，父节点的兄弟节点不存在或为黑色
			isleft := node == node.parent.left
			isparentleft := node.parent == node.getGrandParent().left
			if isleft && isparentleft {
				node.parent.color = BLACK
				node.getGrandParent().color = RED
				rbtree.rotateRight(node.getGrandParent())
			} else if isleft && !isparentleft {
				rbtree.rotateRight(node.parent)
				rbtree.rotateLeft(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else if !isleft && isparentleft {
				rbtree.rotateLeft(node.parent)
				rbtree.rotateRight(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else {
				node.parent.color = BLACK
				node.getGrandParent().color = RED
				rbtree.rotateLeft(node.getGrandParent())
			}
		}
	}
}


func printTreeInLog(pnode *RBNode, front string) {
	if pnode != nil {
		var colorstr string
		if pnode.color == RED {
			colorstr = "红"
		} else {
			colorstr = "黑"
		}
		log.Printf(front+"%d,%s\n", pnode.value, colorstr)
		printTreeInLog(pnode.left, front+"-(l)|")
		printTreeInLog(pnode.right, front+"-(r)|")
	}
}


func (rbtree *RBTree)RBTreeSearch(pnode *RBNode, value int64) bool {
	if pnode == nil {
		return false
	}

	if value < pnode.value {
		rbtree.RBTreeSearch(pnode.left, value)
	} else if value > pnode.value {
		rbtree.RBTreeSearch(pnode.right, value)
	}

	return true
}
