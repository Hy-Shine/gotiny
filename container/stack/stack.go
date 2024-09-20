package stack

import "container/list"

type stack[V any] struct {
	l *list.List
}

func New[V any]() *stack[V] {
	return &stack[V]{l: list.New()}
}

func (st *stack[V]) Push(val V) {
	st.l.PushBack(val)
}

func (st *stack[V]) Pop() (V, bool) {
	if st.IsEmpty() {
		var v V
		return v, false
	}
	return st.l.Remove(st.l.Back()).(V), true
}

func (st *stack[V]) Peek() (V, bool) {
	if st.IsEmpty() {
		var v V
		return v, false
	}
	return st.l.Back().Value.(V), true
}

func (st *stack[V]) Len() int {
	return st.l.Len()
}

func (st *stack[V]) IsEmpty() bool {
	return st.Len() == 0
}
