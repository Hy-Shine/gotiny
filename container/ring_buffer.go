package container

import (
	"sync"
)

type RingBuffer[T any] struct {
	buffer []T
	size   int
	count  int
	head   int
	tail   int
	lock   sync.RWMutex
}

// NewRingBuffer returns a new Ring[T] with the given size.
//
// The capacity of the returned Ring is equal to the given size,
// and size should be greater than 0.
// The underlying buffer is initialized with the given size.
// The returned Ring is empty.
func NewRingBuffer[T any](size uint) *RingBuffer[T] {
	return &RingBuffer[T]{
		size:   int(size),
		buffer: make([]T, size),
	}
}

// Len returns the number of elements in the Ring
func (r *RingBuffer[T]) Len() int {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.count
}

// Read removes and returns the element at the head of the Ring.
func (r *RingBuffer[T]) Read() (ele T, ok bool) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.count == 0 {
		var zero T
		return zero, false // Return zero value and false for empty
	}

	value := r.buffer[r.head]
	r.head = (r.head + 1) % r.size
	r.count--
	return value, true
}

// Write adds an element to the tail of the Ring.
// If the Ring is full, the element is not added and false is returned.
func (r *RingBuffer[T]) Write(ele T) bool {
	return r.write(false, ele)
}

// MustWrite adds an element to the tail of the Ring.
// Even if the Ring is full, the element is added.
func (r *RingBuffer[T]) MustWrite(ele T) {
	r.write(true, ele)
}

// write adds an element to the tail of the Ring.
// If the Ring is full, the element is not added and false is returned.
// Otherwise, true is returned.
func (r *RingBuffer[T]) write(must bool, ele T) bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.size == 0 || r.count == r.size && !must {
		return false
	}

	r.buffer[r.tail] = ele
	r.tail = (r.tail + 1) % r.size
	if r.count < r.size {
		r.count++
	}
	return true
}

// IsFull returns true if the Ring is full
func (r *RingBuffer[T]) IsFull() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.count == r.size
}
