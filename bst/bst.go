// Package bst provides a thread-safe, generic binary search tree implementation.
package bst

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

// Node struct represents a single item in the BST.
// It has fields for a key and a value. The key is used
// to determine where in the BST this node belongs.
// It also have pointers to the left and right nodes.
type Node[K any, V any] struct {
	key K
	val V
	left *Node[K, V]
	right *Node[K ,V]
}

// BST struct represents a binary search tree.
// It has a pointer to the root node, a comparator function for comparing keys,
// a field to keep track of its size, and a mutex for thread-safety.
type BST[K any, V any] struct {
	root *Node[K, V]
	comparator comparators.Comparator[K]
	size int
	mu sync.Mutex
}

// NewEmpty returns a pointer to a new empty BST.
// NewEmpty requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewEmpty[K, V any](comparator comparators.Comparator[K]) *BST[K, V] {
	return &BST[K, V]{comparator: comparator}
}

// Insert inserts a new node into the BST with the provided key and value.
func (bst *BST[K, V]) Insert(key K, value V) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	n := &Node[K, V]{
		key: key,
		val: value,
	}
	if bst.size == 0 {
		bst.root = n
	} else {
		cursor := bst.root
		for {
			comparison := bst.comparator(n.key, cursor.key)
			if comparison == -1 {
				// go left
				if cursor.left == nil {
					cursor.left = n
					break
				} else {
					cursor = cursor.left
				}
			} else {
				// go right
				if cursor.right == nil {
					cursor.right = n
					break
				} else {
					cursor = cursor.right
				}
			}
		}
	}
	bst.size++
}

// Search returns the value of the first node with the provided key.
// If no item with the provided key exists, an error is returned.
func (bst *BST[K, V]) Search(key K) (V, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	cursor := bst.root
	for cursor != nil{
		comparison := bst.comparator(key, cursor.key)
		if comparison == -1 {
			// go left
			cursor = cursor.left
		} else if comparison == 0 {
			return cursor.val, nil
		} else {
			// go right
			cursor = cursor.right
		}
	}
	var zeroValue V
	return zeroValue, fmt.Errorf("Key '%v' is not in the BST.", key)
}

// removeHelper removes a given node and returns a pointer to
// a node that will serve as its replacement.
func (bst *BST[K, V]) removeHelper(n *Node[K, V]) *Node[K, V] {
	bst.size--
	if n.left == nil && n.right == nil {
		return nil
	} else if n.left == nil {
		return n.right
	} else if n.right == nil {
		return n.left
	}
	// get greatest node from the left subtree
	replacementParent := n
	replacement := n.left
	for replacement.right != nil {
		replacementParent = replacement
		replacement = replacement.right
	}
	replacement.right = n.right
	if n.left != replacement {
		replacement.left = n.left
	}
	replacementParent.right = nil
	return replacement
}

// Remove removes the first node with the provided key and returns its value.
// If no node has the provided key, an error is returned.
func (bst *BST[K, V]) Remove(key K) (V, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	cursor := bst.root
	for cursor != nil {
		comparison := bst.comparator(key, cursor.key)
		if comparison == -1 {
			if cursor.left == nil {
				break
			} else {
				if bst.comparator(key, cursor.left.key) == 0 {
					removeNode := cursor.left
					cursor.left = bst.removeHelper(removeNode)
					return removeNode.val, nil
				} else {
					cursor = cursor.left
				}
			}
		} else if  comparison == 0 {
			// Only reachable if removing root.
			val := cursor.val
			bst.root = bst.removeHelper(bst.root)
			return val, nil
		} else {
			if cursor.right == nil {
				break
			} else {
				if bst.comparator(key, cursor.right.key) == 0 {
					removeNode := cursor.right
					cursor.right = bst.removeHelper(removeNode)
					return removeNode.val, nil
				} else {
					cursor = cursor.right
				}
			}
		}
	}
	var zeroValue V
	return zeroValue, fmt.Errorf("Key '%v' is not in the BST.", key)
}

// Size returns the number of nodes in the BST.
func (bst *BST[K, V]) Size() int {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return bst.size
}

// FindMin returns the minimum key in the BST.
// If the BST is empty, an error is returned.
func (bst *BST[K, V]) FindMin() (K, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	if bst.size == 0 {
		var zeroValue K
		return zeroValue, fmt.Errorf("Cannot find min in an empty BST.")
	}
	cursor := bst.root
	for cursor.left != nil {
		cursor = cursor.left
	}
	return cursor.key, nil
}

// FindMax returns the maximum key in the BST.
// If the BST is empty, an error is returned.
func (bst *BST[K, V]) FindMax() (K, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	if bst.size == 0 {
		var zeroValue K
		return zeroValue, fmt.Errorf("Cannot find min in an empty BST.")
	}
	cursor := bst.root
	for cursor.right != nil {
		cursor = cursor.right
	}
	return cursor.key, nil
}

// inOrderTraversal adds keys to the provided slice in an
// in-order fashion.
func (bst *BST[K, V]) inOrderTraversal(node *Node[K, V], slice *[]K) {
	if node == nil {
		return
	}
	bst.inOrderTraversal(node.left, slice)
	*slice = append(*slice, node.key)
	bst.inOrderTraversal(node.right, slice)
}

// InOrderTraversal returns a slice of the keys from the BST using in-order traversal.
func (bst *BST[K, V]) InOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	slice := []K{}
	bst.inOrderTraversal(bst.root, &slice)
	return slice
}

// preOrderTraversal adds keys to the provided slice
// in a pre-order fashion.
func (bst *BST[K, V]) preOrderTraversal(node *Node[K, V], slice *[]K) {
	if node == nil {
		return
	}
	*slice = append(*slice, node.key)
	bst.preOrderTraversal(node.left, slice)
	bst.preOrderTraversal(node.right, slice)
}

// PreOrderTraversal returns a slice of the keys from the BST using pre-order traversal.
func (bst *BST[K, V]) PreOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	slice := []K{}
	bst.preOrderTraversal(bst.root, &slice)
	return slice
}

// postOrderTraversal adds keys to the provided slice
// in a post-order fashion.
func (bst *BST[K, V]) postOrderTraversal(node *Node[K, V], slice *[]K) {
	if node == nil {
		return
	}
	bst.postOrderTraversal(node.left, slice)
	bst.postOrderTraversal(node.right, slice)
	*slice = append(*slice, node.key)
}

// PostOrderTraversal returns a slice of the keys from the BST using post-order traversal.
func (bst *BST[K, V]) PostOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	slice := []K{}
	bst.postOrderTraversal(bst.root, &slice)
	return slice
}

// height returns the height of the given node in the BST.
func (bst *BST[K, V]) height(node *Node[K, V]) int {
	if node == nil {
		return -1
	}
	left := bst.height(node.left)
	right := bst.height(node.right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// Height returns the height of the BST.
// It returns -1 if the BST is empty.
func (bst *BST[K, V]) Height() int {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return bst.height(bst.root)
}

// Clear removes all nodes from the BST.
func (bst *BST[K, V]) Clear() {
    bst.mu.Lock()
    defer bst.mu.Unlock()
    bst.root = nil
    bst.size = 0
}

// Copy returns a pointer to a copy of the BST.
func (bst *BST[K, V]) Copy() *BST[K, V] {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	var copyNode func(node *Node[K, V]) *Node[K, V]
	copyNode = func(node *Node[K, V]) *Node[K, V] {
		if node == nil {
			return nil
		}
		return &Node[K, V]{
			key:   node.key,
			val:   node.val,
			left:  copyNode(node.left),
			right: copyNode(node.right),
		}
	}
	return &BST[K, V]{
		root: copyNode(bst.root),
		size: bst.size,
		comparator: bst.comparator,
	}
}
