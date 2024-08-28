package queue

import (
	"fmt"
	"sync"
)

type Queue[T comparable] struct {
	items []T
	front int
	rear int
	mutex sync.Mutex
}

// Creates a new empty Queue.
func NewEmpty[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// Creates a new Queue from a slice.
// The slice is copied.
func NewFromSlice[T comparable](slice []T) *Queue[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Queue[T]{
		items: copiedSlice,
		front: 0,
		rear: len(copiedSlice),
	}
}

// Adds an item to the rear of the Queue.
func (queue *Queue[T]) Enqueue(newItem T) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.items = append(queue.items, newItem)
	queue.rear++
}

// Returns a bool indicating whether or not the Queue is empty.
func (queue *Queue[T]) IsEmpty() bool {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.front == queue.rear
}

// Removes and returns the item at the front of the Queue.
// Returns an error if the Queue is empty.
func (queue *Queue[T]) Dequeue() (T, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	var zeroValue T
	if queue.front == queue.rear {
		return zeroValue, fmt.Errorf("Cannot dequeue from an empty Queue.")
	}
	first := queue.items[queue.front]
	queue.front++
	return first, nil
}

// Returns the item at the front of the Queue.
// Returns an error if the Queue is empty.
func (queue *Queue[T]) Peek() (T, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	var zeroValue T
	if queue.front == queue.rear {
		return zeroValue, fmt.Errorf("Cannot peak an empty Queue.")
	}
	first := queue.items[queue.front]
	return first, nil
}

// Returns the number of items in the Queue.
func (queue *Queue[T]) Size() int {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.rear - queue.front
}

// Removes all items from the Queue.
func (queue *Queue[T]) Clear() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.items = []T{}
	queue.front = 0
	queue.rear = 0
}

// Returns a nonnegative int indicating the position of the item in the Queue.
// Returns -1 if the item is not in the Queue.
func (queue *Queue[T]) Contains(item T) int {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	for i := range queue.items {
		if queue.items[i] == item {
			return i
		}
	}
	return -1
}

// Returns a copy of the Queue.
func (queue *Queue[T]) Copy() *Queue[T] {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	copiedSlice := make([]T, len(queue.items))
	copy(copiedSlice, queue.items)
	return &Queue[T]{
		items: copiedSlice,
		front: queue.front,
		rear: queue.rear,
	}
}

// Returns the Queue as a slice.
func (queue *Queue[T]) ToSlice() []T {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	copiedSlice := make([]T, queue.rear - queue.front)
	copy(copiedSlice, queue.items[queue.front:queue.rear])
	return copiedSlice
}

// Returns the string representation of the Queue.
func (queue *Queue[T]) String() string {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return fmt.Sprintf("%v", queue.items[queue.front:queue.rear])
}
