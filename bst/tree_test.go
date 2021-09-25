package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBST() *BinarySearchTree {
	bst := CreateBST(20)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(40)
	return bst
}

func TestBinarySearchTree_FindMax(t *testing.T) {
	bst := createBST()
	assert.Equal(t, bst.FindMax(), 20)
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	bst := createBST()
	assert.Equal(t, bst.FindMin(), 2)
}
