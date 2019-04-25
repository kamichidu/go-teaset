package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"path"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

type Import struct {
	Path, PkgName string
}

func Traverse(w io.Writer, fset *token.FileSet, aFile *ast.File, importSpecs []Import, applyer astutil.ApplyFunc) error {
	for _, imp := range importSpecs {
		if imp.PkgName != "" && path.Base(imp.Path) != imp.PkgName {
			astutil.AddNamedImport(fset, aFile, imp.PkgName, imp.Path)
		} else {
			astutil.AddImport(fset, aFile, imp.Path)
		}
	}
	modified := astutil.Apply(aFile, applyer, nil)
	var buffer bytes.Buffer
	if err := format.Node(&buffer, fset, modified); err != nil {
		return err
	}
	src, err := imports.Process("", buffer.Bytes(), &imports.Options{
		AllErrors: true,
		// avoiding auto append import paths.
		// when FormatOnly don't add pkg name if basename of pkg path differs pkg name.
		FormatOnly: true,
		Comments:   true,
	})
	if err != nil {
		return err
	}
	_, err = w.Write(src)
	return err
}
