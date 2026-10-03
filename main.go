package main

import (
	"sort"
)

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
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
	keys := make([]string, 0 , len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {

}
