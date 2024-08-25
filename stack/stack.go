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

func NewFromSlice[T comparable](slice []T) (*Stack[T], error) {
	if slice == nil {
		return nil, fmt.Errorf("Cannot initialize a Stack with a nil slice.")
	}
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &Stack[T]{items: copiedSlice}, nil
}

func (stack *Stack[T]) Pop() (T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var zeroValue T
	if stack == nil {
		return zeroValue, fmt.Errorf("Cannot pop from a nil Stack pointer.")
	}
	if len(stack.items) == 0 {
		return zeroValue, fmt.Errorf("Cannot pop from an empty Stack.")
	}
	last := stack.items[len(stack.items) - 1]
	stack.items = stack.items[:len(stack.items) - 1]
	return last, nil
}

func (stack *Stack[T]) Push(newItem T) error {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return fmt.Errorf("Cannot push onto a nil Stack pointer.")
	}
	stack.items = append(stack.items, newItem)
	return nil
}

func (stack *Stack[T]) Peek() (T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var zeroValue T
	if stack == nil {
		return zeroValue, fmt.Errorf("Cannot peek a nil Stack pointer.")
	}
	if len(stack.items) == 0 {
		return zeroValue, fmt.Errorf("Cannot peek an empty Stack.")
	}
	return stack.items[len(stack.items) - 1], nil
}

func (stack *Stack[T]) IsEmpty() (bool, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return false, fmt.Errorf("Cannot check the emptiness of a nil Stack pointer.")
	}
	return len(stack.items) == 0, nil
}

func (stack *Stack[T]) Size() (int, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return -1, fmt.Errorf("Cannot get the size a nil Stack pointer.")
	}
	return len(stack.items), nil
}

func (stack *Stack[T]) Clear() error {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return fmt.Errorf("Cannot clear a nil Stack pointer.")
	}
	stack.items = []T{}
	return nil
}

func (stack *Stack[T]) Contains(item T) (int, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return -1, fmt.Errorf("Cannot search a nil Stack pointer.")
	}
	for i := range stack.items {
		if stack.items[i] == item {
			return i, nil
		}
	}
	return -1, nil
}

func (stack *Stack[T]) ToSlice() ([]T, error) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return nil, fmt.Errorf("Cannot convert a nil Stack pointer to a slice.")
	}
	copiedSlice := make([]T, len(stack.items))
	copy(copiedSlice, stack.items)
	return copiedSlice, nil
}

func (stack *Stack[T]) String() string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack == nil {
		return "nil Stack pointer"
	}
	return fmt.Sprintf("%v", stack.items)
}
