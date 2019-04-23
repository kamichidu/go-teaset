package generator

import (
	"go/ast"
	"log"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

var (
	Debug = false
)

func ApplyFunc(l ...astutil.ApplyFunc) astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		for _, fn := range l {
			if !fn(cursor) {
				return false
			}
		}
		return true
	}
}

func PrintCursorInfo(cursor *astutil.Cursor) bool {
	if node, ok := cursor.Node().(*ast.Ident); ok {
		log.Printf("cursor.Name = %q, *ast.Ident = %q", cursor.Name(), node.Name)
	}
	return true
}

func ReplaceIdentName(oldStr, newStr string) astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		if node, ok := cursor.Node().(*ast.Ident); ok {
			name := strings.Replace(node.Name, oldStr, newStr, 1)
			if node.Name != name {
				cursor.Replace(ast.NewIdent(name))
				return false
			}
		}
		return true
	}
}

func ReplaceExactIdentName(oldStr, newStr string) astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		if node, ok := cursor.Node().(*ast.Ident); ok {
			if node.Name == oldStr {
				cursor.Replace(ast.NewIdent(newStr))
				return false
			}
		}
		return true
	}
}
