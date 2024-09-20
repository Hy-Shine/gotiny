package queue

import (
	"testing"
)

func TestSliceQ_op(t *testing.T) {
	q := NewSliceQueue[int](5)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	val, _ := q.Pop()
	if val != 1 {
		t.Errorf("Expected 1, but got %v", val)
	}

	q.Pop()

	val, _ = q.Peek()
	if val != 3 {
		t.Errorf("Expected 3, but got %v", val)
	}

	q.Pop()
	q.Pop()
	_, ok := q.Pop()
	if ok {
		t.Errorf("Expected false, but got %v", ok)
	}
}

func BenchmarkSliceQ_PushPop(b *testing.B) {
	q := NewSliceQueue[int](100)
	for i := 0; i < b.N; i++ {
		q.Push(i)
		if i%100 != 20 {
			continue
		}
		for j := 0; j < 20; j++ {
			q.Pop()
		}
	}

	for !q.IsEmpty() {
		q.Pop()
	}
}
