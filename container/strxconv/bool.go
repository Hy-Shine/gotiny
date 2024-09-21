package strxconv

// NotEmpty returns true if v is not empty, otherwise returns false.
func NotEmpty[T comparable](v T) bool {
	var defaultValue T
	return v != defaultValue
}

// IsEmpty returns true if v is empty, otherwise returns false.
func IsEmpty[T comparable](v T) bool {
	return !NotEmpty(v)
}
