package template

import (
	"sync"
)

type TreeSet struct {
	compare func(Element, Element) int

	// elements, smaller index is smaller item
	eles []Element

	mu sync.RWMutex
}

func NewTreeSet(compare func(Element, Element) int) *TreeSet {
	return &TreeSet{
		compare: compare,
	}
}

func (s *TreeSet) Add(v Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.add(v)
}

func (s *TreeSet) AddAll(l ...Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range l {
		s.add(v)
	}
}

func (s *TreeSet) add(v Element) {
	idx := -1
	for i, ele := range s.eles {
		// find greater one
		cmp := s.compare(v, ele)
		if cmp < 0 {
			idx = i
			break
		} else if cmp == 0 {
			// if it is equal one, nothing to do
			return
		}
	}
	if idx >= 0 {
		head, tail := s.eles[:idx], s.eles[idx:]
		s.eles = s.eles[:0]
		s.eles = append(s.eles, head...)
		s.eles = append(s.eles, v)
		s.eles = append(s.eles, tail...)
	} else {
		s.eles = append(s.eles, v)
	}
}

func (s *TreeSet) Remove(v Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.remove(v)
}

func (s *TreeSet) RemoveAll(l ...Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range l {
		s.remove(v)
	}
}

func (s *TreeSet) remove(v Element) {
	idx := -1
	for i, ele := range s.eles {
		// find equal one
		if s.compare(v, ele) == 0 {
			idx = i
			break
		}
	}
	if idx < 0 {
		return
	}
	head, tail := s.eles[:idx], s.eles[idx+1:]
	s.eles = s.eles[:0]
	s.eles = append(s.eles, head...)
	s.eles = append(s.eles, tail...)
}

func (s *TreeSet) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eles = s.eles[:0]
}

func (s *TreeSet) Contains(v Element) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.contains(v)
}

func (s *TreeSet) ContainsAll(l ...Element) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range l {
		if !s.contains(v) {
			return false
		}
	}
	return true
}

func (s *TreeSet) contains(v Element) bool {
	for _, ele := range s.eles {
		if s.compare(v, ele) == 0 {
			return true
		}
	}
	return false
}

func (s *TreeSet) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.eles)
}

func (s *TreeSet) ToSlice() []Element {
	s.mu.RLock()
	defer s.mu.RUnlock()
	l := make([]Element, len(s.eles))
	for i, v := range s.eles {
		l[i] = v
	}
	return l
}
