package tree

type BTree struct {
	root *Node
}



type Node struct {
	data interface{}
	left *Node
	right *Node
}

