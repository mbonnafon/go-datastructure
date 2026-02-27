package stack

import "errors"

// Stack is a last-in, first-out (LIFO) data structure
type Stack[T any] struct {
	items []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return zeroValue[T](), errors.New("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the top item from the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return zeroValue[T](), errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func zeroValue[T any]() T {
	var t T
	return t
}
