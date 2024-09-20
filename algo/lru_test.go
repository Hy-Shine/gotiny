package algo

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[int, int](3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Put(5, 5)
	lru.Put(6, 6)

	t.Run("test for not exists key 1", func(t *testing.T) {
		if _, ok := lru.Get(1); ok {
			t.Errorf("Expected false, but got true")
		}
	})

	t.Run("test for not exists key 2", func(t *testing.T) {
		if _, ok := lru.Get(2); ok {
			t.Errorf("Expected false, but got true")
		}
	})

	t.Run("test for exists key 6", func(t *testing.T) {
		if _, ok := lru.Get(6); !ok {
			t.Errorf("Expected false, but got true")
		}
	})
}
