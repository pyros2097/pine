package code_gen

import (
	"io"
	"yum/ast"
)

var typeMap = map[string]byte{
	"int": 0x7f,
	"float": 0x7c,
}

var operandMap = map[string]map[string]byte{
	"int": map[string]byte{
		"add": 0x6a,
	},
	"float": map[string]byte{
		"add": 0xa0,
	},
}

func GenerateFunc(fun *ast.FunDecl) {
}

func GenerateCode(w io.Writer, ast *ast.Ast) {
	w.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	w.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION

	funcCount := 0
	types := []byte{} // num types, func, num params
	funcs := []byte{}
	exports := []byte{}
	funcBodys := []byte{}
	for _, entry := range ast.Entries {
		for _, fun := range entry.Funs {
			types = append(types, 0x60, byte(len(fun.Parameters))) // func, num params
			for _, v := range fun.Parameters {                     // i32, i32
				types = append(types, typeMap[v.Type.Name])
			}
			types = append(types, 0x01)                         // num results
			types = append(types, typeMap[fun.ReturnType.Name]) // i32
		}
		for _, fun := range entry.Funs {
			funcs = append(funcs, byte(funcCount)) // function 0 signature index
			exports = append(exports, byte(len(fun.Name)))  // name length
			exports = append(exports, []byte(fun.Name)...) // name
			exports = append(exports, 0x00, byte(funcCount))                  	// export kind, export func indexs

			//; section "Code" (10)
			funcbody := []byte{
				0x00, // local decl count
				0x20, 0x00, //  local.get, local index
				0x20, 0x01, //  local.get, local index
				operandMap[fun.ReturnType.Name]["add"], 0x0b, //  i32.add, end
			}
			funcBodys = append(funcBodys, byte(len(funcbody))) // func body size
			funcBodys = append(funcBodys, funcbody...)
			funcCount += 1
		}
	}
	w.Write([]byte{0x01, byte(len(types) + 1), byte(funcCount)})          // Type section code, section size, num types, type data
	w.Write(types)
	w.Write([]byte{0x03, byte(len(funcs) + 1), byte(funcCount)}) // Func sig section code, section size, num types, type data
	w.Write(funcs)
	w.Write([]byte{0x07, byte(len(exports) + 1), byte(funcCount)}) // exports section code, section size, num exports
	w.Write(exports)
	w.Write([]byte{0x0a, byte(len(funcBodys) + 1), byte(funcCount)}) // func body section code, section size, num functions
	w.Write(funcBodys)
}

// ;; A function that takes an `i64` and returns
// ;; three `i32`s.
// (func (param i64) (result i32 i32 i32)
//   ...)

// ;; A loop that consumes an `i32` stack value
// ;; at the start of each iteration.
// loop (param i32)
//   ...
// end

// ;; A block that produces two `i32` stack values.
// block (result i32 i32)
//   ...
// end
