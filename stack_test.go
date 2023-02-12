package tdd_stack_test

import (
	"testing"

	tdd_stack "github.com/rzldimam28/tdd-stack"
	"github.com/stretchr/testify/assert"
)

// requirements:
// * A stack is empty on construction
// * A stack has size 0 on construction
// * After n pushes to an empty stack (n > 0), the stack is non-empty and its size equals n
// * If one pushes x then pops, the value popped is x, and the size is one less than old size
// * If one pushes x then peeks, the value returned is x, but the size stays the same
// * If the size is n, then after n pops, the stack is empty and has size 0
// * Popping from an empty stack return an error: ErrNoSuchElement
// * Peeking into an empty stack return an error: ErrNoSuchElement

func TestNewStack(t *testing.T) {
	t.Run("A stack is empty on construction", func(t *testing.T) {
		s := tdd_stack.New()
		assert.True(t, s.IsEmpty())
	})
	t.Run("A stack has size 0 on construction", func(t *testing.T) {
		s := tdd_stack.New()
		assert.Equal(t, 0, s.Size())
	})
}

func TestInsert(t *testing.T) {
	t.Run("After n pushes to an empty stack (n > 0), the stack is non-empty and its size equals n", func(t *testing.T) {
		s := tdd_stack.New()
		s.Push(5)
		s.Push(3)
		s.Push(7)
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())
	})
	t.Run("If one pushes x then pops, the value popped is x, and the size is one less than old size", func(t *testing.T) {
		s := tdd_stack.New()
		s.Push(5)
		s.Push(3)
		s.Push(7)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Pop()
		assert.Equal(t, 7, val)
		assert.Equal(t, 2, s.Size())
	})
	t.Run("If one pushes x then peeks, the value returned is x, but the size stays the same", func(t *testing.T) {
		s := tdd_stack.New()
		s.Push(5)
		s.Push(3)
		s.Push(7)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Peek()
		assert.Equal(t, 7, val)
		assert.Equal(t, 3, s.Size())
	})
	t.Run("If the size is n, then after n pops, the stack is empty and has size 0", func(t *testing.T) {
		s := tdd_stack.New()
		s.Push(5)
		s.Push(3)
		s.Push(7)
		assert.Equal(t, 3, s.Size())
		s.Pop()
		s.Pop()
		s.Pop()
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())
	})
}

func TestError(t *testing.T) {
	t.Run("Popping from an empty stack return an error: ErrNoSuchElement", func(t *testing.T) {
		s := tdd_stack.New()
		_, err := s.Pop()
		if err == nil {
			t.Fail()
			t.Logf("Expected error is not nill, but got %v", err)
		}
		assert.Equal(t, tdd_stack.ErrNoSuchElement, err)
	})
	t.Run("Peeking into an empty stack return an error: ErrNoSuchElement", func(t *testing.T) {
		s := tdd_stack.New()
		_, err := s.Peek()
		if err == nil {
			t.Fail()
			t.Logf("Expected error is not nill, but got %v", err)
		}
		assert.Equal(t, tdd_stack.ErrNoSuchElement, err)
	})
}