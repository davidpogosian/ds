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
// Duplicate keys are ok.
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

// InOrderTraversal returns a slice of the keys from the BST using in-order traversal.
func (bst *BST[K, V]) InOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	var slice []K
	stack := []*Node[K, V]{}
	current := bst.root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		slice = append(slice, current.key)
		current = current.right
	}
	return slice
}

// PreOrderTraversal returns a slice of the keys from the BST using pre-order traversal.
func (bst *BST[K, V]) PreOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	var slice []K
	if bst.root == nil {
		return slice
	}
	stack := []*Node[K, V]{bst.root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		slice = append(slice, node.key)
		if node.right != nil {
			stack = append(stack, node.right)
		}
		if node.left != nil {
			stack = append(stack, node.left)
		}
	}
	return slice
}

// PostOrderTraversal returns a slice of the keys from the BST using post-order traversal.
func (bst *BST[K, V]) PostOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	var slice []K
	if bst.root == nil {
		return slice
	}
	s1 := []*Node[K, V]{bst.root}
	s2 := []*Node[K, V]{}
	for len(s1) > 0 {
		node := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, node)
		if node.left != nil {
			s1 = append(s1, node.left)
		}
		if node.right != nil {
			s1 = append(s1, node.right)
		}
	}
	for len(s2) > 0 {
		node := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		slice = append(slice, node.key)
	}
	return slice
}

// nodeLevel represents a node and its level in the BST during BFS traversal.
type nodeLevel[K, V any] struct {
	node  *Node[K, V]
	level int
}

// Height returns the height of the BST.
// It returns -1 if the BST is empty.
func (bst *BST[K, V]) Height() int {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	if bst.root == nil {
		return -1
	}
	queue := []nodeLevel[K, V]{}
	queue = append(queue, nodeLevel[K, V]{node: bst.root, level: 0})
	height := 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		height = front.level
		if front.node.left != nil {
			queue = append(queue, nodeLevel[K, V]{
				node: front.node.left,
				level: front.level + 1,
			})
		}
		if front.node.right != nil {
			queue = append(queue, nodeLevel[K, V]{
				node: front.node.right,
				level: front.level + 1,
			})
		}
	}
	return height
}

// Clear removes all nodes from the BST.
func (bst *BST[K, V]) Clear() {
    bst.mu.Lock()
    defer bst.mu.Unlock()
    bst.root = nil
    bst.size = 0
}

// originalAndCopy represents a node and its corresponding node from
// a different BST (for copying).
type originalAndCopy[K, V any] struct {
	original  *Node[K, V]
	copy *Node[K, V]
}

// Copy returns a pointer to a copy of the BST.
func (bst *BST[K, V]) Copy() *BST[K, V] {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	if bst.root == nil {
		return &BST[K, V]{
			comparator: bst.comparator,
		}
	}
	copyNode := func(node *Node[K, V]) *Node[K, V] {
		if node == nil {
			return nil
		}
		return &Node[K, V]{
			key:   node.key,
			val:   node.val,
			left:  nil,
			right: nil,
		}
	}
	copiedRoot := copyNode(bst.root)
	var queue []originalAndCopy[K, V]
	queue = append(queue, originalAndCopy[K, V]{
		original: bst.root,
		copy: copiedRoot,
	})
	for len(queue) > 0 {
		original := queue[0].original
		copy := queue[0].copy
		queue = queue[1:]
		if original.left != nil {
			copyLeftChild := copyNode(original.left)
			copy.left = copyLeftChild
			queue = append(queue, originalAndCopy[K, V]{
				original: original.left,
				copy: copyLeftChild,
			})
		}
		if original.right != nil {
			copyRightChild := copyNode(original.right)
			copy.right = copyRightChild
			queue = append(queue, originalAndCopy[K, V]{
				original: original.right,
				copy: copyRightChild,
			})
		}
	}
	return &BST[K, V]{
		root:       copiedRoot,
		size:       bst.size,
		comparator: bst.comparator,
	}
}
