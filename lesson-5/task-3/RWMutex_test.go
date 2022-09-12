package main

import (
	"sync"
	"testing"
)

//  3. Протестируйте производительность операций чтения и записи на множестве
//действительных чисел, безопасность которого обеспечивается sync.Mutex и
//sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50%
//запись, 50% чтение; 90% запись, 10% чтение

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}
func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}
func (s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkSet10Add90Has(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetAdd50Add50Set(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetAdd90Add10Set(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}
