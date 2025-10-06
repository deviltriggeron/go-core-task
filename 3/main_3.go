package mymap

import (
	"fmt"
	"maps"
	"sync"
)

type MyMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

func NewMyMap[K comparable, V any]() *MyMap[K, V] {
	return &MyMap[K, V]{
		m:  make(map[K]V),
		mu: sync.RWMutex{},
	}
}

func (m *MyMap[K, V]) Add(key K, value V) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.m[key] = value
}

func (m *MyMap[K, V]) Remove(key K) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	delete(m.m, key)
}

func (m *MyMap[K, V]) Copy() *MyMap[K, V] {
	newMap := NewMyMap[K, V]()

	maps.Copy(newMap.m, m.m)

	return newMap
}

func (m *MyMap[K, V]) Exists(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.m[key]
	return ok
}

func (m *MyMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	exists, ok := m.m[key]
	return exists, ok
}

func main() {
	m := NewMyMap[string, float32]()
	m.Add("asd", 12)
	fmt.Println(m.Get("asda"))
}
