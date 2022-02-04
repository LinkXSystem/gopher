package avl

import (
	"log"
	"testing"
)

func TestAVLTreeRotateRight(t *testing.T) {
	tree := NewAVLTree()

	tree.Put(25, string("this is node"))
	tree.Put(15, string("this is node"))
	tree.Put(10, string("this is node"))

	if tree.Root().Key() != 15 {
		t.Errorf("the rotate is error !")
		t.Errorf("the key    is error !")
	}

	if tree.Root().Height() != 2 {
		t.Errorf("the rotate is error !")
		t.Errorf("the height is error !")
	}

}

func TestAVLTreeRotateRightLeft(t *testing.T) {
	tree := NewAVLTree()

	tree.Put(25, string("this is node"))
	tree.Put(15, string("this is node"))
	tree.Put(16, string("this is node"))

	if tree.Root().Key() != 16 {
		t.Errorf("the rotate is error !")
		t.Errorf("the key    is error !")
	}

	if tree.Root().Height() != 2 {
		t.Errorf("the rotate is error !")
		t.Errorf("the height is error !")
	}

}

func TestAVLTreeRotateLeft(t *testing.T) {
	tree := NewAVLTree()

	tree.Put(11, string("this is node"))
	tree.Put(12, string("this is node"))
	tree.Put(13, string("this is node"))

	if tree.Root().Key() != 12 {
		t.Errorf("the rotate is error !")
		t.Errorf("the key    is error !")
	}

	if tree.Root().Height() != 2 {
		t.Errorf("the rotate is error !")
		t.Errorf("the height is error !")
	}

}

func TestAVLTreeRotateLeftRight(t *testing.T) {
	tree := NewAVLTree()

	tree.Put(25, string("this is node"))
	tree.Put(42, string("this is node"))
	tree.Put(36, string("this is node"))

	if tree.Root().Key() != 36 {
		t.Errorf("the rotate is error !")
		t.Errorf("the key    is error !")
	}

	if tree.Root().Height() != 2 {
		t.Errorf("the rotate is error !")
		t.Errorf("the height is error !")
	}

}

func TestAVLTreeRotate(t *testing.T) {

	tree := NewAVLTree()

	tree.Put(25, string("this is node"))
	tree.Put(15, string("this is node"))
	tree.Put(10, string("this is node"))

	if tree.Root().Key() != 15 {
		t.Errorf("the rotate is error !")
	}

	tree.Put(11, string("this is node"))
	tree.Put(12, string("this is node"))
	tree.Put(13, string("this is node"))

	log.Println("the tree is ", tree)
}
