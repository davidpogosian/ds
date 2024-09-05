// Package stack provides a thread-safe, generic stack implementation.
package stack

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

// Stack is a struct representing a stack. It contains a slice to store items, a comparator function
// that is used to compare elements for advanced methods such as Find, and a mutex for thread-safety.
type Stack[T any] struct {
	items []T
	comparator comparators.Comparator[T]
	mutex sync.Mutex
}

// NewEmpty creates a new empty Stack and returns a pointer to it.
// NewEmpty requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewEmpty[T any](comparator comparators.Comparator[T]) *Stack[T] {
	return &Stack[T]{comparator: comparator}
}

// NewFromSlice creates a new Stack from a slice and returns a pointer to it.
// The slice is copied prior to being handed over to the Stack.
// NewFromSlice requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewFromSlice[T any](slice []T, comparator comparators.Comparator[T]) *Stack[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Stack[T]{
		items: copiedSlice,
		comparator: comparator,
	}
}

// Pop removes and returns the top item off of the Stack.
// An error is returned if the Stack is empty.
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

// Push adds a new item to the top of the Stack.
func (stack *Stack[T]) Push(newItem T) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = append(stack.items, newItem)
}

// Peek returns the top item from the Stack.
// It returns an error if the Stack is empty.
func (stack *Stack[T]) Peek() (T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var zeroValue T
	if len(stack.items) == 0 {
		return zeroValue, fmt.Errorf("Cannot peek an empty Stack.")
	}
	return stack.items[len(stack.items) - 1], nil
}

// IsEmpty returns a bool indicating if the Stack is empty.
func (stack *Stack[T]) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items) == 0
}

// Size returns the the number of items in the Stack.
func (stack *Stack[T]) Size() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items)
}

// Clear removes all items from the Stack.
func (stack *Stack[T]) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = []T{}
}

// Find returns nonnegative int indicating the poistion of the item in the Stack.
// Returns -1 if the item is not in the Stack.
func (stack *Stack[T]) Find(item T) int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	for i := range stack.items {
		if stack.comparator(stack.items[i], item) == 0 {
			return i
		}
	}
	return -1
}

// ToSlice returns the Stack as a slice.
func (stack *Stack[T]) ToSlice() []T {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return copiedSlice
}

// Copy returns a pointer to a copy of the Stack.
func (stack *Stack[T]) Copy() *Stack[T] {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return &Stack[T]{
		items: copiedSlice,
		comparator: stack.comparator,
	}
}

// String returns the string representation of the Stack.
func (stack *Stack[T]) String() string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return fmt.Sprintf("%v", stack.items)
}
