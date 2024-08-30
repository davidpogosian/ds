package set

import (
	"fmt"
	"strings"
	"sync"
)

type Set[T comparable] struct {
	items map[T]bool
	size int
	mu sync.Mutex
}

// Returns a pointer to a new empty Set.
func NewEmpty[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

// Returns a pointer to a new Set initialized with a slice.
func NewFromSlice[T comparable](slice []T) *Set[T] {
	s := Set[T]{items: make(map[T]bool)}
	for i := 0; i < len(slice); i++ {
		s.Add(slice[i])
	}
	return &s
}

// Adds an item to the Set.
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

// Removes an item from the Set.
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

// Returns the string representation of the Set.
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

// Returns a copy of the Set.
func (s *Set[T]) Copy() *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	copy := NewEmpty[T]()
	for key := range s.items {
		copy.Add(key)
	}
	return copy
}

// Returns a bool indicating whether or not the item is in the Set.
func (s *Set[T]) Contains(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.items[item]
	return exists
}

// Returns the number of items in the Set as an int.
func (s *Set[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size
}

// Returns a bool indicating the emptiness of the Set.
func (s *Set[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size == 0
}

// Removes all items from the Set.
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[T]bool)
	s.size = 0
}

// Returns the Set as a slice.
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

// Returns a new Set that is the union of this Set and the input Set.
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

// Returns a new Set that is the intersection of this Set and the input Set.
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

// Returns a new Set that is the difference between this Set and the input Set.
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

// Returns a bool that indicates if the this Set is a subset of the input Set.
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

// Returns a bool that indicates if the this Set is a superset of the input Set.
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

// Returns a bool that indicates if the this Set is equal to the input Set.
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
