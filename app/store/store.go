package store

import (
	"fmt"
)

type ErrKeyNotFound struct {
	Key string
}

func (e ErrKeyNotFound) Error() string {
	return fmt.Sprintf("key not found: %s", e.Key)
}

type MemStore struct {
	Store map[string]string
}

func (m *MemStore) Set(key, value string) {
	m.Store[key] = value
}

func (m *MemStore) Get(key string) (string, error) {
	value, found := m.Store[key]
	if found == false {
		return "", ErrKeyNotFound{Key: key}
	}

	return value, nil
}

func New() *MemStore {
	return &MemStore{Store: make(map[string]string)}
}
