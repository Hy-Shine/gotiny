package cal

import "testing"

func TestMax(t *testing.T) {
	// Test when v1 is greater than v2
	result := Max(5, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %v", result)
	}
	if Max(0, 0) != 0 {
		t.Errorf("Expected 0, but got %v", result)
	}

	// max for string
	s := Max("a", "b")
	if s != "b" {
		t.Errorf("Expected b, but got %v", s)
	}

	// max for uint8
	if Max(uint8(5), uint8(3)) != uint8(5) {
		t.Errorf("Expected 5, but got %v", result)
	}
}

func TestMin(t *testing.T) {
	// Test when v1 is greater than v2
	result := Min(5, 3)
	if result != 3 {
		t.Errorf("Expected 3, but got %v", result)
	}
	if Min(0, 0) != 0 {
		t.Errorf("Expected 0, but got %v", result)
	}

	// min for string
	s := Min("a", "b")
	if s != "a" {
		t.Errorf("Expected a, but got %v", s)
	}

	// min for uint8
	if Min(uint8(5), uint8(3)) != uint8(3) {
		t.Errorf("Expected 3, but got %v", result)
	}
}
