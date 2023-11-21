package maputil

import (
	"sync"
)

type SyncMap[K comparable, V any] sync.Map

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.m().Store(k, v)
}

func (m *SyncMap[K, V]) Load(k K) (V, bool) {
	t, ok := m.m().Load(k)
	if !ok {
		return m.empty(), false
	}
	return t.(V), true
}

func (m *SyncMap[K, V]) Exists(k K) bool {
	_, ok := m.m().Load(k)
	return ok
}

func (m *SyncMap[K, V]) Delete(k K) {
	m.m().Delete(k)
}

func (m *SyncMap[K, V]) Len() (l int) {
	m.m().Range(func(k, v any) bool { l++; return true })
	return
}

func (m *SyncMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	t, loaded := m.m().LoadOrStore(k, v)
	return t.(V), loaded
}

func (m *SyncMap[K, V]) LoadAndDelete(k K) (V, bool) {
	t, loaded := m.m().LoadAndDelete(k)
	if !loaded {
		return m.empty(), false
	}
	return t.(V), loaded
}

func (m *SyncMap[K, V]) CompareAndSwap(k K, old, new V) bool {
	return m.m().CompareAndSwap(k, old, new)
}

func (m *SyncMap[K, V]) CompareAndDelete(k K, old V) bool {
	return m.m().CompareAndDelete(k, old)
}

func (m *SyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.m().Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (m *SyncMap[K, V]) m() *sync.Map {
	return (*sync.Map)(m)
}

func (m *SyncMap[K, V]) empty() V {
	var v V
	return v
}
