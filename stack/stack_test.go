package stack_test

import (
	"go-datastructure/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	t.Run("creates an empty stack", func(t *testing.T) {
		s := stack.NewStack[int]()
		assert.Equal(t, 0, s.Size())
		assert.True(t, s.IsEmpty())
	})
}

func TestPush(t *testing.T) {
	t.Run("pushes a single item", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(42)
		assert.Equal(t, 1, s.Size())
		assert.False(t, s.IsEmpty())
	})

	t.Run("pushes multiple items", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Size())
	})

	t.Run("pushes items with string type", func(t *testing.T) {
		s := stack.NewStack[string]()
		s.Push("hello")
		s.Push("world")
		assert.Equal(t, 2, s.Size())
		item, err := s.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "world", item)
	})
}

func TestPop(t *testing.T) {
	t.Run("pops the last pushed item (LIFO order)", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		item, err := s.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 3, item)

		item, err = s.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 2, item)

		item, err = s.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 1, item)
	})

	t.Run("decreases the size after each pop", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(10)
		s.Push(20)
		assert.Equal(t, 2, s.Size())

		s.Pop()
		assert.Equal(t, 1, s.Size())

		s.Pop()
		assert.Equal(t, 0, s.Size())
	})

	t.Run("returns an error when popping from an empty stack", func(t *testing.T) {
		s := stack.NewStack[int]()
		item, err := s.Pop()
		assert.Error(t, err)
		assert.Equal(t, "stack is empty", err.Error())
		assert.Equal(t, 0, item)
	})

	t.Run("returns an error when popping after all items have been removed", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Pop()

		item, err := s.Pop()
		assert.Error(t, err)
		assert.Equal(t, "stack is empty", err.Error())
		assert.Equal(t, 0, item)
	})

	t.Run("returns zero value for string type on empty pop", func(t *testing.T) {
		s := stack.NewStack[string]()
		item, err := s.Pop()
		assert.Error(t, err)
		assert.Equal(t, "", item)
	})
}

func TestPeek(t *testing.T) {
	t.Run("returns the top item without removing it", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Push(2)

		item, err := s.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 2, item)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("returns the same item on consecutive peeks", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(99)

		item1, _ := s.Peek()
		item2, _ := s.Peek()
		assert.Equal(t, item1, item2)
		assert.Equal(t, 1, s.Size())
	})

	t.Run("returns an error when peeking an empty stack", func(t *testing.T) {
		s := stack.NewStack[int]()
		item, err := s.Peek()
		assert.Error(t, err)
		assert.Equal(t, "stack is empty", err.Error())
		assert.Equal(t, 0, item)
	})

	t.Run("reflects the new top after a pop", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(10)
		s.Push(20)

		s.Pop()
		item, err := s.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 10, item)
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("returns true for a new stack", func(t *testing.T) {
		s := stack.NewStack[int]()
		assert.True(t, s.IsEmpty())
	})

	t.Run("returns false after pushing an item", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		assert.False(t, s.IsEmpty())
	})

	t.Run("returns true after pushing and popping all items", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Pop()
		s.Pop()
		assert.True(t, s.IsEmpty())
	})
}

func TestSize(t *testing.T) {
	t.Run("returns 0 for a new stack", func(t *testing.T) {
		s := stack.NewStack[int]()
		assert.Equal(t, 0, s.Size())
	})

	t.Run("increases after push", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		assert.Equal(t, 1, s.Size())
		s.Push(2)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("decreases after pop", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Pop()
		assert.Equal(t, 1, s.Size())
	})
}

func TestPushAndPopInterleaved(t *testing.T) {
	t.Run("handles interleaved push and pop operations", func(t *testing.T) {
		s := stack.NewStack[int]()

		s.Push(1)
		s.Push(2)
		item, _ := s.Pop()
		assert.Equal(t, 2, item)

		s.Push(3)
		item, _ = s.Pop()
		assert.Equal(t, 3, item)

		item, _ = s.Pop()
		assert.Equal(t, 1, item)

		assert.True(t, s.IsEmpty())
	})
}
