package mapx

import "testing"

func TestConMapRange(t *testing.T) {
	// Test case 1: Empty map
	m1 := NewConMap[int, string](0)
	m1.Range(func(k int, v string) bool {
		t.Errorf("Function should not be called for empty map")
		return false
	})

	// Test case 2: Non-empty map
	m2 := NewConMap[string, string](5)
	m2.Set("key1", "value1")
	m2.Set("key2", "value2")
	count := 0
	m2.Range(func(k string, v string) bool {
		count++
		return true
	})
	if count != 2 {
		t.Errorf("Expected 2 iterations, got %d", count)
	}

	// Test case 3: delete
	m2.Delete("key1")

	// Test case 4: length
	if m2.Len() != 1 {
		t.Errorf("Expected 2 elements, got %d", m2.Len())
	}

	// Test case 5: clear
	m2.Clear()
	if m2.Len() != 0 {
		t.Errorf("Expected 0 elements, got %d", m2.Len())
	}
}
