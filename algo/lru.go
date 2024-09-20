package algo

import (
	"container/list"
)

type Store[K comparable, V any] struct {
	key K
	val V
}

func (s *Store[K, V]) Key() K {
	return s.key
}

func (s *Store[K, V]) Val() V {
	return s.val
}

type LRUCache[K comparable, V any] struct {
	cap  int
	m    map[K]*list.Element
	list *list.List
}

func NewLRU[K comparable, V any](cap int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		cap:  cap,
		m:    make(map[K]*list.Element, cap),
		list: list.New(),
	}
}

func (lru *LRUCache[K, V]) Get(key K) (*Store[K, V], bool) {
	node, ok := lru.m[key]
	if !ok {
		return nil, false
	}

	// move node to front
	lru.list.MoveToFront(node)
	return node.Value.(*Store[K, V]), true
}

func (lru *LRUCache[K, V]) Put(key K, value V) {
	newVal := &Store[K, V]{key: key, val: value}
	if val, ok := lru.m[key]; ok {
		val.Value = newVal
		lru.list.MoveToFront(val)
		return
	}

	if lru.cap <= len(lru.m) {
		st := lru.list.Remove(lru.list.Back()).(*Store[K, V])
		delete(lru.m, st.Key())
	}
	// push to front
	node := lru.list.PushFront(newVal)
	lru.m[key] = node
}
