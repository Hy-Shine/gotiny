package stack

type stackSlice[K any] struct {
	l []K
}

func NewSliceStack[K any](cap uint) *stackSlice[K] {
	return &stackSlice[K]{l: make([]K, 0, cap)}
}

func (ss *stackSlice[K]) Push(val K) {
	ss.l = append(ss.l, val)
}

func (ss *stackSlice[K]) Pop() (K, bool) {
	var v K
	if ss.IsEmpty() {
		return v, false
	}
	v = ss.l[len(ss.l)-1]
	ss.l = ss.l[:len(ss.l)-1]
	return v, true
}

func (ss *stackSlice[K]) Peek() (K, bool) {
	var v K
	if ss.IsEmpty() {
		return v, false
	}
	v = ss.l[len(ss.l)-1]
	return v, true
}

func (ss *stackSlice[K]) IsEmpty() bool {
	return len(ss.l) == 0
}
