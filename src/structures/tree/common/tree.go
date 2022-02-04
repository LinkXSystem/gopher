package common

type Sortable interface {
}

type Node struct {
	key   Sortable
	value interface{}

	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}
