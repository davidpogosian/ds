package bst

import (

)

type BSTNode[T comparable] struct {
	val T
	left *BSTNode[T]
	right *BSTNode[T]
}

type BST[T comparable] struct {
	root *BSTNode[T]
}
