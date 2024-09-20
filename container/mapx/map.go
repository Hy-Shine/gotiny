package mapx

func mapPreCompare[T comparable, V any](m1, m2 map[T]V) bool {
	switch {
	case m1 == nil && m2 == nil:
		return true
	case m1 == nil && m2 != nil:
		return false
	case m1 != nil && m2 == nil:
		return false
	case len(m1) != len(m2):
		return false
	}
	return true
}

// MapsEqual returns true if both maps have the same keys and corresponding
// values, and false otherwise.
//
// Maps of type map[T]V are considered equal if they have the same number of
// elements and each element in m1 is in m2 with the same value, and vice
// versa. The order of the elements is not significant.
//
// Both m1 and m2 must be of type map[T]V, or the function panics. The
// function returns a bool.
func KeysEqual[T comparable, V any](m1, m2 map[T]V) bool {
	if len(m1) == 0 || len(m2) == 0 {
		return mapPreCompare(m1, m2)
	}

	for k := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}
	return true
}

// Keys returns the keys of the map m as a slice. The keys are in an undefined
// order. The map m must be of type map[T]V, where T is comparable, or the
// function panics. The function returns a slice of type []T.
func Keys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m)>>2)
	Range(m, func(k T, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

func KeysFunc[T comparable, V any](m map[T]V, f func(k T) (T, bool)) []T {
	keys := make([]T, 0, len(m)>>2)
	for k := range m {
		if val, ok := f(k); ok {
			keys = append(keys, val)
		}
	}
	return keys
}

func Range[K comparable, V any](m map[K]V, f func(k K, v V) bool) {
	for k := range m {
		if !f(k, m[k]) {
			continue
		}
	}
}
