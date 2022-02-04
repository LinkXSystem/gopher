package lsm

type LSMTree struct {
	root *LSMNode
}

type LSMNode struct {
	left  *LSMNode
	right *LSMNode
}
