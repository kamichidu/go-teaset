package template

import (
	"strings"
	"testing"
)

func TestTreeSet(t *testing.T) {
	testSet(t, func() set {
		return NewTreeSet(func(a, b Element) int {
			return strings.Compare(a.(string), b.(string))
		})
	})
}

func BenchmarkTreeSet(b *testing.B) {
	newSet := func() set {
		return NewTreeSet(func(a, b Element) int {
			return strings.Compare(a.(string), b.(string))
		})
	}
	b.Run("1", func(b *testing.B) {
		benchmarkSet(b, newSet, 1)
	})
	b.Run("100", func(b *testing.B) {
		benchmarkSet(b, newSet, 100)
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkSet(b, newSet, 1000)
	})
	b.Run("10000", func(b *testing.B) {
		benchmarkSet(b, newSet, 10000)
	})
}
