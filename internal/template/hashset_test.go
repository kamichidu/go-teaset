package template

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	testSet(t, func() set {
		return NewHashSet()
	})
}

func BenchmarkHashSet(b *testing.B) {
	newSet := func() set {
		return NewHashSet()
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
