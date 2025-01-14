package container

import (
	"errors"
	"sync"
)

type RingBuffer[T any] struct {
	mu     sync.RWMutex
	buffer []T
	size   int
	count  int
	head   int
	tail   int
}

// NewRingBuffer returns a new Ring[T] with the given size.
//
// The capacity of the returned Ring is equal to the given size,
// and size should be greater than 0.
// The underlying buffer is initialized with the given size.
// The returned Ring is empty.
func NewRingBuffer[T any](size int) (*RingBuffer[T], error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}

	return &RingBuffer[T]{
		size:   size,
		buffer: make([]T, size),
	}, nil
}

// Capacity returns the capacity of the Ring
func (r *RingBuffer[T]) Capacity() int {
	return r.size
}

// Len returns the number of elements in the Ring
func (r *RingBuffer[T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.count
}

// Read reads and returns the element at the head of the Ring.
func (r *RingBuffer[T]) Read() (ele T, ok bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.count == 0 {
		var zero T
		// Return zero value and false for empty
		return zero, false
	}

	value := r.buffer[r.head]
	r.head = (r.head + 1) % r.size
	r.count--
	return value, true
}

// ReadBatch reads and returns the elements at the head of the Ring.
func (r *RingBuffer[T]) ReadBatch(count int) []T {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.count == 0 || count <= 0 {
		return nil
	}

	if count > r.count {
		count = r.count
	}

	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = r.buffer[r.head]
		r.head = (r.head + 1) % r.size
		r.count--
	}
	return result
}

// Write adds an element to the tail of the Ring.
// If the Ring is full, the element is not added and false is returned.
func (r *RingBuffer[T]) Write(ele T) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.write(false, ele)
}

// MustWrite adds an element to the tail of the Ring.
// Even if the Ring is full, the element is added.
func (r *RingBuffer[T]) MustWrite(ele T) {
	r.mu.Lock()
	r.write(true, ele)
	r.mu.Unlock()
}

// write adds an element to the tail of the Ring.
// If the Ring is full, the element is not added and false is returned.
// Otherwise, true is returned.
func (r *RingBuffer[T]) write(must bool, ele T) bool {
	if r.count == r.size && !must {
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
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.count == r.size
}

// Reset clears the Ring
func (r *RingBuffer[T]) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.head = 0
	r.tail = 0
	r.count = 0
}
