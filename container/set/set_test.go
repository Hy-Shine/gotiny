package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s == nil {
		t.Error("NewSet() returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	s = NewSet[int](10)
	if s == nil {
		t.Error("NewSet(10) returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	s.Add(4)
	if s.Len() != 4 {
		t.Errorf("Expected length 4, got %d", s.Len())
	}

	s.Add(1) // Adding duplicate
	if s.Len() != 4 {
		t.Errorf("Expected length 4, got %d", s.Len())
	}
}

func TestDelete(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	s.Delete(2)
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}
	if s.IsExists(2) {
		t.Error("Expected element 2 to be deleted")
	}

	s.Delete(4) // Deleting non-existent element
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}
}

func TestIsExists(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	if !s.IsExists(1) {
		t.Error("Expected element 1 to exist")
	}
	if s.IsExists(4) {
		t.Error("Expected element 4 to not exist")
	}
}

func TestLen(t *testing.T) {
	s := NewSet[int]()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	s.Add(1, 2, 3)
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	s.Delete(1)
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}
}

func TestElems(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	elems := s.Elems()
	if len(elems) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(elems))
	}

	// Check if all elements are present
	expected := map[int]bool{1: true, 2: true, 3: true}
	for _, elem := range elems {
		if !expected[elem] {
			t.Errorf("Unexpected element %d", elem)
		}
	}
}

func TestRange(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)

	count := 0
	s.Range(func(k int) bool {
		count++
		return true
	})

	if count != 3 {
		t.Errorf("Expected 3 iterations, got %d", count)
	}

	// Test early exit
	count = 0
	s.Range(func(k int) bool {
		count++
		return count < 2
	})

	if count != 2 {
		t.Errorf("Expected 2 iterations, got %d", count)
	}
}

func TestClear(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}
}

func TestSet(t *testing.T) {
	set := NewSet[int](10)

	t.Run("add_elements", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			set.Add(i)
		}
	})

	t.Run("set_len", func(t *testing.T) {
		if set.Len() != 10 {
			t.Errorf("Expected %v but got %v", 10, set.Len())
		}
	})

	t.Run("set_delete", func(t *testing.T) {
		set.Delete(5)
		if set.Len() != 9 {
			t.Errorf("Expected %v but got %v", 9, set.Len())
		}
	})

	t.Run("set_exists", func(t *testing.T) {
		exists := set.IsExists(1)
		if !exists {
			t.Errorf("Expected %v but got %v", true, exists)
		}
		exists = set.IsExists(5)
		if exists {
			t.Errorf("Expected %v but got %v", false, exists)
		}
	})

	t.Run("", func(t *testing.T) {
		list := set.Elems()
		sort.Ints(list)
		Expected := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
		if !reflect.DeepEqual(Expected, list) {
			t.Errorf("Expected %v but got %v", Expected, list)
		}
	})

	t.Run("set_range", func(t *testing.T) {
		var list []int
		set.Range(func(k int) bool {
			if k%2 == 0 {
				list = append(list, k)
				return true
			}
			return false
		})
		sort.Ints(list)
		Expected := []int{0, 2, 4, 6, 8}
		if !reflect.DeepEqual(Expected, list) {
			t.Errorf("Expected %v but got %v", Expected, list)
		}
	})

	t.Run("set_clear", func(t *testing.T) {
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected %v but got %v", 0, set.Len())
		}
	})
}
