package set

import "github.com/hy-shine/gotiny/cal"

type gset[K comparable] struct {
	c int
	m map[K]struct{}
}

func setCap(cap ...int) int {
	var c int
	if len(cap) > 0 {
		c = cap[0]
	}
	c = cal.If[int](c > 0, c, 0)
	return c
}

func NewSet[K comparable](cap ...int) *gset[K] {
	initCap := setCap(cap...)
	return &gset[K]{
		c: initCap,
		m: make(map[K]struct{}, initCap),
	}
}

func (s *gset[K]) Add(vs ...K) {
	for i := range vs {
		s.m[vs[i]] = struct{}{}
	}
}

func (s *gset[K]) Delete(v K) {
	delete(s.m, v)
}

func (s *gset[K]) IsExists(v K) bool {
	_, ok := s.m[v]
	return ok
}

func (s *gset[K]) Len() int {
	return len(s.m)
}

func (s *gset[K]) Keys() []K {
	keys := make([]K, 0, s.Len())
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *gset[K]) Range(f func(k K) bool) {
	for k := range s.m {
		if !f(k) {
			continue
		}
	}
}

func (s *gset[K]) Clear() {
	s.m = make(map[K]struct{}, s.c)
}
