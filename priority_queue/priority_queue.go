// Package priority_queue provides a thread-safe, generic priority queue implementation.
package priority_queue

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

// Node struct represents a single item in the priority queue.
// It consists of two fields, one for determining priority,
// and another for storing a value.
type Node[P, V any] struct {
	p P
	v V
}

// PriorityQueue struct represents a priority queue.
// It contains a slice of the Node type that is used as a heap.
// It also has a field to keep track of its size, a minHeap flag
// (to specify whether a min heap or a max heap is used),
// a comparator function for comparing priorities, and a mutex
// for thread-safety.
type PriorityQueue[P, V any] struct {
	heap []Node[P, V]
	size int
	minHeap bool
	comparator comparators.Comparator[P]
	mu sync.Mutex
}

// NewEmpty returns a pointer to a new empty PriorityQueue.
// NewEmpty requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
// NewEmpty also requires a boolean value "minHeap" to indicate whether
// to sort items in the PriorityQueue by increasing or decreasing priority.
func NewEmpty[P, V any](comparator comparators.Comparator[P], minHeap bool) *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{
		minHeap: minHeap,
		comparator: comparator,
	}
}

// heapifyUp restores the heap property of the PriorityQueue's heap by moving the
// element at the given index up to its correct position.
func (pq *PriorityQueue[P, V]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.minHeap {
			if pq.comparator(pq.heap[index].p, pq.heap[parentIndex].p) >= 0 {
				break
			}
		} else {
			if pq.comparator(pq.heap[index].p, pq.heap[parentIndex].p) <= 0 {
				break
			}
		}
		pq.heap[index], pq.heap[parentIndex] = pq.heap[parentIndex], pq.heap[index]
		index = parentIndex
	}
}

// Enqueue enqueues a given value with given priority into the heap
// of the PriorityQueue.
func (pq *PriorityQueue[P, V]) Enqueue(p P, v V) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	n := Node[P, V] {
		p: p,
		v: v,
	}
	pq.heap = append(pq.heap, n)
	pq.size++
	pq.heapifyUp(pq.size - 1)
}

// Peek returns the priority and the value of the node at the top of heap
// of the PriorityQueue.
// If the heap is empty, an error is returned.
func (pq *PriorityQueue[P, V]) Peek() (P, V, error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if pq.size == 0 {
		var zeroPriority P
		var zeroValue V
		return zeroPriority, zeroValue, fmt.Errorf("Cannot peek an empty PriorityQueue")
	}
	return pq.heap[0].p, pq.heap[0].v, nil
}

// heapifyDown restores the heap property of the PriorityQueue's heap by moving the
// element at the given index down to its correct position.
func (pq *PriorityQueue[P, V]) heapifyDown(index int) {
	for {
		leftChild := 2 * index + 1
		rightChild := 2 * index + 2
		smallestOrLargest := index
		if pq.minHeap {
			if leftChild < pq.size && pq.comparator(pq.heap[leftChild].p, pq.heap[smallestOrLargest].p) == -1 {
				smallestOrLargest = leftChild
			}
			if rightChild < pq.size && pq.comparator(pq.heap[rightChild].p, pq.heap[smallestOrLargest].p) == -1 {
				smallestOrLargest = rightChild
			}
		} else {
			if leftChild < pq.size && pq.comparator(pq.heap[leftChild].p, pq.heap[smallestOrLargest].p) == 1 {
				smallestOrLargest = leftChild
			}
			if rightChild < pq.size && pq.comparator(pq.heap[rightChild].p, pq.heap[smallestOrLargest].p) == 1 {
				smallestOrLargest = rightChild
			}
		}
		if smallestOrLargest == index {
			break
		}
		pq.heap[index], pq.heap[smallestOrLargest] = pq.heap[smallestOrLargest], pq.heap[index]
		index = smallestOrLargest
	}
}

// ExtractTop removes the node at the top of the heap
// and returns the corresponding priority and value.
// If the heap is empty, an error is returned.
func (pq *PriorityQueue[P, V]) ExtractTop() (P, V, error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if pq.size == 0 {
		var zeroPriority P
		var zeroValue V
		return zeroPriority, zeroValue, fmt.Errorf("Cannot extract top on an empty PriorityQueue")
	}
	p := pq.heap[0].p
	v := pq.heap[0].v
	pq.heap[0] = pq.heap[pq.size - 1]
	pq.size--
	pq.heap = pq.heap[:pq.size]
	pq.heapifyDown(0)
	return p, v, nil
}

// Clear removes all items from the PriorityQueue.
func (pq *PriorityQueue[P, V]) Clear() {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.heap = []Node[P, V]{}
	pq.size = 0
}

// Size returns the number of items in the PriorityQueue.
func (pq *PriorityQueue[P, V]) Size() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.size
}

// IsEmpty returns a bool indicating the emptiness of the PriorityQueue.
func (pq *PriorityQueue[P, V]) IsEmpty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.size == 0
}

// Copy returns a pointer to a copy of this PriorityQueue.
func (pq *PriorityQueue[P, V]) Copy() *PriorityQueue[P, V] {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	newHeap := make([]Node[P, V], pq.size)
	for i, node := range pq.heap {
		newHeap[i] = Node[P, V]{
			p: node.p,
			v: node.v,
		}
	}
	return &PriorityQueue[P, V]{
		heap: newHeap,
		size: pq.size,
		minHeap: pq.minHeap,
		comparator: pq.comparator,
	}
}
