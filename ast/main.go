package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
)

func test(fix bool) string {
	// parse file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "ast/test.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	newImport := &ast.ImportSpec{
		Name: ast.NewIdent("name"),
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", "path"),
		},
	}
	impDecl := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			newImport,
		},
	}

	node.Decls = append(node.Decls, nil)
	copy(node.Decls[1:], node.Decls[0:])
	node.Decls = append(node.Decls, impDecl)
	node.Decls[0] = impDecl

	if fix {
		impDecl.TokPos = 1
		newImport.Name.NamePos = 0
	}

	buf := bytes.NewBufferString("")
	printer.Fprint(buf, fset, node)

	return buf.String()
}

func main() {
	print("-----bug:-----\n")
	fmt.Print(test(false))
	print("-----fix:-----\n")
	fmt.Print(test(true))
}
