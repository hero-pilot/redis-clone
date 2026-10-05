package main

import (
	"errors"
	"sort"
)

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (value string, err error) {
	v, ok := s.data[key]
	if !ok {
		return v, errors.New("value doesn;t exist")
	}
	return v, nil
}
func (s *Store) Set(key string, value string) {
	s.data[key] = value
}
func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (s *Store) Len() int {
	return len(s.data)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Rename(oldKey, newKey string) {
	v, error := s.Get(oldKey)
	if error != nil {
		return
	}
	s.Set(newKey, v)
	s.Delete(oldKey)
}

func (s *Store) Pop(key string) (string, error) {
	v, error := s.Get(key)
	if error != nil {
		return "", error
	}

	s.Delete(key)
	return v, nil
}
