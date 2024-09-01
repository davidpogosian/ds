package bst

import (

)

type BSTNode[K any, V any] struct {
	key K
	val V
	left *BSTNode[K, V]
	right *BSTNode[K ,V]
}

type BST[K any, V any] struct {
	root *BSTNode[K, V]
}
