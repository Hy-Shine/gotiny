package rand

import "testing"

func TestRandIntRange(t *testing.T) {
	// Test case: Testing the function with valid minimum and maximum values.
	min := 1
	max := 10
	result := RandIntRange(min, max)
	if result < min || result > max {
		t.Errorf("Expected a random number between %d and %d, but got %d", min, max, result)
	}

	// Test case: Testing the function with minimum value greater than maximum value.
	min = 10
	max = 5
	result = RandIntRange(min, max)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}

	// Test case: Testing the function with minimum value equal to maximum value.
	min = 5
	max = 5
	result = RandIntRange(min, max)
	if result != min {
		t.Errorf("Expected %d, but got %d", min, result)
	}
}

func TestRandInt64Range(t *testing.T) {
	// Test case: Testing the function with valid minimum and maximum values.
	min := int64(1)
	max := int64(10)
	result := RandInt64Range(min, max)
	if result < min || result > max {
		t.Errorf("Expected a random number between %d and %d, but got %d", min, max, result)
	}

	// Test case: Testing the function with minimum value greater than maximum value.
	min = int64(10)
	max = int64(5)
	result = RandInt64Range(min, max)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}

	// Test case: Testing the function with minimum value equal to maximum value.
	min = int64(5)
	max = int64(5)
	result = RandInt64Range(min, max)
	if result != min {
		t.Errorf("Expected %d, but got %d", min, result)
	}
}
