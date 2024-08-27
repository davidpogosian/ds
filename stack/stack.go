package stack

import (
	"fmt"
	"sync"
)

type Stack[T comparable] struct {
	items []T
	mutex sync.Mutex
}

// Creates a new empty Stack.
func NewEmpty[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Creates a new Stack from a slice. The slice is copied.
func NewFromSlice[T comparable](slice []T) *Stack[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Stack[T]{items: copiedSlice}
}

// Removes and returns the top item off of the Stack.
// Returns an error if the Stack is empty.
func (stack *Stack[T]) Pop() (T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var zeroValue T
	if len(stack.items) == 0 {
		return zeroValue, fmt.Errorf("Cannot pop from an empty Stack.")
	}
	last := stack.items[len(stack.items) - 1]
	stack.items = stack.items[:len(stack.items) - 1]
	return last, nil
}

// Adds a new item to the top of the Stack.
func (stack *Stack[T]) Push(newItem T) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = append(stack.items, newItem)
}

// Returns the top item from the Stack.
// Returns an error if the Stack is empty.
func (stack *Stack[T]) Peek() (T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var zeroValue T
	if len(stack.items) == 0 {
		return zeroValue, fmt.Errorf("Cannot peek an empty Stack.")
	}
	return stack.items[len(stack.items) - 1], nil
}

// Returns bool indicating if the Stack is empty.
func (stack *Stack[T]) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items) == 0
}

// Returns the the number of items in the Stack.
func (stack *Stack[T]) Size() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items)
}

// Removes all items from the Stack.
func (stack *Stack[T]) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = []T{}
}

// Returns nonnegative int indicating poistion of item in the Stack.
// Returns -1 if the item could not be found.
func (stack *Stack[T]) Contains(item T) int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	for i := range stack.items {
		if stack.items[i] == item {
			return i
		}
	}
	return -1
}

// Returns the Stack as a slice.
func (stack *Stack[T]) ToSlice() []T {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return copiedSlice
}

// Returns a copy of the Stack.
// The Stack's underlying slice is copied, then returned.
func (stack *Stack[T]) Copy() *Stack[T] {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return &Stack[T]{items: copiedSlice}
}

// Returns a string representation of the Stack.
func (stack *Stack[T]) String() string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return fmt.Sprintf("%v", stack.items)
}
