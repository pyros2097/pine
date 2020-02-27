package main

import (
	"fmt"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	// if len(os.Args) != 2 {
	// 	println("Usage: yum file.yu")
	// 	return
	// }
	// filename := os.Args[1]
	// nodes, err := ast.ParseFile("./examples/main.yum")
	// if err != nil {
	// 	panic(err)
	// }
	// println(repr.String(nodes, repr.Indent("  ")))
	bytes, err := wasm.ReadBytes("./code_gen/.snapshots/main.wasm")
	if err != nil {
		panic(err)
	}
	// Instantiates the WebAssembly module.
	instance, err := wasm.NewInstance(bytes)
	if err != nil {
		panic(err)
	}
	defer instance.Close()

	mainFunc := instance.Exports["main"]

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, err := mainFunc(5.0, 1.0)
	if err != nil {
		panic(err)
	}

	fmt.Println(result) // 42!
}
