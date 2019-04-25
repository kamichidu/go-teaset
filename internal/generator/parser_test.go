package generator

import (
	"testing"

	_ "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestParseElementType(t *testing.T) {
	cases := []struct {
		PkgName, PkgPath, TypName, Src string
	}{{
		PkgName: "uuid",
		PkgPath: "github.com/satori/go.uuid",
		TypName: "UUID",
		Src:     `"github.com/satori/go.uuid".UUID`,
	}, {
		PkgName: "time",
		PkgPath: "time",
		TypName: "Time",
		Src:     `time.Time`,
	}}
	for _, c := range cases {
		pkgPath, pkgName, typName := ParseElementType(c.Src)
		assert.Equal(t, c.PkgPath, pkgPath, "pkgPath")
		assert.Equal(t, c.PkgName, pkgName, "pkgName")
		assert.Equal(t, c.TypName, typName, "typName")
	}
}
