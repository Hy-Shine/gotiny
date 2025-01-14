package set

// gset is a generic set implementation using a map with empty struct values.
type gset[K comparable] struct {
	m map[K]struct{}
}

// setCap determines the initial capacity of the set based on the provided arguments.
// If no capacity is provided, it defaults to 0.
func setCap(cap ...int) int {
	var c int
	if len(cap) > 0 {
		c = cap[0]
	}
	if c < 0 {
		c = 0
	}
	return c
}

// NewSet creates and returns a new set with the specified initial capacity.
// If no capacity is provided, the set is initialized with a default capacity of 0.
func NewSet[K comparable](cap ...int) *gset[K] {
	initCap := setCap(cap...)
	return &gset[K]{
		m: make(map[K]struct{}, initCap),
	}
}

// Add adds one or more elements to the set.
// If an element already exists in the set, it will not be added again.
func (s *gset[K]) Add(elems ...K) {
	for i := range elems {
		s.m[elems[i]] = struct{}{}
	}
}

// Delete removes the specified element from the set.
// If the element does not exist in the set, the operation is a no-op.
func (s *gset[K]) Delete(v K) {
	delete(s.m, v)
}

// IsExists checks if the specified element exists in the set.
// It returns true if the element is found, otherwise false.
func (s *gset[K]) IsExists(v K) bool {
	_, ok := s.m[v]
	return ok
}

// Len returns the number of elements in the set.
func (s *gset[K]) Len() int {
	return len(s.m)
}

// Elems returns a slice containing all the elements in the set.
// The order of elements in the slice is not guaranteed.
func (s *gset[K]) Elems() []K {
	keys := make([]K, 0, s.Len())
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

// Range iterates over the elements of the set and calls the provided function f for each element.
// If f returns false, the iteration stops.
// The order of iteration is not guaranteed, and elements may be visited multiple times.
// The function f may be called concurrently.
func (s *gset[K]) Range(f func(k K) bool) {
	for k := range s.m {
		if !f(k) {
			continue
		}
	}
}

// Clear removes all elements from the set, effectively resetting it to an empty state.
func (s *gset[K]) Clear() {
	s.m = make(map[K]struct{})
}
