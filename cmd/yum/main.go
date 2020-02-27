package main

import (
	"fmt"
	"os"
	"yum/ast"
	"yum/code_gen"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: yum file.yum")
		return
	}
	filename := os.Args[1]
	nodes, err := ast.ParseFile(filename)
	if err != nil {
		panic(err)
	}
	wasmData, err := code_gen.GenerateCode(nodes)
	if err != nil {
		panic(err)
	}
	instance, err := wasm.NewInstance(wasmData.Bytes())
	if err != nil {
		panic(err)
	}
	defer instance.Close()

	mainFunc := instance.Exports["main"]
	result, err := mainFunc()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
