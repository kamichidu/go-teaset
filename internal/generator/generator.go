package generator

import (
	"go/ast"
	"go/format"
	"go/token"
	"io"

	"golang.org/x/tools/go/ast/astutil"
)

func Traverse(w io.Writer, aFile *ast.File, imports []string, applyer astutil.ApplyFunc) error {
	fset := token.NewFileSet()
	for _, imp := range imports {
		astutil.AddImport(fset, aFile, imp)
	}
	modified := astutil.Apply(aFile, applyer, nil)
	return format.Node(w, fset, modified)
}
