package queue

import (
	"fmt"
	"sync"
)

type Queue[T comparable] struct {
	items []T
	front int
	rear int
	size int
	mutex sync.Mutex
}

// Creates a new empty Queue.
func NewEmpty[T comparable]() *Queue[T] {
	return &Queue[T]{items: make([]T, 4)}
}

// Creates a new Queue from a slice.
// The slice is copied.
func NewFromSlice[T comparable](slice []T) *Queue[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Queue[T]{
		items: copiedSlice,
		front: 0,
		rear: 0,
		size: len(copiedSlice),
	}
}

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

// Adds an item to the rear of the Queue.
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

// Returns a bool indicating whether or not the Queue is empty.
func (queue *Queue[T]) IsEmpty() bool {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.size == 0
}

// Removes and returns the item at the front of the Queue.
// Returns an error if the Queue is empty.
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

// Returns the item at the front of the Queue.
// Returns an error if the Queue is empty.
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

// Returns the number of items in the Queue.
func (queue *Queue[T]) Size() int {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.size
}

// Removes all items from the Queue.
func (queue *Queue[T]) Clear() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.front = 0
	queue.rear = 0
	queue.size = 0
}

// Returns a nonnegative int indicating the position of the item in the Queue.
// Returns -1 if the item is not in the Queue.
func (queue *Queue[T]) Contains(item T) int {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	traversed := 0
	for i := queue.front; traversed != queue.size; i = (i + 1) % len(queue.items) {
		if queue.items[i] == item {
			return traversed
		}
		traversed++
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
		size: queue.size,
	}
}

// Returns the Queue as a slice.
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

// Returns the string representation of the Queue.
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
