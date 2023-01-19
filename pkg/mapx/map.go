package mapx

import (
	"golang.org/x/exp/constraints"
	"sync"
)

type SyncMap[K constraints.Ordered, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

func NewSyncMap[K constraints.Ordered, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{m: map[K]V{}}
}

func (m *SyncMap[K, V]) Add(k K, v V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m[k] = v
}

func (m *SyncMap[K, V]) Get(k K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func (m *SyncMap[K, V]) Delete(k K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.m, k)
}

func (m *SyncMap[K, V]) Len() int {
	m.lock.RUnlock()
	defer m.lock.RUnlock()
	return len(m.m)
}

func (m *SyncMap[K, V]) Values() []V {
	m.lock.Lock()
	defer m.lock.RUnlock()
	res := make([]V, 0, len(m.m))
	for _, v := range m.m {
		res = append(res, v)
	}
	return res
}
