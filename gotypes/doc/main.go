// The doc command prints the doc comment of a package-level object.
package main

import (
	"fmt"
	"go/ast"
	"log"
	"os"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: doc <package> <object>")
	}
	//!+part1
	pkgpath, name := os.Args[1], os.Args[2]

	// Load complete type information for the specified packages,
	// along with type-annotated syntax.
	// Types for dependencies are loaded from export data.
	conf := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(conf, pkgpath)
	if err != nil {
		log.Fatal(err) // failed to load anything
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1) // some packages contained errors
	}

	// Find the package and package-level object.
	pkg := pkgs[0]
	obj := pkg.Types.Scope().Lookup(name)
	if obj == nil {
		log.Fatalf("%s.%s not found", pkg.Types.Path(), name)
	}
	//!-part1
	//!+part2

	// Print the object and its methods (incl. location of definition).
	fmt.Println(obj)
	for _, sel := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
		fmt.Printf("%s: %s\n", pkg.Fset.Position(sel.Obj().Pos()), sel)
	}

	// Find the path from the root of the AST to the object's position.
	// Walk up to the enclosing ast.Decl for the doc comment.
	for _, file := range pkg.Syntax {
		pos := obj.Pos()
		start := file.Pos()
		end := file.End()
		if !(start <= pos && pos < end) {
			continue // not in this file
		}
		path, _ := astutil.PathEnclosingInterval(file, pos, pos)
		for _, n := range path {
			switch n := n.(type) {
			case *ast.GenDecl:
				fmt.Println("\n", n.Doc.Text())
				return
			case *ast.FuncDecl:
				fmt.Println("\n", n.Doc.Text())
				return
			}
		}
	}
	//!-part2
}
