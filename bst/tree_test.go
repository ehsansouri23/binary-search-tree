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
	assert.Equal(t, bst.FindMax(), 40)
	bst2 := CreateBST(17)
	assert.Equal(t, bst2.FindMax(), 17)
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	bst := createBST()
	assert.Equal(t, bst.FindMin(), 2)
	bst2 := CreateBST(17)
	assert.Equal(t, bst2.FindMax(), 17)

}
