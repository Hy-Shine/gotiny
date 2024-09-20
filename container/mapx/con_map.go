package mapx

import "sync"

type conMap[K comparable, V any] struct {
	lock sync.RWMutex
	m    map[K]V
}

func NewConMap[K comparable, V any](cap int) *conMap[K, V] {
	if cap < 0 {
		cap = 0
	}

	return &conMap[K, V]{
		lock: sync.RWMutex{},
		m:    make(map[K]V, cap),
	}
}

func (m *conMap[K, V]) Get(k K) (any, bool) {
	m.lock.RLock()
	v, ok := m.m[k]
	m.lock.RUnlock()
	return v, ok
}

func (m *conMap[K, V]) Set(k K, v V) {
	m.lock.Lock()
	m.m[k] = v
	m.lock.Unlock()
}

func (m *conMap[K, V]) Delete(k K) bool {
	m.lock.Lock()
	_, ok := m.m[k]
	if ok {
		delete(m.m, k)
	}
	m.lock.Unlock()
	return ok
}

func (m *conMap[K, V]) Len() int {
	m.lock.RLock()
	length := len(m.m)
	m.lock.RUnlock()
	return length
}

func (m *conMap[K, V]) Clear() {
	m.lock.Lock()
	m.m = make(map[K]V)
	m.lock.Unlock()
}

func (m *conMap[K, V]) Keys() []K {
	m.lock.RLock()
	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	m.lock.RUnlock()
	return keys
}

func (m *conMap[K, V]) Range(f func(k K, v V) bool) {
	m.lock.RLock()
	for k := range m.m {
		if !f(k, m.m[k]) {
			continue
		}
	}
	m.lock.RUnlock()
}
