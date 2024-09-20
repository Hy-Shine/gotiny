package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	expected := 4
	result := BinarySearch(arr, target)
	if result != expected {
		t.Errorf("Test case 1 failed. Expected: %d, got: %d", expected, result)
	}

	target = 10
	expected = -1
	result = BinarySearch(arr, target)
	if result != expected {
		t.Errorf("Test case 2 failed. Expected: %d, got: %d", expected, result)
	}

	arr = []int{}
	target = 5
	expected = -1
	result = BinarySearch(arr, target)
	if result != expected {
		t.Errorf("Test case 3 failed. Expected: %d, got: %d", expected, result)
	}

	arr = []int{1}
	target = 1
	expected = 0
	result = BinarySearch(arr, target)
	if result != expected {
		t.Errorf("Test case 4 failed. Expected: %d, got: %d", expected, result)
	}

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 9
	expected = 8
	result = BinarySearch(arr, target)
	if result != expected {
		t.Errorf("Test case 5 failed. Expected: %d, got: %d", expected, result)
	}

	arrBytes := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	expected = 2
	result = BinarySearch(arrBytes, 'c')
	if result != expected {
		t.Errorf("Test case 6 failed. Expected: %d, got: %d", expected, result)
	}
}
