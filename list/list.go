// Package list provides a thread-safe, generic doubly-linked list implementation.
package list

import (
	"fmt"
	"sync"

	"github.com/davidpogosian/ds/comparators"
)

// node struct represents a single item in the List.
// It has a field for a value, and pointers to the previous and the next node.
type node[T any] struct {
	val T
	next *node[T]
	prev *node[T]
}

// List struct represents a double-linked list.
// It has pointers to the front and the back of the list.
// A field to keep track of the size of the list.
// A comparator function to compare elements.
// And a mutex for thread-safety.
type List[T any] struct {
	front *node[T]
	back *node[T]
	size int
	comparator comparators.Comparator[T]
	mu sync.Mutex
}

// NewEmpty returns a pointer to a new empty List.
// NewEmpty requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewEmpty[T any](comparator comparators.Comparator[T]) *List[T] {
	return &List[T]{comparator: comparator}
}

// NewFromSlice returns a pointer to a new List initialized with a slice.
// NewFromSlice requires a comparator function to compare elements.
// For built-in types, the comparators package provides ready-made comparators
// (e.g., comparators.CompareInt for int).
// Custom types will require a user-defined comparator.
func NewFromSlice[T any](slice []T, comparator comparators.Comparator[T]) *List[T] {
	l := List[T]{comparator: comparator}
	for _, item := range slice {
		l.InsertBack(item)
	}
	return &l
}

// insertFront inserts new item at the front of the List.
func (l *List[T]) insertFront(newItem T) {
	n := &node[T]{val: newItem}
	if l.size == 0 {
		l.front = n
		l.back = n
	} else {
		n.next = l.front
		l.front.prev = n
		l.front = n
	}
	l.size++
}

// InsertFront inserts new item at the front of the List.
func (l *List[T]) InsertFront(newItem T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.insertFront(newItem)
}

// insertBack inserts new item at the back of the list.
func (l *List[T]) insertBack(newItem T) {
	n := &node[T]{val: newItem}
	if l.size == 0 {
		l.front = n
		l.back = n
	} else {
		n.prev = l.back
		l.back.next = n
		l.back = n
	}
	l.size++
}

// InsertBack inserts new item at the back of the List.
func (l *List[T]) InsertBack(newItem T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.insertBack(newItem)
}

// InsertPosition inserts new item at the specified position.
// If the position is invalid (aka if position < 0 || position > List.size), an error is returned.
func (l *List[T]) InsertPosition(newItem T, position int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if position < 0 || position > l.size {
		return fmt.Errorf("Cannot insert into a List of size %d at index %d.", l.size, position)
	}
	if position == 0 {
		l.insertFront(newItem)
	} else if position == l.size {
		l.insertBack(newItem)
	} else {
		cursor := l.front
		for i := 0; i < position; i++ {
			cursor = cursor.next
		}
		n := &node[T]{val: newItem}
		n.prev = cursor.prev
		cursor.prev.next = n
		n.next = cursor
	 	cursor.prev = n
		l.size++
	}
	return nil
}

// String returns the string representation of the List.
func (l *List[T]) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	s := "["
	cursor := l.front
	for i := 0; i < l.size - 1; i++ {
		s += fmt.Sprintf("%v", cursor.val) + " "
		cursor = cursor.next
	}
	if l.size > 0 {
		s += fmt.Sprintf("%v", cursor.val)
	}
	s += "]"
	return s
}

// Size returns the number of items in the List.
func (l *List[T]) Size() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.size
}

// IsEmpty returns a bool indicating whether or not the List is empty.
func (l *List[T]) IsEmpty() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.size == 0
}

// Clear removes all items from the List.
func (l *List[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.front = nil
	l.back = nil
	l.size = 0
}

// Get returns an item from the specified index of the List.
// If the index is invalid (aka index < 0 || index >= List.size), an error is returned.
func (l *List[T]) Get(index int) (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index < 0 || index >= l.size {
		var zeroValue T
		return zeroValue, fmt.Errorf("Cannot access index %d in a List of size %d.", index, l.size)
	}
	cursor := l.front
	for i := 0; i < index; i++ {
		cursor = cursor.next
	}
	return cursor.val, nil
}

// Copy returns a pointer to a copy of the List.
func (l *List[T]) Copy() *List[T] {
	l.mu.Lock()
	defer l.mu.Unlock()
	newList := &List[T]{comparator: l.comparator}
	cursor := l.front
	for i := 0; i < l.size; i++ {
		newList.insertBack(cursor.val)
		cursor = cursor.next
	}
	return newList
}

// Find returns the index of the first occurence of the given item in the List.
// If the item is not found in the List, -1 is returned.
func (l *List[T]) Find(item T) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	cursor := l.front
	for i := 0; i < l.size; i++ {
		if l.comparator(cursor.val, item) == 0 {
			return i
		}
		cursor = cursor.next
	}
	return -1
}

// removeFront removes the item at the front of the List.
// If the List is empty, an error is returned.
func (l *List[T]) removeFront() (T, error) {
	if l.size == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("Cannot remove the front item from an empty List.")
	}
	value := l.front.val
	if l.size == 1 {
		l.front = nil
		l.back = nil
	} else {
		l.front.next.prev = nil
		l.front = l.front.next
	}
	l.size--
	return value, nil
}

// RemoveFront removes the item at the front of the List.
// If the List is empty, an error is returned.
func (l *List[T]) RemoveFront() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.removeFront()
}

// removeBack removes the item at the back of the List.
// If the List is empty, an error is returned.
func (l *List[T]) removeBack() (T, error) {
	if l.size == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("Cannot remove the back item from an empty List.")
	}
	value := l.back.val
	if l.size == 1 {
		l.front = nil
		l.back = nil
	} else {
		l.back.prev.next = nil
		l.back = l.back.prev
	}
	l.size--
	return value, nil
}

// RemoveBack removes the item at the back of the List.
// If the List is empty, an error is returned.
func (l *List[T]) RemoveBack() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.removeBack()
}

// RemovePosition removes the item at given index from the List.
// If the index is invalid (aka index < 0 || index >= List.size), an error is returned.
func (l *List[T]) RemovePosition(index int) (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index < 0 || index >= l.size {
		var zeroValue T
		return zeroValue, fmt.Errorf("Cannot remove item at index %d in a List of size %d.", index, l.size)
	}
	var value T
	if index == 0 {
		return l.removeFront()
	} else if index == l.size - 1 {
		return l.removeBack()
	} else {
		cursor := l.front
		for i := 0; i < index; i++ {
			cursor = cursor.next
		}
		value = cursor.val
		cursor.prev.next = cursor.next
		cursor.next.prev = cursor.prev
		l.size--
	}
	return value, nil
}

// ToSlice returns the List as a slice.
func (l *List[T]) ToSlice() []T {
	l.mu.Lock()
	defer l.mu.Unlock()
	cursor := l.front
	s := make([]T, l.size)
	for i := 0; i < l.size; i++ {
		s[i] = cursor.val
		cursor = cursor.next
	}
	return s
}

// Reverse reverses the order of the items in the List.
func (l *List[T]) Reverse() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.size == 0 || l.size == 1 {
		return
	} else {
		cursor := l.front
		for i := 0; i < l.size; i++ {
			cursorNext := cursor.next
			cursor.next = cursor.prev
			cursor.prev = cursorNext
			cursor = cursorNext
		}
		tempFront := l.front
		l.front = l.back
		l.back = tempFront
	}
}
