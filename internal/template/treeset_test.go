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
