package main

import (
	"binary-search-tree/bst"
)

func main() {
	binarySearchTree := bst.CreateBST(5)
	binarySearchTree.Insert(10)
	binarySearchTree.Insert(1)
	binarySearchTree.PreOrder()
}
