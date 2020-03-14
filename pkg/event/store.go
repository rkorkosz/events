package event

import "sync"

// Store describes event store methods
type Store interface {
	Put(Event) error
}

// InMemoryStore stores events in memory
type InMemoryStore struct {
	sync.RWMutex
	db map[string]Event
}

// NewInMemoryStore creates InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{db: make(map[string]Event)}
}

// Put implements Store.Put method for InMemoryStore
func (ms *InMemoryStore) Put(e Event) error {
	ms.Lock()
	ms.db[e.ID.String()] = e
	ms.Unlock()
	return nil
}
