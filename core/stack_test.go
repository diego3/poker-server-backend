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

func TestStack_TableDriven(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			op    string // "push", "pop", "clear"
			value int    // value for push operation
		}
		expectedSize   int
		expectedEmpty  bool
		expectedPopVal int // expected value when popping (only relevant if last op is pop)
	}{
		{
			name: "Single push",
			operations: []struct {
				op    string
				value int
			}{
				{op: "push", value: 42},
			},
			expectedSize:  1,
			expectedEmpty: false,
		},
		{
			name: "Multiple pushes",
			operations: []struct {
				op    string
				value int
			}{
				{op: "push", value: 10},
				{op: "push", value: 20},
				{op: "push", value: 30},
			},
			expectedSize:  3,
			expectedEmpty: false,
		},
		{
			name: "Push and pop",
			operations: []struct {
				op    string
				value int
			}{
				{op: "push", value: 100},
				{op: "pop", value: 0},
			},
			expectedSize:   0,
			expectedEmpty:  true,
			expectedPopVal: 100,
		},
		{
			name: "LIFO order verification",
			operations: []struct {
				op    string
				value int
			}{
				{op: "push", value: 1},
				{op: "push", value: 2},
				{op: "push", value: 3},
				{op: "pop", value: 0}, // should return 3
				{op: "pop", value: 0}, // should return 2
			},
			expectedSize:   1,
			expectedEmpty:  false,
			expectedPopVal: 2, // last pop should return 2
		},
		{
			name: "Clear operation",
			operations: []struct {
				op    string
				value int
			}{
				{op: "push", value: 5},
				{op: "push", value: 15},
				{op: "push", value: 25},
				{op: "clear", value: 0},
			},
			expectedSize:  0,
			expectedEmpty: true,
		},
		{
			name: "Pop from empty stack",
			operations: []struct {
				op    string
				value int
			}{
				{op: "pop", value: 0},
			},
			expectedSize:   0,
			expectedEmpty:  true,
			expectedPopVal: -1, // empty stack returns -1
		},
		{
			name: "Complex sequence",
			operations: []struct {
				op    string
				value int
			}{
				{op: "push", value: 1},
				{op: "push", value: 2},
				{op: "pop", value: 0}, // pop 2
				{op: "push", value: 3},
				{op: "push", value: 4},
				{op: "pop", value: 0}, // pop 4
				{op: "pop", value: 0}, // pop 3
			},
			expectedSize:   1,
			expectedEmpty:  false,
			expectedPopVal: 3, // last pop should return 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{}
			var lastPopVal int

			// Execute all operations
			for _, op := range tt.operations {
				switch op.op {
				case "push":
					s.Push(op.value)
				case "pop":
					val, err := s.Pop()
					if err != nil {
						t.Errorf("Unexpected error on pop: %v", err)
					}
					lastPopVal = val
				case "clear":
					s.Clear()
				default:
					t.Fatalf("Unknown operation: %s", op.op)
				}
			}

			// Verify final state
			if s.Size() != tt.expectedSize {
				t.Errorf("Expected size %d, got %d", tt.expectedSize, s.Size())
			}

			if s.Empty() != tt.expectedEmpty {
				t.Errorf("Expected empty=%v, got %v", tt.expectedEmpty, s.Empty())
			}

			// If last operation was pop, verify the value
			if len(tt.operations) > 0 && tt.operations[len(tt.operations)-1].op == "pop" {
				if lastPopVal != tt.expectedPopVal {
					t.Errorf("Expected pop value %d, got %d", tt.expectedPopVal, lastPopVal)
				}
			}
		})
	}
}
