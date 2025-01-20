package main

import (
	"sync"
)

type syncMap[KeyT comparable, ValT any] struct {
	mu   sync.Mutex
	data map[KeyT]ValT
}

func (m *syncMap[KeyT, ValT]) Set(key KeyT, val ValT) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

func (m *syncMap[KeyT, ValT]) Delete(key KeyT) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *syncMap[KeyT, ValT]) Get(key KeyT) (ValT, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.data[key]
	return val, ok
}

func NewSyncMap[KeyT comparable, ValT any]() *syncMap[KeyT, ValT] {
    return &syncMap[KeyT, ValT]{data: make(map[KeyT]ValT)}
}