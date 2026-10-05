package main

import (
	"sort"
)

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (value string, exists bool) {
	v, ok := s.data[key]
	return v, ok
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

func (s *Store) Len()  int{
	return  len(s.data)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}


func (s *Store) Rename(oldKey, newKey string) {
	if v, exists := s.Get(oldKey); exists{
		s.Set(newKey, v)
		s.Delete(oldKey)
	} 
}


func (s *Store) Pop(key string) (string, bool) {
	if v, exists := s.Get(key); exists{
		s.Delete(key)
		return v, exists
	}
	return "", false
}