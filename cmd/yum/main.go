package main

import (
	"os"
	"yum/compiler"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: yum file.yum")
		return
	}
	filename := os.Args[1]
	nodes, err := compiler.ParseFile(filename)
	if err != nil {
		panic(err)
	}
	wasmData, err := compiler.NewEmitter(nodes).EmitAll()
	if err != nil {
		panic(err)
	}
	instance, err := wasm.NewInstance(wasmData.Bytes())
	if err != nil {
		panic(err)
	}
	defer instance.Close()

	if _, ok := instance.Exports["main"]; !ok {
		panic("You need to have a function named 'main' which takes no arguments and return none")
	}
	_, err = instance.Exports["main"]()
	if err != nil {
		panic(err)
	}
}

// register_route("path", pathParameters, queryPameters, callback) will be in c
// callback: will be wasm func which in this case will take pathParams and queryParams as parameters
// @route("/users/:id", "name", "date") will be a decorator function in wasm
// which will get the details of the callback -> fn index, params and call register route

// func route(path string, f ast.Func) {
// 	pathParams := []string{}
// 	queryParams := []string{}
//   f.returnType != 'Html' panic('asda')
// 	register_route("path", pathParameters, queryPameters, f.Index)
// }
