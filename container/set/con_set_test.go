package set

import (
	"fmt"
	"sync"
	"testing"
)

func TestSet_Add(t *testing.T) {
	hash := New[int](0)
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		n := i
		go func() {
			hash.Add(n, n+1)
			wg.Done()
		}()
	}
	wg.Wait()
	for i := range hash.Keys() {
		fmt.Println(i)
	}
}

func BenchmarkSet_Add(b *testing.B) {
	hash := New[int](0)
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		n := i
		go func() {
			hash.Add(n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(hash.Len())
}
