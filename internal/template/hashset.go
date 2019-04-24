package template

import (
	"sync"
)

type HashSet struct {
	eles map[Element]struct{}

	mu sync.RWMutex
}

func NewHashSet() *HashSet {
	return &HashSet{
		eles: map[Element]struct{}{},
	}
}

func (s *HashSet) Add(v Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eles[v] = struct{}{}
}

func (s *HashSet) AddAll(l ...Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range l {
		s.eles[v] = struct{}{}
	}
}

func (s *HashSet) Remove(v Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.eles, v)
}

func (s *HashSet) RemoveAll(l ...Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range l {
		delete(s.eles, v)
	}
}

func (s *HashSet) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eles = map[Element]struct{}{}
}

func (s *HashSet) Contains(v Element) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.eles[v]
	return ok
}

func (s *HashSet) ContainsAll(l ...Element) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range l {
		if _, ok := s.eles[v]; !ok {
			return false
		}
	}
	return true
}

func (s *HashSet) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.eles)
}

// ToSlice returns an slice containing all of the elements in this set.
func (s *HashSet) ToSlice() []Element {
	s.mu.RLock()
	defer s.mu.RUnlock()
	l := make([]Element, 0, len(s.eles))
	for v := range s.eles {
		l = append(l, v)
	}
	return l
}
