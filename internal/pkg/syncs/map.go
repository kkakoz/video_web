package syncs

import "sync"

type Map[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{m: map[K]V{}}
}

func (item *Map[K, V]) Get(key K) V {
	item.lock.RLock()
	defer item.lock.RUnlock()
	return item.m[key]
}

func (item *Map[K, V]) Set(key K, value V) {
	item.lock.Lock()
	defer item.lock.Unlock()
	item.m[key] = value
}

func (item *Map[K, V]) SetFirst(key K, value V) bool {
	item.lock.Lock()
	defer item.lock.Unlock()
	_, ok := item.m[key]
	if ok {
		return false
	}
	item.m[key] = value
	return true
}

func (item *Map[K, V]) Exist(key K) bool {
	item.lock.RLock()
	defer item.lock.RUnlock()
	_, ok := item.m[key]
	return ok
}

func (item *Map[K, V]) Del(key K) {
	item.lock.RLock()
	defer item.lock.RUnlock()
	delete(item.m, key)
}

// 加锁操作
func (item *Map[K, V]) Do(f func(map[K]V)) {
	item.lock.Lock()
	defer item.lock.Unlock()
	f(item.m)
}

func (item *Map[K, V]) Foreach(f func(key K, value V)) {
	item.lock.RLock()
	defer item.lock.RUnlock()
	for k, v := range item.m {
		f(k, v)
	}
}

func (item *Map[K, V]) Len() int {
	item.lock.RLock()
	defer item.lock.RUnlock()
	return len(item.m)
}
