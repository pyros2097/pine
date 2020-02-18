package main

import (
	"os"

	"yum/ast"
	"yum/code_gen"

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
	file, err := os.Create("main.wasm")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	code_gen.GenerateCode(file, nodes)
}

// wasmer run ww.wat
// wat2wasm -d -v ww.wat
// wasmtime run main.wasm
// wasmtime run main.wasm
// vim -b ww.wasm
