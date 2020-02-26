package main

import (
	"yum/ast"

	"github.com/alecthomas/repr"
)

func main() {
	// if len(os.Args) != 2 {
	// 	println("Usage: yum file.yu")
	// 	return
	// }
	// filename := os.Args[1]
	nodes, err := ast.ParseFile("./examples/main.yum")
	if err != nil {
		panic(err)
	}
	println(repr.String(nodes, repr.Indent("  ")))
}
