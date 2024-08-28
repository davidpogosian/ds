package list

import (
	"fmt"
	"sync"
)

type node[T comparable] struct {
	val T
	next *node[T]
	prev *node[T]
}

type List[T comparable] struct {
	front *node[T]
	back *node[T]
	size int
	mu sync.Mutex
}

// Returns a new empty List.
func NewEmpty[T comparable]() *List[T] {
	return &List[T]{}
}

// Returns a new List initialized with a slice.
func NewFromSlice[T comparable](slice []T) *List[T] {
	l := List[T]{}
	for _, item := range slice {
		l.InsertBack(item)
	}
	return &l
}

// Inserts new item at the front of the List.
func (l *List[T]) InsertFront(newItem T) {
	l.mu.Lock()
	defer l.mu.Unlock()
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

// Inserts new item at the back of the List.
func (l *List[T]) InsertBack(newItem T) {
	l.mu.Lock()
	defer l.mu.Unlock()
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

func (l *List[T]) InsertPosition(newItem T, position int) error {
	return fmt.Errorf("todo")
}
