package avl

func abs(target int) int {
	if target >= 0 {
		return target
	}

	return -target
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x int, y int) int {
	if x > y {
		return y
	}

	return x
}

type AVLNode struct {
	key    int
	height int
	value  interface{}

	left  *AVLNode
	right *AVLNode
}

func (node *AVLNode) Key() int {
	if node != nil {
		return node.key
	}

	return 0
}

func (node *AVLNode) Height() int {
	if node != nil {
		return node.height
	}

	return 0
}

func (node *AVLNode) Put(key int, value interface{}) (_ *AVLNode, inserted bool) {
	if node == nil {
		return &AVLNode{
			key:    key,
			value:  value,
			height: 1,
		}, true
	}

	if node.key == key {
		node.value = value
		return node, false
	}

	status := key < node.key

	switch status {
	case true:
		node.left, inserted = node.left.Put(key, value)
	case false:
		node.right, inserted = node.right.Put(key, value)
	}

	if inserted {
		node.height += 1
		return node.balance(), inserted
	}

	return node, inserted
}

func (node *AVLNode) getBalanceFactor() int {
	if node == nil {
		return 0
	}

	return node.left.Height() - node.right.Height()
}

func (node *AVLNode) getHeight() int {
	if node == nil {
		return 0
	}

	return max(node.left.getHeight(), node.right.getHeight()) + 1
}

func (node *AVLNode) balance() *AVLNode {
	factor := node.getBalanceFactor()

	// LL

	if factor > 1 && node.left.getBalanceFactor() > 0 /*    */ {
		return node.rotateRight()
	}

	// LR

	if factor > 1 && node.left.getBalanceFactor() <= 0 /*   */ {
		node.left = node.left.rotateLeft()

		return node.rotateRight()
	}

	// RR

	if factor < -1 && node.right.getBalanceFactor() <= 0 /* */ {
		return node.rotateLeft()
	}

	// RL

	if factor < -1 && node.right.getBalanceFactor() > 0 /*  */ {
		node.right = node.right.rotateRight()

		return node.rotateLeft()
	}

	return node
}

func (node *AVLNode) rotateRight() *AVLNode {
	left := node.left

	node.left = left.right
	left.right = node

	left.height = max(left.left.getHeight(), left.right.getHeight()) + 1

	if left.right != nil {
		left.right.height = max(left.right.left.getHeight(), left.right.right.getHeight()) + 1
	}

	return left
}

func (node *AVLNode) rotateLeft() *AVLNode {
	right := node.right

	node.right = right.left
	right.left = node

	right.height = max(right.left.getHeight(), right.right.getHeight()) + 1

	if right.left != nil {
		right.left.height = max(right.left.left.getHeight(), right.left.right.getHeight()) + 1
	}

	return right
}

type AVLTree struct {
	root *AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (tree *AVLTree) Root() *AVLNode {
	return tree.root
}

func (tree *AVLTree) Put(key int, value interface{}) (err error) {
	tree.root, _ = tree.root.Put(key, value)
	return nil
}

func (tree *AVLTree) Get(key int) (_ interface{}, err error) {
	return tree.root.value, nil
}

func (tree *AVLTree) Has(key int) bool {
	return false
}

func (tree *AVLTree) Earse(key int) (_ interface{}, err error) {
	return tree.root.value, nil
}

func (tree AVLTree) Height() int  {
	return 0
}
