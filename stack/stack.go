package stack

import (
	"fmt"
	"sync"
)

type Stack[T comparable] struct {
	items []T
	mutex sync.Mutex
}

func NewEmpty[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// also shallow
func NewFromSlice[T comparable](slice []T) *Stack[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Stack[T]{items: copiedSlice}
}

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

func (stack *Stack[T]) Push(newItem T) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = append(stack.items, newItem)
}

func (stack *Stack[T]) Peek() (T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var zeroValue T
	if len(stack.items) == 0 {
		return zeroValue, fmt.Errorf("Cannot peek an empty Stack.")
	}
	return stack.items[len(stack.items) - 1], nil
}

func (stack *Stack[T]) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items) == 0
}

func (stack *Stack[T]) Size() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items)
}

func (stack *Stack[T]) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = []T{}
}

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

// shallow
func (stack *Stack[T]) ToSlice() []T {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return copiedSlice
}

// Shallow copy, should I implement deep copy?
func (stack *Stack[T]) Copy() *Stack[T] {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return &Stack[T]{items: copiedSlice}
}

func (stack *Stack[T]) String() string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return fmt.Sprintf("%v", stack.items)
}
