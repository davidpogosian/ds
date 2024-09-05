// Package queue provides a thread-safe, generic queue implementation.
package queue

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

// Queue is a struct representing a queue. It contains a circular slice to store items, pointers to the
// front and the rear of the queue, a field to keep track of the size, a comparator function,
// and a mutex for thread-safety.
type Queue[T any] struct {
	items []T
	front int
	rear int
	size int
	comparator comparators.Comparator[T]
	mutex sync.Mutex
}

// NewEmpty creates a new empty Queue and returns a pointer to it.
// NewEmpty requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewEmpty[T any](comparator comparators.Comparator[T]) *Queue[T] {
	return &Queue[T]{items: make([]T, 4), comparator: comparator}
}

// NewFromSlice creates a new Queue from a slice and returns a pointer to it.
// The slice is copied prior to being handed over to the Queue.
// NewFromSlice requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewFromSlice[T any](slice []T, comparator comparators.Comparator[T]) *Queue[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Queue[T]{
		items: copiedSlice,
		front: 0,
		rear: 0,
		size: len(copiedSlice),
		comparator: comparator,
	}
}

// grow doubles the capacity of the Queue and copies over existing items.
func (queue *Queue[T]) grow() {
	newCapacity := len(queue.items) * 2
	newItems := make([]T, newCapacity)
	if queue.front < queue.rear {
        copy(newItems, queue.items[queue.front:queue.rear])
    } else {
        copy(newItems, queue.items[queue.front:])
        copy(newItems[len(queue.items) - queue.front:], queue.items[:queue.rear])
    }
	queue.front = 0
	queue.rear = queue.size
	queue.items = newItems
}

// Enqueue adds an item to the rear of the Queue.
func (queue *Queue[T]) Enqueue(newItem T) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if queue.size == len(queue.items) {
		queue.grow()
	}
	queue.items[queue.rear] = newItem
	queue.rear = (queue.rear + 1) % len(queue.items)
	queue.size++
}

// IsEmpty returns a bool indicating whether or not the Queue is empty.
func (queue *Queue[T]) IsEmpty() bool {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.size == 0
}

// Dequeue removes and returns the item at the front of the Queue.
// It returns an error if the Queue is empty.
func (queue *Queue[T]) Dequeue() (T, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if queue.size == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("Cannot dequeue from an empty Queue.")
	}
	first := queue.items[queue.front]
	queue.front = (queue.front + 1) % len(queue.items)
	queue.size--
	return first, nil
}

// Peek returns the item at the front of the Queue.
// It returns an error if the Queue is empty.
func (queue *Queue[T]) Peek() (T, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	var zeroValue T
	if queue.size == 0 {
		return zeroValue, fmt.Errorf("Cannot peak an empty Queue.")
	}
	first := queue.items[queue.front]
	return first, nil
}

// Size returns the number of items in the Queue.
func (queue *Queue[T]) Size() int {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.size
}

// Clear removes all items from the Queue.
func (queue *Queue[T]) Clear() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.front = 0
	queue.rear = 0
	queue.size = 0
}

// Find returns a nonnegative int indicating the position of the item in the Queue.
// It returns -1 if the item is not in the Queue.
func (queue *Queue[T]) Find(item T) int {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	traversed := 0
	for i := queue.front; traversed != queue.size; i = (i + 1) % len(queue.items) {
		if queue.comparator(queue.items[i], item) == 0 {
			return traversed
		}
		traversed++
	}
	return -1
}

// Returns a pointer to a copy of the Queue.
func (queue *Queue[T]) Copy() *Queue[T] {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	copiedSlice := make([]T, len(queue.items))
	copy(copiedSlice, queue.items)
	return &Queue[T]{
		items: copiedSlice,
		front: queue.front,
		rear: queue.rear,
		size: queue.size,
		comparator: queue.comparator,
	}
}

// ToSlice returns the Queue as a slice.
func (queue *Queue[T]) ToSlice() []T {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	copiedSlice := make([]T, queue.size)
	if queue.front < queue.rear {
        copy(copiedSlice, queue.items[queue.front:queue.rear])
    } else {
        copy(copiedSlice, queue.items[queue.front:])
        copy(copiedSlice[len(queue.items) - queue.front:], queue.items[:queue.rear])
    }
	return copiedSlice
}

// String returns the string representation of the Queue.
func (queue *Queue[T]) String() string {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if queue.size == 0 {
		return "[]"
	}
	if queue.front < queue.rear {
        return fmt.Sprintf("%v", queue.items[queue.front:queue.rear])
	}
	concatenated := append(queue.items[queue.front:], queue.items[:queue.rear]...)
	return fmt.Sprintf("%v", concatenated)
}
