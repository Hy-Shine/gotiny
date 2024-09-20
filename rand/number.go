package rand

import "math/rand"

// RandIntRange generates a random integer between min and max (inclusive).
func RandIntRange(min int, max int) int {
	if min > max {
		return 0
	}
	return min + rand.Intn(max-min+1)
}

// RandInt64Range generates a random integer between min and max (inclusive).
func RandInt64Range(min, max int64) int64 {
	if min > max {
		return 0
	}
	return min + rand.Int63n(max-min+1)
}
