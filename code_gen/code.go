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
	for _, entry := range ast.Entries {
		for _, importStmt := range entry.Imports {
			if importStmt != nil {
				// sb.WriteString("const ")
				// sb.WriteString(importStmt.Name)
				// sb.WriteString(" = ")
				// sb.WriteString("require('" + importStmt.Name + "');")
				// sb.WriteString("\n")
			}
		}
		for _, fun := range entry.Funs {
			// ; section "Type" (1)
			// ; type 0
			types := []byte{0x01, 0x60, byte(len(fun.Fun.Parameters))} // num types, func, num params
			for _, v := range fun.Fun.Parameters {                     // i32, i32
				types = append(types, typeMap[v.Type.Name])
			}
			types = append(types, 0x01)                             // num results
			types = append(types, typeMap[fun.Fun.ReturnType.Name]) //i32

			w.Write([]byte{0x01, byte(len(types))}) // section code, section size
			w.Write(types)
		}
		for _, fun := range entry.Funs {
			// ; section "Function" (3)
			all := []byte{0x01, 0x00}             // num functions, function 0 signature index
			w.Write([]byte{0x03, byte(len(all))}) // section code, section size
			w.Write(all)

			// ; section "Export" (7)s
			exports := []byte{0x01, byte(len(fun.Fun.Name))}   // num exports, name length
			exports = append(exports, []byte(fun.Fun.Name)...) // func name
			exports = append(exports, 0x00, 0x00)              // export kind, export func index
			w.Write([]byte{0x07, byte(len(exports))})          // section code, section size
			w.Write(exports)

			//; section "Code" (10)
			funcbody := []byte{
				0x20, 0x00, //  local.get, local index
				0x20, 0x01, //  local.get, local index
				0x6a, 0x0b, //  i32.add, end
			}

			func0 := []byte{0x00}
			sec10 := []byte{0x01, byte(len(func0) + len(funcbody))}              // num functions,
			w.Write([]byte{0x0a, byte(len(sec10) + len(func0) + len(funcbody))}) // section code, section size
			w.Write(sec10)
			w.Write(func0)
			w.Write(funcbody)
		}
	}
}

// 0000000: 0061 736d                                 ; WASM_BINARY_MAGIC
// 0000004: 0100 0000                                 ; WASM_BINARY_VERSION
// ; section "Type" (1)
// 0000008: 01                                        ; section code
// 0000009: 00                                        ; section size (guess)
// 000000a: 01                                        ; num types
// ; type 0
// 000000b: 60                                        ; func
// 000000c: 02                                        ; num params
// 000000d: 7f                                        ; i32
// 000000e: 7f                                        ; i32
// 000000f: 01                                        ; num results
// 0000010: 7f                                        ; i32
// 0000009: 07                                        ; FIXUP section size
// ; section "Function" (3)
// 0000011: 03                                        ; section code
// 0000012: 00                                        ; section size (guess)
// 0000013: 01                                        ; num functions
// 0000014: 00                                        ; function 0 signature index
// 0000012: 02                                        ; FIXUP section size
// ; section "Export" (7)
// 0000015: 07                                        ; section code
// 0000016: 00                                        ; section size (guess)
// 0000017: 01                                        ; num exports
// 0000018: 04                                        ; string length
// 0000019: 6d61 696e                                main  ; export name
// 000001d: 00                                        ; export kind
// 000001e: 00                                        ; export func index
// 0000016: 08                                        ; FIXUP section size
// ; section "Code" (10)
// 000001f: 0a                                        ; section code
// 0000020: 00                                        ; section size (guess)
// 0000021: 01                                        ; num functions
// ; function body 0
// 0000022: 00                                        ; func body size (guess)
// 0000023: 00                                        ; local decl count
// 0000024: 20                                        ; local.get
// 0000025: 00                                        ; local index
// 0000026: 20                                        ; local.get
// 0000027: 01                                        ; local index
// 0000028: 6a                                        ; i32.add
// 0000029: 0b                                        ; end
// 0000022: 07                                        ; FIXUP func body size
// 0000020: 09                                        ; FIXUP section size
// ;; dump
// 0000000: 0061 736d 0100 0000 0107 0160 027f 7f01
// 0000010: 7f03 0201 0007 0801 046d 6169 6e00 000a
// 0000020: 0901 0700 2000 2001 6a0b

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
