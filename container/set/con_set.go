package set

import "sync"

type conSet[K comparable] struct {
	lock sync.RWMutex
	m    map[K]struct{}
}

func New[K comparable](cap ...int) *conSet[K] {
	initCap := setCap(cap...)
	return &conSet[K]{
		lock: sync.RWMutex{},
		m:    make(map[K]struct{}, initCap),
	}
}

func (h *conSet[K]) Add(l ...K) {
	h.lock.Lock()
	for _, v := range l {
		h.m[v] = struct{}{}
	}
	h.lock.Unlock()
}

func (h *conSet[K]) Delete(v K) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.m, v)
}

func (h *conSet[K]) Clear() {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.m = make(map[K]struct{})
}

func (h *conSet[K]) Keys() []K {
	h.lock.Lock()
	keys := make([]K, 0, h.Len())
	for k := range h.m {
		keys = append(keys, k)
	}
	h.lock.Unlock()
	return keys
}

func (h *conSet[K]) Len() int {
	h.lock.Lock()
	length := len(h.m)
	h.lock.Unlock()
	return length
}

func (h *conSet[K]) IsExists(e K) bool {
	h.lock.RLock()
	_, ok := h.m[e]
	h.lock.Unlock()
	return ok
}

func (h *conSet[K]) Range(f func(k K) bool) {
	h.lock.RLock()
	for k := range h.m {
		if !f(k) {
			continue
		}
	}
	h.lock.Unlock()
}
