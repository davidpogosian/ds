package priorityqueue

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

type Node[P, V any] struct {
	p P
	v V
}

type PriorityQueue[P, V any] struct {
	heap []Node[P, V]
	size int
	minHeap bool
	comparator comparators.Comparator[P]
	mu sync.Mutex
}

// Returns a new empty PriorityQueue.
func NewEmpty[P, V any](comparator comparators.Comparator[P], minHeap bool) *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{
		minHeap: minHeap,
		comparator: comparator,
	}
}

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

// Enqueues a given value with given priority.
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

// Returns the value at the top of the heap.
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

// Removes and returns the value at the top of the heap.
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

// Removes all items from the PriorityQueue.
func (pq *PriorityQueue[P, V]) Clear() {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.heap = []Node[P, V]{}
	pq.size = 0
}

// Returns the number of items in the PriorityQueue.
func (pq *PriorityQueue[P, V]) Size() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.size
}

// Returns the number of items in the PriorityQueue.
func (pq *PriorityQueue[P, V]) IsEmpty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.size == 0
}
