package bst

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func CreateBST(value int) *BinarySearchTree {
	return &BinarySearchTree{Root: &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}}
}

func (node *Node) insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			node.Left.insert(value)
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			node.Right.insert(value)
		}
	}
}

func (bst *BinarySearchTree) Insert(value int) {
	if bst.Root != nil {
		bst.Root.insert(value)
	}
}

func preOrder(node *Node) {
	if node.Left != nil {
		preOrder(node.Left)
	}
	if node.Right != nil {
		preOrder(node.Right)
	}
	fmt.Print(node.Value, ", ")
}

func (bst *BinarySearchTree) PreOrder() {
	if bst.Root != nil {
		preOrder(bst.Root)
	}
}
