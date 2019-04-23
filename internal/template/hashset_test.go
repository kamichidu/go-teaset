package template

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	testSet(t, func() set {
		return NewHashSet()
	})
}
