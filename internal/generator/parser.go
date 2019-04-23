package generator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"

	"github.com/kamichidu/go-teaset/internal/assets"
)

const (
	templateDir = "/internal/template/"
)

var (
	// for debugging
	UseLocal = false
)

func ParseFile(baseImpl string) (*ast.File, error) {
	name := path.Join(templateDir, strings.ToLower(baseImpl)+".go")
	src := assets.FSMustByte(UseLocal, name)
	fset := token.NewFileSet()
	return parser.ParseFile(fset, name, src, parser.AllErrors)
}

func ParseElementType(s string) (pkgPath, typName string) {
	idx := strings.LastIndex(s, ".")
	if idx < 0 {
		return "", s
	}
	x, sel := s[:idx], s[idx+1:]
	return strings.Trim(x, `"`), sel
}
