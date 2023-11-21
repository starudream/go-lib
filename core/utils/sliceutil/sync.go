package sliceutil

import (
	"sync"
)

type SyncSlice[V any] struct {
	data []V
	mu   sync.RWMutex
}

func (s *SyncSlice[V]) Append(v ...V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, v...)
}

func (s *SyncSlice[V]) Index(i int) V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[i]
}

func (s *SyncSlice[V]) Delete(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data[:i], s.data[i+1:]...)
}

func (s *SyncSlice[V]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *SyncSlice[V]) Data() []V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data
}

func (s *SyncSlice[V]) Range(f func(i int, v V) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, v := range s.data {
		if !f(i, v) {
			break
		}
	}
}
