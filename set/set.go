// Package set provides a thread-safe, generic set implementation.
package set

import (
	"fmt"
	"strings"
	"sync"
)

// Set struct represents a set.
// Important: Set can only be used with types that have the comparable constraint.
// Set stores items in a field of type map[T comparable]bool.
// Set also has a field to keep track of its size, as well as a mutex for thread-safety.
type Set[T comparable] struct {
	items map[T]bool
	size int
	mu sync.Mutex
}

// NewEmpty returns a pointer to a new empty Set.
func NewEmpty[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

// NewFromSlice returns a pointer to a new Set initialized with a slice.
func NewFromSlice[T comparable](slice []T) *Set[T] {
	s := Set[T]{items: make(map[T]bool)}
	for i := 0; i < len(slice); i++ {
		s.Add(slice[i])
	}
	return &s
}

// Add adds an item to the Set.
// If the item is already in the Set, nothing happens.
func (s *Set[T]) Add(newItem T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.items[newItem]
	if !exists {
		s.items[newItem] = true
		s.size++
	}
}

// Remove removes an item from the Set.
// If the item is not in the Set, nothing happens.
func (s *Set[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.items[item]
	if exists {
		delete(s.items, item)
		s.size--
	}
}

// String returns the string representation of the Set.
func (s *Set[T]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	str := "["
	for key := range s.items {
		str += fmt.Sprintf("%v ", key)
	}
	if strings.HasSuffix(str, " ") {
		str = strings.TrimSuffix(str, " ")
	}
	return str + "]"
}

// Copy returns a pointer to a copy of the Set.
func (s *Set[T]) Copy() *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	copy := NewEmpty[T]()
	for key := range s.items {
		copy.Add(key)
	}
	return copy
}

// Contains returns a bool indicating whether or not the item is in the Set.
func (s *Set[T]) Contains(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.items[item]
	return exists
}

// Size returns the number of items in the Set as an int.
func (s *Set[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size
}

// IsEmpty returns a bool indicating the emptiness of the Set.
func (s *Set[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size == 0
}

// Clear removes all items from the Set.
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[T]bool)
	s.size = 0
}

// ToSlice returns the Set as a slice.
func (s *Set[T]) ToSlice() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	slice := make([]T, s.size)
	i := 0
	for key := range s.items {
		slice[i] = key
		i++
	}
	return slice
}

// Union returns a pointer to a new Set that is the union of this Set
// and the Set provided as an argument.
func (s1 *Set[T]) Union(s2 *Set[T]) *Set[T] {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.Lock()
	defer s2.mu.Unlock()
	union := NewEmpty[T]()
	for key := range s1.items {
		union.Add(key)
	}
	for key := range s2.items {
		union.Add(key)
	}
	return union
}

// Intersection returns a pointer to a new Set that is the intersection of
// this Set and the Set provided as an argument.
func (s1 *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.Lock()
	defer s2.mu.Unlock()
	intersection := NewEmpty[T]()
	for key := range s1.items {
		if _, exists := s2.items[key]; exists {
			intersection.Add(key)
		}
	}
	return intersection
}

// Difference returns a pointer to a new Set that is
// the difference between this Set and the Set provided as an argument.
func (s1 *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.Lock()
	defer s2.mu.Unlock()
	difference := NewEmpty[T]()
	for key := range s1.items {
		if _, exists := s2.items[key]; !exists {
			difference.Add(key)
		}
	}
	return difference
}

// IsSubset returns a bool that indicates if this Set is a
// subset of the Set provided as an argument.
func (s1 *Set[T]) IsSubset(s2 *Set[T]) bool {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.Lock()
	defer s2.mu.Unlock()
	for key := range s1.items {
		if _, exists := s2.items[key]; !exists {
			return false
		}
	}
	return true
}

// IsSuperset returns a bool that indicates if this Set is a
// superset of the Set provided as an argument.
func (s1 *Set[T]) IsSuperset(s2 *Set[T]) bool {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.Lock()
	defer s2.mu.Unlock()
	for key := range s2.items {
		if _, exists := s1.items[key]; !exists {
			return false
		}
	}
	return true
}

// Equals returns a bool that indicates if this Set is
// equal to the Set provided as an argument.
func (s1 *Set[T]) Equals(s2 *Set[T]) bool {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.Lock()
	defer s2.mu.Unlock()
	if s1.size != s2.size {
		return false
	}
	for key := range s1.items {
		if _, exists := s2.items[key]; !exists {
			return false
		}
	}
	return true
}
