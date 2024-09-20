package mapx

import (
	"reflect"
	"testing"
)

func TestMapKeysEqual(t *testing.T) {
	t.Run("nil maps", func(t *testing.T) {
		if !KeysEqual[string, int](nil, nil) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("Both maps are nil or empty", func(t *testing.T) {
		var m map[string]int
		var m2 map[string]int
		if !KeysEqual(m, m2) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("One map is nil, the other is not", func(t *testing.T) {
		m1 := map[string]int{"a": 1}
		if KeysEqual(m1, nil) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("Maps have different lengths", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1}
		if KeysEqual(m1, m2) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("Maps are equal", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1, "b": 2}
		if !KeysEqual(m1, m2) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("Maps have different values for the same keys", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"a": 1, "b": 3}
		if !KeysEqual(m1, m2) {
			t.Errorf("Expected true, got false")
		}
	})
}

func TestKeys(t *testing.T) {
	// Testing for an empty map
	m1 := map[int]string{}
	expected1 := []int{}
	result1 := Keys(m1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Testing for a map with non-empty values
	m2 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	expected2 := []string{"apple", "banana", "cherry"}
	result2 := Keys(m2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}

	// Testing for a map with duplicate keys
	m3 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "two again",
	}
	expected3 := []int{1, 2, 3, 4}
	result3 := Keys(m3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, result3)
	}
}

func TestKeysFunc(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		f        func(k int) (int, bool)
		expected []int
	}{
		{
			name:     "Empty map",
			m:        map[int]string{},
			f:        func(k int) (int, bool) { return k, true },
			expected: []int{},
		},
		{
			name:     "Function returns true for all keys",
			m:        map[int]string{1: "a", 2: "b", 3: "c"},
			f:        func(k int) (int, bool) { return k, true },
			expected: []int{1, 2, 3},
		},
		{
			name:     "Function returns false for all keys",
			m:        map[int]string{1: "a", 2: "b", 3: "c"},
			f:        func(k int) (int, bool) { return k, false },
			expected: []int{},
		},
		{
			name: "Mix of keys returning true and false",
			m:    map[int]string{1: "a", 2: "b", 3: "c", 4: ""},
			f: func(k int) (int, bool) {
				if k%2 == 0 {
					return k, true
				}
				return k, false
			},
			expected: []int{2, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := KeysFunc(test.m, test.f)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
