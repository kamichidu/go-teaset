package generator

import (
	"go/ast"
	"go/importer"
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

func ParseFile(baseImpl string) (*token.FileSet, *ast.File, error) {
	name := path.Join(templateDir, strings.ToLower(baseImpl)+".go")
	src := assets.FSMustByte(UseLocal, name)
	fset := token.NewFileSet()
	aFile, err := parser.ParseFile(fset, name, src, parser.AllErrors|parser.ParseComments)
	return fset, aFile, err
}

func ParseElementType(s string) (pkgPath, pkgName, typName string) {
	idx := strings.LastIndex(s, ".")
	if idx < 0 {
		return "", "", s
	}
	x, sel := s[:idx], s[idx+1:]
	pkgPath = strings.Trim(x, `"`)
	typName = sel
	tPkg, err := importer.ForCompiler(token.NewFileSet(), "source", nil).Import(pkgPath)
	if err != nil {
		panic(err)
	}
	return pkgPath, tPkg.Name(), typName
}
