package rbtree

import (
	"log"
	"testing"
)

type State bool

const (
	LEFTCHILD State = true
	RIGHTCHILD State = false
)

func addSon(value int64, parent *RBNode, isleft State) *RBNode {
	son := NewRBNode(value)
	son.parent = parent
	if isleft {
		parent.left = son
	} else {
		parent.right = son
	}

	return son
}

func travers(root *RBNode) []int64 {
	if root == nil {
		return nil
	}

	var rbnodes []int64
	rbnodes = append(rbnodes, travers(root.left)...)
	rbnodes = append(rbnodes, root.value)
	rbnodes = append(rbnodes, travers(root.right)...)

	return rbnodes
}

func createTree() (*RBNode, error) {
	root := NewRBNode(4)
	l := addSon(2, root, LEFTCHILD)
	r := addSon(6, root, RIGHTCHILD)
	addSon(1, l, LEFTCHILD)
	addSon(3, l, RIGHTCHILD)
	addSon(5, r, LEFTCHILD)
	addSon(7, r, RIGHTCHILD)

	nodes := travers(root)
	log.Println(nodes)

	return root, nil
}

func TestGetGrandParent(t *testing.T) {
	root, err := createTree()
	if err != nil {
		log.Println(err)
		return
	}

	n := root.left.left.getGrandParent()
	if n.value == root.value {
		t.Log("getGrandParent() function OK!")
	} else {
		t.Errorf("Error: expected: %d, acutal: %d", root.value, n.value)
	}


}

func TestGetSibling(t *testing.T) {
	root, err := createTree()
	if err != nil {
		log.Println(err)
		return
	}

	n := root.left.left.getSibling()
	if n.value == root.left.right.value {
		t.Log("getSibling() function ok!")
	} else {
		t.Errorf("Error: expected: %d, actual: %d", root.left.right.value, n.value)
	}

}


func TestGetUncle(t *testing.T) {
	root, err := createTree()
	if err != nil {
		log.Println(err)
		return
	}

	n := root.left.left.getUncle()
	if n.value == root.right.value {
		t.Log("getUncle() function ok!")
	} else {
		t.Errorf("Error: expected: %d, actual: %d", root.right.value, n.value)
	}
}



func TestRotate(t *testing.T) {
	root, err := createTree()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = root.left.left.rotate(LEFTROTATE)
	if err.Error() == "左旋右节点不能为空" {
		t.Logf("rotate left valid!!!")
	} else {
		t.Errorf(err.Error())
	}

	_, err = root.left.left.rotate(RIGHTROTATE)
	if err.Error() == "右旋左节点不能为空" {
		t.Logf("rotate right valid!!!")
	} else {
		t.Errorf(err.Error())
	}
	
	n1, err := root.left.rotate(LEFTROTATE)
	if err != nil {
		t.Errorf(err.Error())
	}

	if n1.left.value == 3 {
		t.Logf("rotate left successed!!!")
	} else {
		t.Errorf("Error: expected : %d, actual: %d", 3, n1.left.value)
	}

	n2, err := root.left.rotate(RIGHTROTATE)
	if err != nil {
		t.Errorf(err.Error())
	}

	if n2.left.value == 2 {
		t.Logf("rotate left successed!!!")
	} else {
		t.Errorf("Error: expected : %d, actual: %d", 2, n2.left.value)
	}
}
