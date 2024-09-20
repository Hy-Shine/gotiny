package queue

type sliceQ[K any] struct {
	l []K
}

func NewSliceQueue[K any](cap uint) *sliceQ[K] {
	return &sliceQ[K]{l: make([]K, 0, cap)}
}

func (sq *sliceQ[K]) IsEmpty() bool {
	return len(sq.l) == 0
}

func (sq *sliceQ[K]) Push(val K) {
	sq.l = append(sq.l, val)
}

func (sq *sliceQ[K]) Pop() (K, bool) {
	var v K
	if sq.IsEmpty() {
		return v, false
	}
	v = sq.l[0]
	sq.l = sq.l[1:]
	return v, true
}

func (sq *sliceQ[K]) Peek() (K, bool) {
	var v K
	if sq.IsEmpty() {
		return v, false
	}
	v = sq.l[0]
	return v, true
}
