package main

import (
	"sync"
	"testing"
)

type Small struct {
	a int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Small)
	},
}

func inc(s *Small) {
	s.a++
}

func BenchmarkWithoutPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		s = &Small{a: 1}
		inc(s)
	}
}
func BenchmarkWithPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		s = pool.Get().(*Small)
		s.a = 1
		inc(s)
		pool.Put(s)
	}
}
