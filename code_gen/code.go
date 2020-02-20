package code_gen

import (
	"io"
	"yum/ast"
)

var typeMap = map[string]byte{
	"int": 0x7f,
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
			types = append(types, typeMap[fun.ReturnType.Name]) //i32
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
				0x6a, 0x0b, //  i32.add, end
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

// 0000000: 0061 736d                                 ; WASM_BINARY_MAGIC
// 0000004: 0100 0000                                 ; WASM_BINARY_VERSION
// ; section "Type" (1)
// 0000008: 01                                        ; section code
// 0000009: 00                                        ; section size (guess)
// 000000a: 02                                        ; num types
// ; type 0
// 000000b: 60                                        ; func
// 000000c: 02                                        ; num params
// 000000d: 7c                                        ; f64
// 000000e: 7c                                        ; f64
// 000000f: 01                                        ; num results
// 0000010: 7c                                        ; f64
// ; type 1
// 0000011: 60                                        ; func
// 0000012: 02                                        ; num params
// 0000013: 7f                                        ; i32
// 0000014: 7f                                        ; i32
// 0000015: 01                                        ; num results
// 0000016: 7f                                        ; i32
// 0000009: 0d                                        ; FIXUP section size
// ; section "Function" (3)
// 0000017: 03                                        ; section code
// 0000018: 00                                        ; section size (guess)
// 0000019: 02                                        ; num functions
// 000001a: 00                                        ; function 0 signature index
// 000001b: 01                                        ; function 1 signature index
// 0000018: 03                                        ; FIXUP section size
// ; section "Export" (7)
// 000001c: 07                                        ; section code
// 000001d: 00                                        ; section size (guess)
// 000001e: 02                                        ; num exports
// 000001f: 03                                        ; string length
// 0000020: 7375 62                                  sub  ; export name
// 0000023: 00                                        ; export kind
// 0000024: 00                                        ; export func index
// 0000025: 04                                        ; string length
// 0000026: 6d61 696e                                main  ; export name
// 000002a: 00                                        ; export kind
// 000002b: 01                                        ; export func index
// 000001d: 0e                                        ; FIXUP section size
// ; section "Code" (10)
// 000002c: 0a                                        ; section code
// 000002d: 00                                        ; section size (guess)
// 000002e: 02                                        ; num functions
// ; function body 0
// 000002f: 00                                        ; func body size (guess)
// 0000030: 00                                        ; local decl count
// 0000031: 20                                        ; local.get
// 0000032: 00                                        ; local index
// 0000033: 20                                        ; local.get
// 0000034: 01                                        ; local index
// 0000035: a0                                        ; f64.add
// 0000036: 0b                                        ; end
// 000002f: 07                                        ; FIXUP func body size
// ; function body 1
// 0000037: 00                                        ; func body size (guess)
// 0000038: 00                                        ; local decl count
// 0000039: 20                                        ; local.get
// 000003a: 00                                        ; local index
// 000003b: 20                                        ; local.get
// 000003c: 01                                        ; local index
// 000003d: 6a                                        ; i32.add
// 000003e: 0b                                        ; end
// 0000037: 07                                        ; FIXUP func body size
// 000002d: 11                                        ; FIXUP section size
// ;; dump
// 0000000: 0061 736d 0100 0000 010d 0260 027c 7c01
// 0000010: 7c60 027f 7f01 7f03 0302 0001 070e 0203
// 0000020: 7375 6200 0004 6d61 696e 0001 0a11 0207
// 0000030: 0020 0020 01a0 0b07 0020 0020 016a 0b
