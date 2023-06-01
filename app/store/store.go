package store

import (
	"fmt"
	"time"
)

type Item struct {
	Value  string
	Expiry time.Time
}

func (i Item) IsExpired() bool {
	if i.Expiry.IsZero() {
		return false
	}

	return i.Expiry.Before(time.Now())
}

type ErrKeyNotFound struct {
	Key string
}

func (e ErrKeyNotFound) Error() string {
	return fmt.Sprintf("key not found: %s", e.Key)
}

type MemStore struct {
	Store map[string]Item
}

func (m *MemStore) Set(key, value string, px time.Duration) {
	i := Item{
		Value: value,
	}
	if px != time.Millisecond*0 {
		i.Expiry = time.Now().Add(px)
	}
	// if px
	m.Store[key] = i
}

func (m *MemStore) Get(key string) (string, error) {
	value, found := m.Store[key]
	if found == false {
		return "", ErrKeyNotFound{Key: key}
	}

	if value.IsExpired() {
		delete(m.Store, key)
		return "", ErrKeyNotFound{Key: key}
	}

	return value.Value, nil

}

func New() *MemStore {
	return &MemStore{Store: make(map[string]Item)}
}
