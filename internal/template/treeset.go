package template

import (
	"sync"

	"github.com/google/btree"
)

type TreeSet struct {
	compare func(Element, Element) int

	eles *btree.BTree

	mu sync.RWMutex
}

type _TreeSetElement struct {
	compare func(Element, Element) int

	Value Element
}

func (e *_TreeSetElement) Less(other_ btree.Item) bool {
	other := other_.(*_TreeSetElement)
	return e.compare(e.Value, other.Value) < 0
}

func NewTreeSet(compare func(Element, Element) int) *TreeSet {
	return &TreeSet{
		compare: compare,
		// 2-3-4 tree
		eles: btree.New(2),
	}
}

func (s *TreeSet) toItem(v Element) btree.Item {
	return &_TreeSetElement{
		compare: s.compare,
		Value:   v,
	}
}

func (s *TreeSet) fromItem(v btree.Item) Element {
	return v.(*_TreeSetElement).Value
}

func (s *TreeSet) Add(v Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eles.ReplaceOrInsert(s.toItem(v))
}

func (s *TreeSet) AddAll(l ...Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range l {
		s.eles.ReplaceOrInsert(s.toItem(v))
	}
}

func (s *TreeSet) Remove(v Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eles.Delete(s.toItem(v))
}

func (s *TreeSet) RemoveAll(l ...Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range l {
		s.eles.Delete(s.toItem(v))
	}
}

func (s *TreeSet) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.eles.Clear(false)
}

func (s *TreeSet) Contains(v Element) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.eles.Has(s.toItem(v))
}

func (s *TreeSet) ContainsAll(l ...Element) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range l {
		if !s.eles.Has(s.toItem(v)) {
			return false
		}
	}
	return true
}

func (s *TreeSet) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.eles.Len()
}

func (s *TreeSet) ToSlice() []Element {
	s.mu.RLock()
	defer s.mu.RUnlock()
	l := make([]Element, 0, s.eles.Len())
	s.eles.Ascend(func(v btree.Item) bool {
		l = append(l, s.fromItem(v))
		return true
	})
	return l
}
