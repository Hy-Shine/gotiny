package set

import (
	"reflect"
	"sort"
	"testing"
)

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
		list := set.Keys()
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
