package core

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := Stack{}

	s.Push(10)
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
	if s.Empty() {
		t.Error("Stack should not be empty after push")
	}

	s.Push(20)
	s.Push(30)
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	t.Run("Pop from non-empty stack", func(t *testing.T) {
		s := Stack{}
		s.Push(10)
		s.Push(20)
		s.Push(30)

		val, err := s.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if val != 30 {
			t.Errorf("Expected 30, got %d", val)
		}
		if s.Size() != 2 {
			t.Errorf("Expected size 2, got %d", s.Size())
		}

		val, err = s.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if val != 20 {
			t.Errorf("Expected 20, got %d", val)
		}

		val, err = s.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if val != 10 {
			t.Errorf("Expected 10, got %d", val)
		}
		if !s.Empty() {
			t.Error("Stack should be empty after popping all items")
		}
	})

	t.Run("Pop from empty stack", func(t *testing.T) {
		s := Stack{}
		val, err := s.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if val != -1 {
			t.Errorf("Expected -1 for empty stack, got %d", val)
		}
		if !s.Empty() {
			t.Error("Stack should remain empty")
		}
	})
}

func TestStack_Empty(t *testing.T) {
	s := Stack{}
	if !s.Empty() {
		t.Error("New stack should be empty")
	}

	s.Push(10)
	if s.Empty() {
		t.Error("Stack should not be empty after push")
	}

	s.Pop()
	if !s.Empty() {
		t.Error("Stack should be empty after popping all items")
	}
}

func TestStack_Size(t *testing.T) {
	s := Stack{}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	s.Push(10)
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}

	s.Push(20)
	s.Push(30)
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	s.Pop()
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
}

func TestStack_Clear(t *testing.T) {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)

	s.Clear()
	if !s.Empty() {
		t.Error("Stack should be empty after Clear()")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", s.Size())
	}

	// Test that we can still use the stack after clearing
	s.Push(40)
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
}

func TestStack_LIFO(t *testing.T) {
	// Test Last In First Out behavior
	s := Stack{}
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		s.Push(v)
	}

	// Pop should return values in reverse order
	for i := len(values) - 1; i >= 0; i-- {
		val, err := s.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if val != values[i] {
			t.Errorf("Expected %d, got %d", values[i], val)
		}
	}

	if !s.Empty() {
		t.Error("Stack should be empty after popping all items")
	}
}
