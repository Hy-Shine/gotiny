package stack

import (
	"testing"
)

func TestStack_int(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if val, _ := s.Peek(); val != 3 {
		t.Error("Expected 3, but got", val)
	}

	s.Pop()
	if val, _ := s.Peek(); val != 2 {
		t.Error("Expected 2, but got", val)
	}
	s.Pop()
	last, _ := s.Pop()
	if last != 1 {
		t.Error("Expected 1, but got", last)
	}

	s.Pop()
	_, b := s.Pop()
	if b {
		t.Errorf("Expected false, but got %v", b)
	}
	if !s.IsEmpty() {
		t.Error("Expected true, but got", s.IsEmpty())
	}
}

func BenchmarkStack_int(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for !s.IsEmpty() {
		s.Pop()
	}
}

func TestStack_string(t *testing.T) {
	s := New[string]()
	s.Push("a")
	s.Push("b")
	s.Push("c")

	if val, _ := s.Peek(); val != "c" {
		t.Error("Expected c, but got", val)
	}

	s.Pop()
	if val, _ := s.Peek(); val != "b" {
		t.Error("Expected b, but got", val)
	}

	s.Pop()
	s.Pop()
	if !s.IsEmpty() {
		t.Error("Expected true, but got", s.IsEmpty())
	}
}

func TestStack_any(t *testing.T) {
	s := New[any]()
	s.Push(1)
	s.Push("b")
	s.Push('c')

	if val, _ := s.Peek(); val != 'c' {
		t.Error("Expected c, but got", val)
	}

	s.Pop()
	if val, _ := s.Peek(); val != "b" {
		t.Error("Expected b, but got", val)
	}

	s.Pop()
	s.Pop()
	if !s.IsEmpty() {
		t.Error("Expected true, but got", s.IsEmpty())
	}
}
