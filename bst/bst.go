package bst

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

type Node[K any, V any] struct {
	key K
	val V
	left *Node[K, V]
	right *Node[K ,V]
}

type BST[K any, V any] struct {
	root *Node[K, V]
	comparator comparators.Comparator[K]
	size int
	mu sync.Mutex
}

// Returns a new empty BST.
func NewEmpty[K, V any](comparator comparators.Comparator[K]) *BST[K, V] {
	return &BST[K, V]{comparator: comparator}
}

// Inserts a new item into the BST.
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

// Returns the value of the first node with the inputted key.
// If no item with the inputted key exists, an error is returned.
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

// Removes give node and returns its replacement.
func (bst *BST[K, V]) removeHelper(n *Node[K, V]) *Node[K, V] {
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
	bst.size--
	return replacement
}

// Removes the first node with the inputted key and returns its value.
// If no node has the inputted key, an error is returned.
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

// Returns the number of nodes in the BST.
func (bst *BST[K, V]) Size() int {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return bst.size
}

// Returns the minimum key in the BST.
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

// Returns the maximum key in the BST.
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

func (bst *BST[K, V]) inOrderTraversal(node *Node[K, V], slice *[]K) {
	if node == nil {
		return
	}
	bst.inOrderTraversal(node.left, slice)
	*slice = append(*slice, node.key)
	bst.inOrderTraversal(node.right, slice)
}

// Returns a slice of the keys from the BST using in-order traversal.
func (bst *BST[K, V]) InOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	slice := []K{}
	bst.inOrderTraversal(bst.root, &slice)
	return slice
}

func (bst *BST[K, V]) preOrderTraversal(node *Node[K, V], slice *[]K) {
	if node == nil {
		return
	}
	*slice = append(*slice, node.key)
	bst.inOrderTraversal(node.left, slice)
	bst.inOrderTraversal(node.right, slice)
}

// Returns a slice of the keys from the BST using pre-order traversal.
func (bst *BST[K, V]) PreOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	slice := []K{}
	bst.preOrderTraversal(bst.root, &slice)
	return slice
}

func (bst *BST[K, V]) postOrderTraversal(node *Node[K, V], slice *[]K) {
	if node == nil {
		return
	}
	bst.inOrderTraversal(node.left, slice)
	bst.inOrderTraversal(node.right, slice)
	*slice = append(*slice, node.key)
}

// Returns a slice of the keys from the BST using post-order traversal.
func (bst *BST[K, V]) PostOrderTraversal() []K {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	slice := []K{}
	bst.postOrderTraversal(bst.root, &slice)
	return slice
}

func (bst *BST[K, V]) height(node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	left := bst.height(node.left)
	right := bst.height(node.right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// Returns the height of the BST.
func (bst *BST[K, V]) Height() int {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return bst.height(bst.root)
}

// Removes all nodes from the BST.
func (bst *BST[K, V]) Clear() {
    bst.mu.Lock()
    defer bst.mu.Unlock()
    bst.root = nil
    bst.size = 0
}

// Copy returns a copy of the BST.
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
