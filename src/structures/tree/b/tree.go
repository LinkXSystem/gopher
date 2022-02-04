package b

// B-Tree 的实现

type BTree struct {
	root *BNode
}

type BNode struct {
	left  *BNode
	right *BNode
}

// B+-Tree 的实现

