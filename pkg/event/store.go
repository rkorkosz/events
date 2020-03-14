package event

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

// Store describes event store methods
type Store interface {
	Put(Event) error
	Get(ID uuid.UUID) (*Event, error)
}

// InMemoryStore stores events in memory
type InMemoryStore struct {
	sync.RWMutex
	db map[string]*Event
}

// NewInMemoryStore creates InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{db: make(map[string]*Event)}
}

// Put implements Store.Put method for InMemoryStore
func (ms *InMemoryStore) Put(e *Event) error {
	ms.Lock()
	ms.db[e.ID.String()] = e
	ms.Unlock()
	return nil
}

// Get retrieves Event by id
func (ms *InMemoryStore) Get(id uuid.UUID) (*Event, error) {
	event, ok := ms.db[id.String()]
	if !ok {
		return nil, fmt.Errorf("There is no event with ID: %s", id)
	}
	return event, nil
}

// BoltStore implements Store interface with bbolt database
type BoltStore struct {
	dbPath    string
	dbOptions *bolt.Options
}

// NewBoltStore creates BoltStore object
func NewBoltStore(dbPath string, dbOptions *bolt.Options) (*BoltStore, error) {
	db, err := bolt.Open(dbPath, 0600, dbOptions)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("events"))
		return err
	})
	return &BoltStore{dbPath: dbPath}, err
}

// Put implements Store.Put method
func (bs *BoltStore) Put(e *Event) error {
	db, err := bolt.Open(bs.dbPath, 0600, bs.dbOptions)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("events"))
		data, err := json.Marshal(e)
		if err != nil {
			return err
		}
		return b.Put([]byte(e.ID.String()), data)
	})
}

// Get gets event from store
func (bs *BoltStore) Get(id uuid.UUID) (*Event, error) {
	event := &Event{}
	db, err := bolt.Open(bs.dbPath, 0600, bs.dbOptions)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("events"))
		data := b.Get([]byte(id.String()))
		return json.Unmarshal(data, event)
	})
	return event, nil
}
