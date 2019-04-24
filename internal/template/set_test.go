package template

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type set interface {
	Add(Element)
	AddAll(...Element)
	Remove(Element)
	RemoveAll(...Element)
	Clear()
	Contains(Element) bool
	ContainsAll(...Element) bool
	Len() int
	ToSlice() []Element
}

func testSet(t *testing.T, newSet func() set) {
	equalStringSet := func(t *testing.T, s set, expect ...string) {
		assert.Equal(t, len(expect), s.Len(), "equal length")
		l := s.ToSlice()
		if _, ok := s.(*HashSet); ok {
			sort.Slice(l, func(i, j int) bool {
				return strings.Compare(l[i].(string), l[j].(string)) < 0
			})
		}
		eles := []Element{}
		for _, v := range l {
			eles = append(eles, v)
		}
		assert.Equal(t, eles, l)
	}
	t.Run("Add", func(t *testing.T) {
		s := newSet()
		equalStringSet(t, s)
		s.Add("a")
		equalStringSet(t, s, "a")
	})
	t.Run("Add", func(t *testing.T) {
		s := newSet()
		equalStringSet(t, s)
		for i := 0; i < 100; i++ {
			s.Add("a")
		}
		equalStringSet(t, s, "a")
	})
	t.Run("AddAll", func(t *testing.T) {
		s := newSet()
		equalStringSet(t, s)
		s.AddAll("a", "b")
		equalStringSet(t, s, "a", "b")
	})
	t.Run("AddAll", func(t *testing.T) {
		s := newSet()
		equalStringSet(t, s)
		s.AddAll("a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "a", "b")
		equalStringSet(t, s, "a", "b")
	})
	t.Run("Remove", func(t *testing.T) {
		s := newSet()
		s.AddAll("a", "b", "c")
		s.Remove("b")
		equalStringSet(t, s, "a", "c")
	})
	t.Run("Remove", func(t *testing.T) {
		s := newSet()
		s.Remove("b")
		equalStringSet(t, s)
	})
	t.Run("RemoveAll", func(t *testing.T) {
		s := newSet()
		s.AddAll("a", "b", "c", "d")
		s.RemoveAll("b", "c")
		equalStringSet(t, s, "a", "d")
	})
	t.Run("RemoveAll", func(t *testing.T) {
		s := newSet()
		s.RemoveAll("b", "c")
		equalStringSet(t, s)
	})
	t.Run("Clear", func(t *testing.T) {
		s := newSet()
		s.Clear()
		equalStringSet(t, s)
	})
	t.Run("Clear", func(t *testing.T) {
		s := newSet()
		s.AddAll("a", "b", "c", "d")
		s.Clear()
		equalStringSet(t, s)
	})
	t.Run("Contains", func(t *testing.T) {
		s := newSet()
		s.AddAll("a", "b", "c", "d")
		assert.Equal(t, true, s.Contains("b"), "contains b")
		assert.Equal(t, false, s.Contains("e"), "contains e")
	})
	t.Run("Contains", func(t *testing.T) {
		s := newSet()
		assert.Equal(t, false, s.Contains("b"), "contains b")
		assert.Equal(t, false, s.Contains("e"), "contains e")
	})
	t.Run("ContainsAll", func(t *testing.T) {
		s := newSet()
		s.AddAll("a", "b", "c", "d")
		assert.Equal(t, true, s.ContainsAll("b", "c"), "contains b, c")
		assert.Equal(t, false, s.ContainsAll("b", "e"), "contains b, e")
	})
	t.Run("ContainsAll", func(t *testing.T) {
		s := newSet()
		assert.Equal(t, false, s.ContainsAll("b", "c"), "contains b, c")
		assert.Equal(t, false, s.ContainsAll("b", "e"), "contains b, e")
	})
	t.Run("Len", func(t *testing.T) {
		s := newSet()
		s.AddAll("a", "b", "c", "d")
		assert.Equal(t, 4, s.Len())
	})
	t.Run("Len", func(t *testing.T) {
		s := newSet()
		assert.Equal(t, 0, s.Len())
	})
}
